package notification

import (
	"model"
	"fmt"
	"sort"
	"time"
	"strings"
	"net/http"
	"github.com/labstack/echo"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"github.com/gorilla/websocket"
	"github.com/satori/go.uuid"
)

type(
	// Handler : Handler for managing websockets and notifications.
	Handler struct {
		upgrader	websocket.Upgrader
		Manager 	clientManager
	}

	// ClientManager : Manages all connected websockets and forwards incoming notifications.
	clientManager struct {
		clients		map[string]map[uuid.UUID]*client
		Operator	chan model.Notification
		register	chan *client
		unregister	chan *client
		db 			*mgo.Session
		dbName 		string
	}

	// Client : Data structure that holds a single websocket connection.
	client struct {
		id			uuid.UUID
		username	string
		Socket		*websocket.Conn
		incoming	chan model.Notification	
	}
)

// NewHandler : Create a Handler instance.
func NewHandler(dbURL string) (h *Handler) {
	// Database connection
	session, err := mgo.Dial(dbURL)
	if err != nil {
		session.Close()
		panic(err)
	}

	h = &Handler{websocket.Upgrader{CheckOrigin: func(r *http.Request) bool {return true}}, clientManager{make(map[string]map[uuid.UUID]*client), 
		make(chan model.Notification), make(chan *client), make(chan *client), session, strings.Split(dbURL, "/")[3]}}

	go h.Manager.start()

	return h
}

// GetConnection : Upgrade a client connection to websocket.
//				   URL: "/api/v1/ws/:username"
//				   Method: GET
//				   Return 200 OK on success.
//				   Return 400 Bad Request if failed to make connection.
//				   Return 404 Not Found if the user is not in the database.
func (h *Handler) GetConnection(c echo.Context) (err error) {
	username := c.Param("username")

	db := h.Manager.db.Clone()
	defer db.Close()

	user := model.User{}
	err = db.DB(h.Manager.dbName).C(model.UserCollection).Find(bson.M{"username": username}).One(&user)
	if err != nil {
		if err == mgo.ErrNotFound {
			return &echo.HTTPError{Code: http.StatusNotFound, Message: "User does not exist."}
		}
		return
	}

	conn, err := h.upgrader.Upgrade(c.Response(), c.Request(), nil)
	if err != nil {
		return &echo.HTTPError{Code: http.StatusBadRequest, Message: "Failed to make web socket connection."}
	}

	cc := client{uuid.NewV1(), username, conn, make(chan model.Notification)}
	h.Manager.register <- &cc

	return c.NoContent(http.StatusOK)
}

func (manager *clientManager) start() {
	for {
		select {
		case message := <- manager.Operator:
			switch message.Detail.(type) {
			case model.NewTweetNotif:
				manager.forwardNewTweetNotif(message)
			case model.FollowNotif:
				manager.forwardFollowNotif(message)
			default:
				fmt.Println("Invalid notification type.")
			}
		case conn := <- manager.register:
			if _, ok := manager.clients[conn.username]; !ok {
				manager.clients[conn.username] = make(map[uuid.UUID]*client)
			}
			manager.clients[conn.username][conn.id] = conn
			go conn.read(manager)
			go conn.write(manager)
			go manager.FlushNotif(conn)
		case conn := <- manager.unregister:
			if _, ok := manager.clients[conn.username][conn.id]; ok {
				time.Sleep(2 * time.Second)
				conn.incoming <- model.Notification{Timestamp: time.Now(), Type: "", Detail: nil}
				conn.Socket.WriteMessage(websocket.CloseMessage, []byte{})
				conn.Socket.Close()
				delete(manager.clients[conn.username], conn.id)
			}
		}
	}
}

func (c *client) read(manager *clientManager) {
	defer func() {manager.unregister <- c}()

	for {
		t, m, err := c.Socket.ReadMessage()
		if err != nil {
			manager.unregister <- c
			break
		}

		switch t {
		case websocket.TextMessage:
			switch string(m) {
			case "Clear notifications.":
				manager.ClearNotif(c.username)
			default:
				c.Socket.WriteMessage(websocket.TextMessage, []byte("Invalid text message."))
				fmt.Println("Invalid text message from client.")
			}
		default:
			c.Socket.WriteMessage(websocket.TextMessage, []byte("Invalid message type."))
			fmt.Println("Invalid message type from client.")
		}
	}
}

func (c *client) write(manager *clientManager) {
	defer c.Socket.Close()

	for {
		message, ok := <- c.incoming
		if !ok {
			return
		}

		switch message.Detail.(type) {
		case model.NewTweetNotif:
			c.Socket.WriteMessage(websocket.TextMessage, []byte("New tweets."))
		case model.FollowNotif:
			c.Socket.WriteJSON(struct {
				Timestamp	time.Time	`json:"timestamp"`
				Type		string		`json:"type"`
				Detail		string		`json:"detail"`
			}{
				Timestamp: message.Timestamp,
				Type: message.Type,
				Detail: message.Detail.(model.FollowNotif).Follower,
			})
		case nil:
			return
		default:
			fmt.Println("Invalid notification type.")
		}
	}
}

// FlushNotif : Dump all notifications in the database to the client.
func (manager *clientManager) FlushNotif(conn *client) {
	time.Sleep(1 * time.Second)

	db := manager.db.Clone()
	defer db.Close()

	target := model.Individual{}
	err := db.DB(manager.dbName).C(model.NotificationCollection).Find(bson.M{"username": conn.username}).One(&target)
	if err != nil {
		panic(err)
	}

	sort.Slice(target.Notifications, func(i, j int) bool {return target.Notifications[i].Timestamp.Before(target.Notifications[j].Timestamp)})

	for _, m := range target.Notifications {
		conn.incoming <- m
	}
}

func (manager *clientManager) forwardNewTweetNotif(m model.Notification) {
	db := manager.db.Clone()
	defer db.Close()

	targets := []model.User{}
	err := db.DB(manager.dbName).C(model.UserCollection).Find(bson.M{"following": m.Detail.(model.NewTweetNotif).Publisher}).All(&targets)
	if err != nil {
		panic(err)
	}

	for _, t := range targets {
		if val, ok := manager.clients[t.Username]; ok {
			for _, c := range val {
				c.incoming <- m
			}
		}
	}
}

func (manager *clientManager) forwardFollowNotif(m model.Notification) {
	db := manager.db.Clone()
	defer db.Close()

	err := db.DB(manager.dbName).C(model.NotificationCollection).Update(bson.M{"username": m.Detail.(model.FollowNotif).Followee}, bson.M{"$addToSet": bson.M{"notifications": m}})
	if err != nil {
		panic(err)
	}

	if cs, ok := manager.clients[m.Detail.(model.FollowNotif).Followee]; ok {
		for _, c := range cs {
			c.incoming <- m
		}
	}
}

// ClearNotif : Empty the unacked notification list in the database for a specific user.
func (manager *clientManager) ClearNotif(username string) {
	db := manager.db.Clone()
	defer db.Close()

	err := db.DB(manager.dbName).C(model.NotificationCollection).Update(bson.M{"username": username}, bson.M{"$set": bson.M{"notifications": make([]model.Notification, 0)}})
	if err != nil {
		panic(err)
	}
}

// Shutdown : Gracefully shutdown notification handler.
func (h *Handler) Shutdown() {
	h.Manager.db.Close()

	for _, u := range h.Manager.clients {
		for _, c := range u {
			h.Manager.unregister <- c
		}
	}
}
