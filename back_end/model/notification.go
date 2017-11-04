package model

import "time"
import "gopkg.in/mgo.v2/bson"

// NotificationCollection : Notification collection name in database.
const NotificationCollection = "notifications"

type (
	// Individual : Data structure that contains notifications for a single user.
	Individual struct {
		ID				bson.ObjectId	`bson:"_id"`
		Username		string			`bson:"username"`
		Notifications	[]Notification	`bson:"notifications"`
	}

	// Notification : Data structure that holds a notification.
	Notification struct {
		Timestamp	time.Time	`bson:"timestamp"`
		Detail		interface{}	`bson:"detail"`
	}

	// NewTweetNotif : Data structure that holds a new tweet notification.
	NewTweetNotif struct {
		Publisher	string		`bson:"publisher"`
	}

	// FollowNotif : Data structure that holds a follow notification.
	FollowNotif struct {
		Followee	string		`bson:"followee"`
		Follower	string		`bson:"follower"`
	}
)
