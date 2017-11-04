package model

import "time"
import "gopkg.in/mgo.v2/bson"

const (

)

type (
	// Individual : Data structure that contains notifications for a single user.
	Individual struct {
		ID				bson.ObjectId	`bson:"_id"`
		Username		string			`bson:"username"`
		Notifications	[]interface{}	`bson:"notifications"`
	}

	// NewTweetNotif : Data structure that holds a new tweet notification.
	NewTweetNotif struct {
		Timestamp	time.Time	`bson:"timestamp"`
		Publisher	string		`bson:"publisher"`
	}

	// FollowNotif : Data structure that holds a follow notification.
	FollowNotif struct {
		Followee	string	`bson:"followee"`
		Follower	string	`bson:"follower"`
	}

	/*// Data structure that holds a 
	Notification struct {
		ID        	bson.ObjectId 	`json:"-" bson:"_id"`
		Username	string			`json:"username" bson:"username"`
	
	}*/
)
