package model

import "gopkg.in/mgo.v2/bson"

type User struct {
	ID        	bson.ObjectId 	`json:"id" bson:"_id,omitempty"`
	FirstName 	string			`json:"firstname" bson:"firstname"`
	LastName	string			`json:"lastname,omitempty" bson:"lastname,omitempty"`
	Password 	string			`json:"password,omitempty" bson:"password"`
	Email 		string			`json:"email" bson:"email"`
	Followers	[]string		`json:"followers,omitempty" bson:"followers,omitempty"`
	Followed	[]string		`json:"followed,omitempty" bson:"followed,omitempty"`
	Token		string			`json:"token,omitempty" bson:"-"`
	
	UserIDdev	string			`json:"useriddev,omitempty" bson:"useriddev,omitempty"`
}
