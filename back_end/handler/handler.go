package handler

import "gopkg.in/mgo.v2"

// Handler : Data structure that holds a handler for a session.
type Handler struct {
	DB *mgo.Session
}