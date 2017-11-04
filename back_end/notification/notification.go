package notification

import (
	"model"
	"fmt"
	"time"
	"net/http"
	"github.com/labstack/echo"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"github.com/gorilla/websocket"
	"github.com/satori/go.uuid"
	//"github.com/golang-collections/collections/stack"
)

type(
	// Handler : Handler for managing websockets and notifications.
	Handler struct {
		upgrader	websocket.Upgrader
		Manager 	ClientManager
	}

	// ClientManager : Manages all connected websockets and forwards incoming notifications.
	ClientManager struct {
		Clients		map[string]map[uuid.UUID]*Client
		Operator	chan interface{}
		register	chan *Client
		Unregister	chan *Client
		db 			*mgo.Session
		dbName 		string
	}

	// Client : Data structure that holds a single websocket connection.
	Client struct {
		id			uuid.UUID
		username	string
		Socket		*websocket.Conn
		incoming	chan interface{}
		//notifStack	
	}
)

// NewHandler : Create a Handler instance
func NewHandler(db *mgo.Session, dbName string) (h *Handler) {
	h = &Handler{websocket.Upgrader{CheckOrigin: func(r *http.Request) bool {return true}}, ClientManager{make(map[string]map[uuid.UUID]*Client), 
		make(chan interface{}), make(chan *Client), make(chan *Client), db, dbName}}
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
	err = db.DB(h.Manager.dbName).C("users").Find(bson.M{"username": username}).One(&user)
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

	cc := Client{uuid.NewV1(), username, conn, make(chan interface{})}
	h.Manager.register <- &cc

	return c.NoContent(http.StatusOK)
}

func (manager *ClientManager) start() {
	for {
		select {
		case message := <- manager.Operator:
			switch message.(type) {
			case model.NewTweetNotif:
				manager.forwardNewTweetNotif(message.(model.NewTweetNotif))
			case model.FollowNotif:
				manager.forwardFollowNotif(message.(model.FollowNotif))
			default:
				fmt.Println("Invalid notification type.")
			}
		case conn := <- manager.register:
			if _, ok := manager.Clients[conn.username]; !ok {
				manager.Clients[conn.username] = make(map[uuid.UUID]*Client)
			}
			manager.Clients[conn.username][conn.id] = conn
			go conn.read(manager)
			go conn.write(manager)
			go manager.flushNotif(conn)
		case conn := <- manager.Unregister:
			if _, ok := manager.Clients[conn.username][conn.id]; ok {
				conn.incoming <- nil
				close(conn.incoming)
				conn.Socket.WriteMessage(websocket.CloseMessage, []byte{})
				conn.Socket.Close()
				delete(manager.Clients[conn.username], conn.id)
			}
		}
	}
}

func (c *Client) read(manager *ClientManager) {
	defer func() {manager.Unregister <- c}()

	for {
		if _, _, err := c.Socket.ReadMessage(); err != nil {
			manager.Unregister <- c
			break
		}


	}
}

func (c *Client) write(manager *ClientManager) {
	defer c.Socket.Close()

	for {
		message := <- c.incoming
		switch message.(type) {
		case model.NewTweetNotif:
			c.Socket.WriteMessage(websocket.TextMessage, []byte("New tweets."))
		case model.FollowNotif:
			
		case nil:
			return
		default:
			fmt.Println("Invalid notification type.")
		}
	}
}

func (manager *ClientManager) flushNotif(conn *Client) {
	time.Sleep(500 * time.Millisecond)

	db := manager.db.Clone()
	defer db.Close()

	target := model.Individual{}
	err := db.DB(manager.dbName).C("notification").Find(bson.M{"username": conn.username}).One(&target)
	if err != nil {
		if err == mgo.ErrNotFound {
			return
		}
		panic(err)
	}



	/*for target.Notifications {

	}*/



}

func (manager *ClientManager) forwardNewTweetNotif(m model.NewTweetNotif) {
	db := manager.db.Clone()
	defer db.Close()

	targets := []model.User{}
	err := db.DB(manager.dbName).C("user").Find(bson.M{"following": bson.M{"$in": m.Publisher}}).All(&targets)
	if err != nil {
		if err == mgo.ErrNotFound {
			return
		}
		panic(err)
	}

	for _, t := range targets {
		if val, ok := manager.Clients[t.Username]; ok {
			for _, c := range val {
				c.incoming <- m
			}
		}
	}
}

func (manager *ClientManager) forwardFollowNotif(m model.FollowNotif) {
	db := manager.db.Clone()
	defer db.Close()


}
