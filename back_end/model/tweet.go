package model

import "gopkg.in/mgo.v2/bson"
import "time"

// Tweet : Data structure that holds a single tweet.
type Tweet struct {
	ID			bson.ObjectId	`json:"id" bson:"_id"`
	Owner 		string			`json:"owner" bson:"owner"`
	From 		string			`json:"from,omitempty" bson:"from,omitempty"`
	Message 	string			`json:"message" bson:"message"`
	Timestamp	time.Time		`json:"timestamp,omitempty" bson:"timestamp"`
}