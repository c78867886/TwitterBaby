package model

import "gopkg.in/mgo.v2/bson"
import "time"

// CommentCollection : Comment collection name in database.
const CommentCollection = "comments"

// Comment : Data structure that holds a single comment.
type Comment struct {
	ID				bson.ObjectId	`json:"id" bson:"_id"`
	FromTweetID 	string			`json:"fromtweetid" bson:"fromtweetid"`
	FromUsername 		string		`json:"fromusername,omitempty" bson:"fromusername,omitempty"`
	Message 		string			`json:"message" bson:"message"`
	Timestamp	time.Time		`json:"timestamp,omitempty" bson:"timestamp"`
}