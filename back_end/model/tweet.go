package model

import "gopkg.in/mgo.v2/bson"
import "time"

type Tweet struct {
	ID			bson.ObjectId	`json:"id" bson:"_id,omitempty"`
	//To 			string			`json:"to" bson:"to"`
	From 		string			`json:"from" bson:"from"`
	Message 	string			`json:"message" bson:"message"`
	Timestamp	time.Time		`json:"timestamp" bson:"timstamp"`
}