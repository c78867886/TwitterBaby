package model

import "gopkg.in/mgo.v2/bson"
import "time"

// Tweet : Data structure that holds a single tweet.
type Tweet struct {
	ID			bson.ObjectId	`json:"id" bson:"_id"`
	//Owner 		bson.ObjectId	`json:"owner" bson:"owner"`
	Owner 		bson.ObjectId	`json:"owner,omitempty" bson:"owner,omitempty"`
	//From 		bson.ObjectId	`json:"from,omitempty" bson:"from,omitempty"`
	From 		string			`json:"from,omitempty" bson:"from,omitempty"`
	Message 	string			`json:"message" bson:"message"`
	Timestamp	time.Time		`json:"timestamp" bson:"timstamp"`
}