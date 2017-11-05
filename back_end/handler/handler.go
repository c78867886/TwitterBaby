package handler

import (
	"user"
	"tweet"
	"comment"
	"notification"
	"strings"
	"gopkg.in/mgo.v2"
)

// Handler : Handler for dealing with requests.
type Handler struct {
	DB 				*mgo.Session
	DBName 			string
	UserHandler		*user.Handler
	TweetHandler	*tweet.Handler
	CommentHandler	*comment.Handler
	NotifHandler	*notification.Handler
}

// Key : Key for signing tokens.
const Key  = "UYrtPaa0Pky7QZyVrkIfnouatG7LjTKystf0FGdOuDiXCZyCSuVz8YdK7OBeSrC"

// NewHandler : Create a Handler instance
func NewHandler(dbURL string) (h *Handler) {
	// Database connection
	session, err := mgo.Dial(dbURL)
	if err != nil {
		session.Close()
		panic(err)
	}

	// Initialize handler
	nh := notification.NewHandler(dbURL)
	uh := user.NewHandler(dbURL, Key, nh.Manager.Operator)
	th := tweet.NewHandler(dbURL, Key, nh.Manager.Operator)
	ch := comment.NewHandler(dbURL, Key)
	h = &Handler{DB: session, DBName: strings.Split(dbURL, "/")[3], UserHandler: uh, TweetHandler: th, CommentHandler: ch, NotifHandler: nh}

	return h
}

// Shutdown : Gracefully shutdown handler.
func (h *Handler) Shutdown() {
	h.DB.Close()
	h.UserHandler.Shutdown()
	h.TweetHandler.Shutdown()
	h.CommentHandler.Shutdown()
	h.NotifHandler.Shutdown()
}
