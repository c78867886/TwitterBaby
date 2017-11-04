package model

import "gopkg.in/mgo.v2/bson"

// User : Data structure that holds a single user.
type User struct {
	ID        	bson.ObjectId 	`json:"-" bson:"_id"`
	Username	string			`json:"username" bson:"username"`						// unique
	FirstName 	string			`json:"firstname" bson:"firstname"`
	LastName	string			`json:"lastname,omitempty" bson:"lastname,omitempty"`
	Password 	string			`json:"password,omitempty" bson:"password"`
	Email 		string			`json:"email" bson:"email"`								// unique
	Followers	[]string		`json:"followers,omitempty" bson:"followers,omitempty"`
	Following	[]string		`json:"following,omitempty" bson:"following,omitempty"`
	Bio			string			`json:"bio,omitempty" bson:"bio,omitempty"`
	Tag			string			`json:"tag,omitempty" bson:"tag,omitempty"`
	Picture		string			`json:"picture,omitempty" bson:"picture,omitempty"`
	Token		string			`json:"token,omitempty" bson:"-"`
}
