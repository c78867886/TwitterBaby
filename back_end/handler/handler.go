package handler

import (
	"strings"
	"gopkg.in/mgo.v2"
)

// Handler : Handler for dealing with requests.
type Handler struct {
	DB 		*mgo.Session
	DBName 	string
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
	h = &Handler{DB: session, DBName: strings.Split(dbURL, "/")[3]}

	return h
}