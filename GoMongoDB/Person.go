package main

import (
	"gopkg.in/mgo.v2/bson"
	"time"
)

type Person struct {
	ID        bson.ObjectId `bson:"_id,omitempty"`
	Firstname      string
	Lastname      string
	Email     string
	Psw	string
	Age int
	Timestamp time.Time
}