package model

import "gopkg.in/mgo.v2/bson"

type Tweet struct {
	ID			bson.ObjectId	`json:"id" bson:"_id,omitempty"`
	//To 			string			`json:"to" bson:"to"`
	From 		string			`json:"from" bson:"from"`
	Message 	string			`json:"message" bson:"message"`
	Timestamp	string			`json:"timestamp" bson:"timstamp"`
}