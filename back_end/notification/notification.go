package notification

import (
	"model"
	"fmt"
	"net/http"
	"github.com/labstack/echo"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"github.com/gorilla/websocket"
	"github.com/dgrijalva/jwt-go"
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
		clients		map[string]*client
		Operator	chan interface{}
		DB 			*mgo.Session
		DBName 		string
	}

	client struct {
		username	string
		socket		*websocket.Conn
		incoming	chan interface{}
		//notifStack	
	}

	// NewTweetNotif : Requests for forwarding new tweet notifications.
	NewTweetNotif struct {
		Publisher	string
	}

	// FollowNotif : Requests for forwarding follow notifications.
	FollowNotif struct {
		Followee	string
		Follower	string
	}
)

// NewHandler : Create a Handler instance
func NewHandler(db *mgo.Session, dbName string) (h *Handler) {
	h = &Handler{websocket.Upgrader{CheckOrigin: func(r *http.Request) bool {return true}}, ClientManager{make(map[string]*client), make(chan interface{}), db, dbName}}
	go h.Manager.start()

	return h
}

// GetConnection : Upgrade a client connection to websocket.
func (h *Handler) GetConnection(c echo.Context) (err error) {
	username := usernameFromToken(c)

	conn, err := h.upgrader.Upgrade(c.Response(), c.Request(), nil)
	if err != nil {
		return &echo.HTTPError{Code: http.StatusNotFound, Message: "Failed to make web socket connection."}
	}

	cc := client{username, conn, make(chan interface{})}
	h.Manager.clients[username] = &cc
	go cc.read()
	go cc.write()

	return c.NoContent(http.StatusOK)
}

func (manager *ClientManager) start() {
	for {
		message := <- manager.Operator
		switch message.(type) {
		case NewTweetNotif:
			manager.forwardNewTweetNotif(message.(NewTweetNotif))
		case FollowNotif:
			manager.forwardFollowNotif(message.(FollowNotif))
		default:
			fmt.Println("Invalid notification type.")
		}
	}
}

func (c *client) read() {
	defer c.socket.Close()

	for {

	}
}

func (c *client) write() {
	defer c.socket.Close()

	for {
		message := <- c.incoming
		switch message.(type) {
		case NewTweetNotif:
			c.socket.WriteMessage(websocket.TextMessage, []byte("New tweets."))
		case FollowNotif:
			
		default:
			fmt.Println("Invalid notification type.")
		}
	}
}

func (manager *ClientManager) forwardNewTweetNotif(m NewTweetNotif) {
	db := manager.DB.Clone()
	defer db.Close()

	target := []model.User{}
	err := db.DB(manager.DBName).C("user").Find(bson.M{"following": bson.M{"$in": m.Publisher}}).All(&target)
	if err != nil {
		if err == mgo.ErrNotFound {
			return
		}
		panic(err)
	}

	for _, t := range target {
		manager.clients[t.Username].incoming <- m
	}
}

func (manager *ClientManager) forwardFollowNotif(m FollowNotif) {
	db := manager.DB.Clone()
	defer db.Close()

	
}

func usernameFromToken(c echo.Context) string {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)

	return claims["username"].(string)
}