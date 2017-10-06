package handler

import "gopkg.in/mgo.v2"

// Handler : Handler for dealing with requests.
type Handler struct {
	DB *mgo.Session
}

// Key : Key for signing tokens.
const Key  = "UYrtPaa0Pky7QZyVrkIfnouatG7LjTKystf0FGdOuDiXCZyCSuVz8YdK7OBeSrC"