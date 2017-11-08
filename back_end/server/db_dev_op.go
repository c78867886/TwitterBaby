package server

import (
	"model"
	"handler"
	"time"
	"gopkg.in/mgo.v2/bson"
)

func dbReinsert() {
	h := handler.NewHandler("mongodb://SEavenger:SEavenger@ds149324.mlab.com:49324/se_avengers")

	err := h.DB.DB(h.DBName).DropDatabase()
	if err != nil {
		panic(err)
	}

	userC := h.DB.DB(h.DBName).C(model.UserCollection)

	users := []model.User {
		model.User{ID: bson.NewObjectId(), Username: "JasonHo", FirstName: "Jason", LastName: "Ho", Password: "test1", Email: "hojason117@gmail.com", Followers: []string{"MarsLee", "TomRiddle"}, Following: []string{"MarsLee"}, 
			Bio: "Hi everyone, this is Jason Ho.", Tag: "Albert Einstein"},
		model.User{ID: bson.NewObjectId(), Username: "MarsLee", FirstName: "Chih-Yin", LastName: "Lee", Password: "test2", Email: "c788678867886@gmail.com", Followers: []string{"JasonHo"}, Following: []string{"JasonHo"}, 
			Bio: "Hi everyone, this is Mars Lee.", Tag: "Bruno Mars"},
		model.User{ID: bson.NewObjectId(), Username: "JasonHe", FirstName: "Jason", LastName: "He", Password: "test3", Email: "hexing_h@hotmail.com", Followers: []string{}, Following: []string{}, 
			Bio: "Hi everyone, this is Jason He.", Tag: "Jason hehehehe"},
		model.User{ID: bson.NewObjectId(), Username: "DianeLin", FirstName: "Diane", LastName: "Lin", Password: "test4", Email: "diane@gmail.com", Followers: []string{}, Following: []string{}, 
			Bio: "Hi everyone, this is Diane Lin.", Tag: "Diane Kruger"}, 
		model.User{ID: bson.NewObjectId(), Username: "TomRiddle", FirstName: "Tom", LastName: "Riddle", Password: "test5", Email: "triddle@gmail.com", Followers: []string{}, Following: []string{"JasonHo"}, 
			Bio: "Hi everyone, this is Lord Voldemort.", Tag: "Voldemort"}, 
		model.User{ID: bson.NewObjectId(), Username: "JS", FirstName: "Jon", Password: "pass", Email: "json@gmail.com", Followers: []string{}, Following: []string{}},
	}

	for _, u := range users {
		err := userC.Insert(u)
		if err != nil {
			panic(err)
		}
	}

	notificationC := h.DB.DB(h.DBName).C(model.NotificationCollection)

	notifications := []model.Individual {
		model.Individual{ID: bson.NewObjectId(), Username: "JasonHo", Notifications: []model.Notification{model.Notification{Timestamp: time.Now(), Type: model.FollowType, Detail: model.FollowNotif{Followee: "JasonHo", Follower: "MarsLee"}},
			model.Notification{Timestamp: time.Now(), Type: model.FollowType, Detail: model.FollowNotif{Followee: "JasonHo", Follower: "TomRiddle"}}}},
		model.Individual{ID: bson.NewObjectId(), Username: "MarsLee", Notifications: make([]model.Notification, 0)},
		model.Individual{ID: bson.NewObjectId(), Username: "JasonHe", Notifications: make([]model.Notification, 0)},
		model.Individual{ID: bson.NewObjectId(), Username: "DianeLin", Notifications: make([]model.Notification, 0)},
		model.Individual{ID: bson.NewObjectId(), Username: "TomRiddle", Notifications: make([]model.Notification, 0)},
		model.Individual{ID: bson.NewObjectId(), Username: "JS", Notifications: make([]model.Notification, 0)},
	}

	for _, c := range notifications {
		err := notificationC.Insert(c)
		if err != nil {
			panic(err)
		}
	}

	tweetC := h.DB.DB(h.DBName).C(model.TweetCollection)

	tweets := []model.Tweet {
		model.Tweet{ID: bson.NewObjectId(), Owner: "JasonHo", Message: "Hi, I am Jason Ho. Weather sucks.", Timestamp: time.Now()}, 
		model.Tweet{ID: bson.NewObjectId(), Owner: "JasonHo", Message: "Hello from Jason Ho.", Timestamp: time.Now()}, 
		model.Tweet{ID: bson.NewObjectId(), Owner: "JasonHo", Message: "Hello world!", Timestamp: time.Now()}, 
		model.Tweet{ID: bson.NewObjectId(), Owner: "JasonHo", Message: "Harvest Moon: All you need to know It's the full moon closest to Sept equinox - coming up Oct 5.", Timestamp: time.Now()}, 
		model.Tweet{ID: bson.NewObjectId(), Owner: "JasonHo", Message: "Hey yo!!", Timestamp: time.Now()}, 
		model.Tweet{ID: bson.NewObjectId(), Owner: "MarsLee", Message: "Hi, I am Chih-Yin Lee. Weather sucks.", Timestamp: time.Now()}, 
		model.Tweet{ID: bson.NewObjectId(), Owner: "MarsLee", Message: "Hello from Chih-Yin Lee.", Timestamp: time.Now()}, 
		model.Tweet{ID: bson.NewObjectId(), Owner: "JasonHe", Message: "Hi, I am Jason He. Weather sucks.", Timestamp: time.Now()}, 
		model.Tweet{ID: bson.NewObjectId(), Owner: "JasonHe", Message: "Hello from Jason He.", Timestamp: time.Now()}, 
		model.Tweet{ID: bson.NewObjectId(), Owner: "DianeLin", Message: "Hi, I am Diane Lin. Weather sucks.", Timestamp: time.Now()}, 
		model.Tweet{ID: bson.NewObjectId(), Owner: "DianeLin", Message: "Hello from Diane Lin.", Timestamp: time.Now()},
		model.Tweet{ID: bson.NewObjectId(), Owner: "JasonHo", Message: "Hi, I am Jason Ho. Weather sucks.", Timestamp: time.Now()}, 
		model.Tweet{ID: bson.NewObjectId(), Owner: "JasonHo", Message: "Hello from Jason Ho.", Timestamp: time.Now()}, 
		model.Tweet{ID: bson.NewObjectId(), Owner: "JasonHo", Message: "Hello world!", Timestamp: time.Now()}, 
		model.Tweet{ID: bson.NewObjectId(), Owner: "JasonHo", Message: "Harvest Moon: All you need to know It's the full moon closest to Sept equinox - coming up Oct 5.", Timestamp: time.Now()}, 
		model.Tweet{ID: bson.NewObjectId(), Owner: "JasonHo", Message: "Hey yo!!", Timestamp: time.Now()}, 
		model.Tweet{ID: bson.NewObjectId(), Owner: "MarsLee", Message: "Hi, I am Chih-Yin Lee. Weather sucks.", Timestamp: time.Now()}, 
		model.Tweet{ID: bson.NewObjectId(), Owner: "MarsLee", Message: "Hello from Chih-Yin Lee.", Timestamp: time.Now()}, 
		model.Tweet{ID: bson.NewObjectId(), Owner: "JasonHe", Message: "Hi, I am Jason He. Weather sucks.", Timestamp: time.Now()}, 
		model.Tweet{ID: bson.NewObjectId(), Owner: "JasonHe", Message: "Hello from Jason He.", Timestamp: time.Now()}, 
		model.Tweet{ID: bson.NewObjectId(), Owner: "DianeLin", Message: "Hi, I am Diane Lin. Weather sucks.", Timestamp: time.Now()}, 
		model.Tweet{ID: bson.NewObjectId(), Owner: "DianeLin", Message: "Hello from Diane Lin.", Timestamp: time.Now()},
		model.Tweet{ID: bson.NewObjectId(), Owner: "JasonHo", Message: "Hi, I am Jason Ho. Weather sucks.", Timestamp: time.Now()}, 
		model.Tweet{ID: bson.NewObjectId(), Owner: "JasonHo", Message: "Hello from Jason Ho.", Timestamp: time.Now()}, 
		model.Tweet{ID: bson.NewObjectId(), Owner: "JasonHo", Message: "Hello world!", Timestamp: time.Now()}, 
		model.Tweet{ID: bson.NewObjectId(), Owner: "JasonHo", Message: "Harvest Moon: All you need to know It's the full moon closest to Sept equinox - coming up Oct 5.", Timestamp: time.Now()}, 
		model.Tweet{ID: bson.NewObjectId(), Owner: "JasonHo", Message: "Hey yo!!", Timestamp: time.Now()}, 
		model.Tweet{ID: bson.NewObjectId(), Owner: "MarsLee", Message: "Hi, I am Chih-Yin Lee. Weather sucks.", Timestamp: time.Now()}, 
		model.Tweet{ID: bson.NewObjectId(), Owner: "MarsLee", Message: "Hello from Chih-Yin Lee.", Timestamp: time.Now()}, 
		model.Tweet{ID: bson.NewObjectId(), Owner: "JasonHe", Message: "Hi, I am Jason He. Weather sucks.", Timestamp: time.Now()}, 
		model.Tweet{ID: bson.NewObjectId(), Owner: "JasonHe", Message: "Hello from Jason He.", Timestamp: time.Now()}, 
		model.Tweet{ID: bson.NewObjectId(), Owner: "DianeLin", Message: "Hi, I am Diane Lin. Weather sucks.", Timestamp: time.Now()}, 
		model.Tweet{ID: bson.NewObjectId(), Owner: "DianeLin", Message: "Hello from Diane Lin.", Timestamp: time.Now()},
		model.Tweet{ID: bson.NewObjectId(), Owner: "JasonHo", Message: "Hi, I am Jason Ho. Weather sucks.", Timestamp: time.Now()}, 
		model.Tweet{ID: bson.NewObjectId(), Owner: "JasonHo", Message: "Hello from Jason Ho.", Timestamp: time.Now()}, 
		model.Tweet{ID: bson.NewObjectId(), Owner: "JasonHo", Message: "Hello world!", Timestamp: time.Now()}, 
		model.Tweet{ID: bson.NewObjectId(), Owner: "JasonHo", Message: "Harvest Moon: All you need to know It's the full moon closest to Sept equinox - coming up Oct 5.", Timestamp: time.Now()}, 
		model.Tweet{ID: bson.NewObjectId(), Owner: "JasonHo", Message: "Hey yo!!", Timestamp: time.Now()}, 
		model.Tweet{ID: bson.NewObjectId(), Owner: "MarsLee", Message: "Hi, I am Chih-Yin Lee. Weather sucks.", Timestamp: time.Now()}, 
		model.Tweet{ID: bson.NewObjectId(), Owner: "MarsLee", Message: "Hello from Chih-Yin Lee.", Timestamp: time.Now()}, 
		model.Tweet{ID: bson.NewObjectId(), Owner: "JasonHe", Message: "Hi, I am Jason He. Weather sucks.", Timestamp: time.Now()}, 
		model.Tweet{ID: bson.NewObjectId(), Owner: "JasonHe", Message: "Hello from Jason He.", Timestamp: time.Now()}, 
		model.Tweet{ID: bson.NewObjectId(), Owner: "DianeLin", Message: "Hi, I am Diane Lin. Weather sucks.", Timestamp: time.Now()}, 
		model.Tweet{ID: bson.NewObjectId(), Owner: "DianeLin", Message: "Hello from Diane Lin.", Timestamp: time.Now()},
		model.Tweet{ID: bson.NewObjectId(), Owner: "JasonHo", Message: "Hi, I am Jason Ho. Weather sucks.", Timestamp: time.Now()}, 
		model.Tweet{ID: bson.NewObjectId(), Owner: "JasonHo", Message: "Hello from Jason Ho.", Timestamp: time.Now()}, 
		model.Tweet{ID: bson.NewObjectId(), Owner: "JasonHo", Message: "Hello world!", Timestamp: time.Now()}, 
		model.Tweet{ID: bson.NewObjectId(), Owner: "JasonHo", Message: "Harvest Moon: All you need to know It's the full moon closest to Sept equinox - coming up Oct 5.", Timestamp: time.Now()}, 
		model.Tweet{ID: bson.NewObjectId(), Owner: "JasonHo", Message: "Hey yo!!", Timestamp: time.Now()}, 
		model.Tweet{ID: bson.NewObjectId(), Owner: "MarsLee", Message: "Hi, I am Chih-Yin Lee. Weather sucks.", Timestamp: time.Now()}, 
		model.Tweet{ID: bson.NewObjectId(), Owner: "MarsLee", Message: "Hello from Chih-Yin Lee.", Timestamp: time.Now()}, 
		model.Tweet{ID: bson.NewObjectId(), Owner: "JasonHe", Message: "Hi, I am Jason He. Weather sucks.", Timestamp: time.Now()}, 
		model.Tweet{ID: bson.NewObjectId(), Owner: "JasonHe", Message: "Hello from Jason He.", Timestamp: time.Now()}, 
		model.Tweet{ID: bson.NewObjectId(), Owner: "DianeLin", Message: "Hi, I am Diane Lin. Weather sucks.", Timestamp: time.Now()}, 
		model.Tweet{ID: bson.NewObjectId(), Owner: "DianeLin", Message: "Hello from Diane Lin.", Timestamp: time.Now()},
		model.Tweet{ID: bson.NewObjectId(), Owner: "JasonHo", Message: "Hi, I am Jason Ho. Weather sucks.", Timestamp: time.Now()}, 
		model.Tweet{ID: bson.NewObjectId(), Owner: "JasonHo", Message: "Hello from Jason Ho.", Timestamp: time.Now()}, 
		model.Tweet{ID: bson.NewObjectId(), Owner: "JasonHo", Message: "Hello world!", Timestamp: time.Now()}, 
		model.Tweet{ID: bson.NewObjectId(), Owner: "JasonHo", Message: "Harvest Moon: All you need to know It's the full moon closest to Sept equinox - coming up Oct 5.", Timestamp: time.Now()}, 
		model.Tweet{ID: bson.NewObjectId(), Owner: "JasonHo", Message: "Hey yo!!", Timestamp: time.Now()}, 
		model.Tweet{ID: bson.NewObjectId(), Owner: "MarsLee", Message: "Hi, I am Chih-Yin Lee. Weather sucks.", Timestamp: time.Now()}, 
		model.Tweet{ID: bson.NewObjectId(), Owner: "MarsLee", Message: "Hello from Chih-Yin Lee.", Timestamp: time.Now()}, 
		model.Tweet{ID: bson.NewObjectId(), Owner: "JasonHe", Message: "Hi, I am Jason He. Weather sucks.", Timestamp: time.Now()}, 
		model.Tweet{ID: bson.NewObjectId(), Owner: "JasonHe", Message: "Hello from Jason He.", Timestamp: time.Now()}, 
		model.Tweet{ID: bson.NewObjectId(), Owner: "DianeLin", Message: "Hi, I am Diane Lin. Weather sucks.", Timestamp: time.Now()}, 
		model.Tweet{ID: bson.NewObjectId(), Owner: "DianeLin", Message: "Hello from Diane Lin.", Timestamp: time.Now()},
		model.Tweet{ID: bson.NewObjectId(), Owner: "JasonHo", Message: "Hi, I am Jason Ho. Weather sucks.", Timestamp: time.Now()}, 
		model.Tweet{ID: bson.NewObjectId(), Owner: "JasonHo", Message: "Hello from Jason Ho.", Timestamp: time.Now()}, 
		model.Tweet{ID: bson.NewObjectId(), Owner: "JasonHo", Message: "Hello world!", Timestamp: time.Now()}, 
		model.Tweet{ID: bson.NewObjectId(), Owner: "JasonHo", Message: "Harvest Moon: All you need to know It's the full moon closest to Sept equinox - coming up Oct 5.", Timestamp: time.Now()}, 
		model.Tweet{ID: bson.NewObjectId(), Owner: "JasonHo", Message: "Hey yo!!", Timestamp: time.Now()}, 
		model.Tweet{ID: bson.NewObjectId(), Owner: "MarsLee", Message: "Hi, I am Chih-Yin Lee. Weather sucks.", Timestamp: time.Now()}, 
		model.Tweet{ID: bson.NewObjectId(), Owner: "MarsLee", Message: "Hello from Chih-Yin Lee.", Timestamp: time.Now()}, 
		model.Tweet{ID: bson.NewObjectId(), Owner: "JasonHe", Message: "Hi, I am Jason He. Weather sucks.", Timestamp: time.Now()}, 
		model.Tweet{ID: bson.NewObjectId(), Owner: "JasonHe", Message: "Hello from Jason He.", Timestamp: time.Now()}, 
		model.Tweet{ID: bson.NewObjectId(), Owner: "DianeLin", Message: "Hi, I am Diane Lin. Weather sucks.", Timestamp: time.Now()}, 
		model.Tweet{ID: bson.NewObjectId(), Owner: "DianeLin", Message: "Hello from Diane Lin.", Timestamp: time.Now()},
		model.Tweet{ID: bson.NewObjectId(), Owner: "JasonHo", Message: "Hi, I am Jason Ho. Weather sucks.", Timestamp: time.Now()}, 
		model.Tweet{ID: bson.NewObjectId(), Owner: "JasonHo", Message: "Hello from Jason Ho.", Timestamp: time.Now()}, 
		model.Tweet{ID: bson.NewObjectId(), Owner: "JasonHo", Message: "Hello world!", Timestamp: time.Now()}, 
		model.Tweet{ID: bson.NewObjectId(), Owner: "JasonHo", Message: "Harvest Moon: All you need to know It's the full moon closest to Sept equinox - coming up Oct 5.", Timestamp: time.Now()}, 
		model.Tweet{ID: bson.NewObjectId(), Owner: "JasonHo", Message: "Hey yo!!", Timestamp: time.Now()}, 
		model.Tweet{ID: bson.NewObjectId(), Owner: "MarsLee", Message: "Hi, I am Chih-Yin Lee. Weather sucks.", Timestamp: time.Now()}, 
		model.Tweet{ID: bson.NewObjectId(), Owner: "MarsLee", Message: "Hello from Chih-Yin Lee.", Timestamp: time.Now()}, 
		model.Tweet{ID: bson.NewObjectId(), Owner: "JasonHe", Message: "Hi, I am Jason He. Weather sucks.", Timestamp: time.Now()}, 
		model.Tweet{ID: bson.NewObjectId(), Owner: "JasonHe", Message: "Hello from Jason He.", Timestamp: time.Now()}, 
		model.Tweet{ID: bson.NewObjectId(), Owner: "DianeLin", Message: "Hi, I am Diane Lin. Weather sucks.", Timestamp: time.Now()}, 
		model.Tweet{ID: bson.NewObjectId(), Owner: "DianeLin", Message: "Hello from Diane Lin.", Timestamp: time.Now()},
		model.Tweet{ID: bson.NewObjectId(), Owner: "JasonHo", Message: "Hi, I am Jason Ho. Weather sucks.", Timestamp: time.Now()}, 
		model.Tweet{ID: bson.NewObjectId(), Owner: "JasonHo", Message: "Hello from Jason Ho.", Timestamp: time.Now()}, 
		model.Tweet{ID: bson.NewObjectId(), Owner: "JasonHo", Message: "Hello world!", Timestamp: time.Now()}, 
		model.Tweet{ID: bson.NewObjectId(), Owner: "JasonHo", Message: "Harvest Moon: All you need to know It's the full moon closest to Sept equinox - coming up Oct 5.", Timestamp: time.Now()}, 
		model.Tweet{ID: bson.NewObjectId(), Owner: "JasonHo", Message: "Hey yo!!", Timestamp: time.Now()}, 
		model.Tweet{ID: bson.NewObjectId(), Owner: "MarsLee", Message: "Hi, I am Chih-Yin Lee. Weather sucks.", Timestamp: time.Now()}, 
		model.Tweet{ID: bson.NewObjectId(), Owner: "MarsLee", Message: "Hello from Chih-Yin Lee.", Timestamp: time.Now()}, 
		model.Tweet{ID: bson.NewObjectId(), Owner: "JasonHe", Message: "Hi, I am Jason He. Weather sucks.", Timestamp: time.Now()}, 
		model.Tweet{ID: bson.NewObjectId(), Owner: "JasonHe", Message: "Hello from Jason He.", Timestamp: time.Now()}, 
		model.Tweet{ID: bson.NewObjectId(), Owner: "DianeLin", Message: "Hi, I am Diane Lin. Weather sucks.", Timestamp: time.Now()}, 
		model.Tweet{ID: bson.NewObjectId(), Owner: "DianeLin", Message: "Hello from Diane Lin.", Timestamp: time.Now()},
		model.Tweet{ID: bson.NewObjectId(), Owner: "JasonHo", Message: "Hi, I am Jason Ho. Weather sucks.", Timestamp: time.Now()}, 
		model.Tweet{ID: bson.NewObjectId(), Owner: "JasonHo", Message: "Hello from Jason Ho.", Timestamp: time.Now()}, 
		model.Tweet{ID: bson.NewObjectId(), Owner: "JasonHo", Message: "Hello world!", Timestamp: time.Now()}, 
		model.Tweet{ID: bson.NewObjectId(), Owner: "JasonHo", Message: "Harvest Moon: All you need to know It's the full moon closest to Sept equinox - coming up Oct 5.", Timestamp: time.Now()}, 
		model.Tweet{ID: bson.NewObjectId(), Owner: "JasonHo", Message: "Hey yo!!", Timestamp: time.Now()}, 
		model.Tweet{ID: bson.NewObjectId(), Owner: "MarsLee", Message: "Hi, I am Chih-Yin Lee. Weather sucks.", Timestamp: time.Now()}, 
		model.Tweet{ID: bson.NewObjectId(), Owner: "MarsLee", Message: "Hello from Chih-Yin Lee.", Timestamp: time.Now()}, 
		model.Tweet{ID: bson.NewObjectId(), Owner: "JasonHe", Message: "Hi, I am Jason He. Weather sucks.", Timestamp: time.Now()}, 
		model.Tweet{ID: bson.NewObjectId(), Owner: "JasonHe", Message: "Hello from Jason He.", Timestamp: time.Now()}, 
		model.Tweet{ID: bson.NewObjectId(), Owner: "DianeLin", Message: "Hi, I am Diane Lin. Weather sucks.", Timestamp: time.Now()}, 
		model.Tweet{ID: bson.NewObjectId(), Owner: "DianeLin", Message: "Hello from Diane Lin.", Timestamp: time.Now()},
		model.Tweet{ID: bson.NewObjectId(), Owner: "JasonHo", Message: "Hi, I am Jason Ho. Weather sucks.", Timestamp: time.Now()}, 
		model.Tweet{ID: bson.NewObjectId(), Owner: "JasonHo", Message: "Hello from Jason Ho.", Timestamp: time.Now()}, 
		model.Tweet{ID: bson.NewObjectId(), Owner: "JasonHo", Message: "Hello world!", Timestamp: time.Now()}, 
		model.Tweet{ID: bson.NewObjectId(), Owner: "JasonHo", Message: "Harvest Moon: All you need to know It's the full moon closest to Sept equinox - coming up Oct 5.", Timestamp: time.Now()}, 
		model.Tweet{ID: bson.NewObjectId(), Owner: "JasonHo", Message: "Hey yo!!", Timestamp: time.Now()}, 
		model.Tweet{ID: bson.NewObjectId(), Owner: "MarsLee", Message: "Hi, I am Chih-Yin Lee. Weather sucks.", Timestamp: time.Now()}, 
		model.Tweet{ID: bson.NewObjectId(), Owner: "MarsLee", Message: "Hello from Chih-Yin Lee.", Timestamp: time.Now()}, 
		model.Tweet{ID: bson.NewObjectId(), Owner: "JasonHe", Message: "Hi, I am Jason He. Weather sucks.", Timestamp: time.Now()}, 
		model.Tweet{ID: bson.NewObjectId(), Owner: "JasonHe", Message: "Hello from Jason He.", Timestamp: time.Now()}, 
		model.Tweet{ID: bson.NewObjectId(), Owner: "DianeLin", Message: "Hi, I am Diane Lin. Weather sucks.", Timestamp: time.Now()}, 
		model.Tweet{ID: bson.NewObjectId(), Owner: "DianeLin", Message: "Hello from Diane Lin.", Timestamp: time.Now()},
		model.Tweet{ID: bson.NewObjectId(), Owner: "JasonHo", Message: "Hi, I am Jason Ho. Weather sucks.", Timestamp: time.Now()}, 
		model.Tweet{ID: bson.NewObjectId(), Owner: "JasonHo", Message: "Hello from Jason Ho.", Timestamp: time.Now()}, 
		model.Tweet{ID: bson.NewObjectId(), Owner: "JasonHo", Message: "Hello world!", Timestamp: time.Now()}, 
		model.Tweet{ID: bson.NewObjectId(), Owner: "JasonHo", Message: "Harvest Moon: All you need to know It's the full moon closest to Sept equinox - coming up Oct 5.", Timestamp: time.Now()}, 
		model.Tweet{ID: bson.NewObjectId(), Owner: "JasonHo", Message: "Hey yo!!", Timestamp: time.Now()}, 
		model.Tweet{ID: bson.NewObjectId(), Owner: "MarsLee", Message: "Hi, I am Chih-Yin Lee. Weather sucks.", Timestamp: time.Now()}, 
		model.Tweet{ID: bson.NewObjectId(), Owner: "MarsLee", Message: "Hello from Chih-Yin Lee.", Timestamp: time.Now()}, 
		model.Tweet{ID: bson.NewObjectId(), Owner: "JasonHe", Message: "Hi, I am Jason He. Weather sucks.", Timestamp: time.Now()}, 
		model.Tweet{ID: bson.NewObjectId(), Owner: "JasonHe", Message: "Hello from Jason He.", Timestamp: time.Now()}, 
		model.Tweet{ID: bson.NewObjectId(), Owner: "DianeLin", Message: "Hi, I am Diane Lin. Weather sucks.", Timestamp: time.Now()}, 
		model.Tweet{ID: bson.NewObjectId(), Owner: "DianeLin", Message: "Hello from Diane Lin.", Timestamp: time.Now()},
	}

	for _, t := range tweets {
		err := tweetC.Insert(t)
		if err != nil {
			panic(err)
		}
	}

	h.DB.Close()
}

func reconstructTestDB() {
	h := handler.NewHandler("mongodb://SEavenger:SEavenger@ds121225.mlab.com:21225/se_avengers_test")

	err := h.DB.DB(h.DBName).DropDatabase()
	if err != nil {
		panic(err)
	}

	userC := h.DB.DB(h.DBName).C(model.UserCollection)

	users := []model.User {
		model.User{ID: bson.NewObjectId(), Username: "testSignup", FirstName: "test", LastName: "signup", Password: "test", Email: "testSignup@gmail.com", Followers: []string{}, Following: []string{}, 
			Bio: "testtest", Token: ""},
		model.User{ID: bson.NewObjectId(), Username: "testSignup_dup", FirstName: "test", LastName: "signup_dup", Password: "test", Email: "testSignup_dup@gmail.com", Followers: []string{}, Following: []string{}, 
			Bio: "testtest", Token: ""},
		model.User{ID: bson.NewObjectId(), Username: "testLogin", FirstName: "testLogin", Password: "test", Email: "testLogin@gmail.com", Followers: []string{}, Following: []string{}},
		model.User{ID: bson.NewObjectId(), Username: "testUserInfo_1", FirstName: "testUserInfo_1", Password: "test", Email: "testUserInfo_1@gmail.com", Followers: []string{}, Following: []string{"testUserInfo_2"}},
		model.User{ID: bson.NewObjectId(), Username: "testUserInfo_2", FirstName: "testUserInfo_2", Password: "test", Email: "testUserInfo_2@gmail.com", Followers: []string{"testUserInfo_1"}, Following: []string{}},
		model.User{ID: bson.NewObjectId(), Username: "testShowFollower_1", FirstName: "testShowFollower_1", Password: "test", Email: "testShowFollower_1@gmail.com", Followers: []string{}, Following: []string{}},
		model.User{ID: bson.NewObjectId(), Username: "testShowFollower_2", FirstName: "test", LastName: "ShowFollower_2", Password: "test", Email: "testShowFollower_2@gmail.com",
			Followers: []string{"testShowFollower_1", "testShowFollower_3"}, Following: []string{}, Bio: "testtest"},
		model.User{ID: bson.NewObjectId(), Username: "testShowFollower_3", FirstName: "testShowFollower_3", Password: "test", Email: "testShowFollower_3@gmail.com",
			Followers: []string{"testShowFollower_2"}, Following: []string{}},
		model.User{ID: bson.NewObjectId(), Username: "testShowFollowing_1", FirstName: "testShowFollowing_1", Password: "test", Email: "testShowFollowing_1@gmail.com", Followers: []string{}, Following: []string{}},
		model.User{ID: bson.NewObjectId(), Username: "testShowFollowing_2", FirstName: "test", LastName: "ShowFollowing_2", Password: "test", Email: "testShowFollowing_2@gmail.com",
			Followers: []string{}, Following: []string{"testShowFollowing_1", "testShowFollowing_3"}, Bio: "testtest"},
		model.User{ID: bson.NewObjectId(), Username: "testShowFollowing_3", FirstName: "testShowFollowing_3", Password: "test", Email: "testShowFollowing_3@gmail.com",
			Followers: []string{}, Following: []string{"testShowFollowing_2"}},
		model.User{ID: bson.NewObjectId(), Username: "testFollow", FirstName: "testFollow", Password: "test", Email: "testFollow@gmail.com", Followers: []string{}, Following: []string{"testFollow_1"}},
		model.User{ID: bson.NewObjectId(), Username: "testFollow_1", FirstName: "testFollow_1", Password: "test", Email: "testFollow_1@gmail.com", Followers: []string{}, Following: []string{}},
		model.User{ID: bson.NewObjectId(), Username: "testFollow_2", FirstName: "testFollow_2", Password: "test", Email: "testFollow_2@gmail.com", Followers: []string{}, Following: []string{}},
		model.User{ID: bson.NewObjectId(), Username: "testUnfollow", FirstName: "testUnfollow", Password: "test", Email: "testUnfollow@gmail.com", Followers: []string{}, Following: []string{"testUnfollow_1", "testUnfollow_2"}},
		model.User{ID: bson.NewObjectId(), Username: "testUnfollow_1", FirstName: "testUnfollow_1", Password: "test", Email: "testUnfollow_1@gmail.com", Followers: []string{}, Following: []string{}},
		model.User{ID: bson.NewObjectId(), Username: "testUnfollow_2", FirstName: "testUnfollow_2", Password: "test", Email: "testUnfollow_2@gmail.com", Followers: []string{"testUnfollow"}, Following: []string{}},
		model.User{ID: bson.NewObjectId(), Username: "testUpdateUserInfo", FirstName: "testUpdateUserInfo", Password: "test", Email: "testUpdateUserInfo@gmail.com", Followers: []string{}, Following: []string{}},
		model.User{ID: bson.NewObjectId(), Username: "testUpdateUserInfo_empty_firstname", FirstName: "testUpdateUserInfo_empty_firstname", Password: "test", Email: "testUpdateUserInfo_empty_firstname@gmail.com", 
			Followers: []string{}, Following: []string{}, Tag: "testUpdate"},
		model.User{ID: bson.NewObjectId(), Username: "testUpdateProfilePicture", FirstName: "testUpdateProfilePicture", Password: "test", Email: "testUpdateProfilePicture@gmail.com", Followers: []string{}, Following: []string{}},

		model.User{ID: bson.NewObjectId(), Username: "testFlushNotif", FirstName: "testFlushNotif", Password: "test", Email: "testFlushNotif@gmail.com", Followers: []string{}, Following: []string{}},
		model.User{ID: bson.NewObjectId(), Username: "testForwardNewTweetNotif_1", FirstName: "testForwardNewTweetNotif", Password: "test", Email: "testForwardNewTweetNotif@gmail.com", Followers: []string{}, Following: []string{"xxx"}},
		model.User{ID: bson.NewObjectId(), Username: "testForwardNewTweetNotif_2", FirstName: "testForwardNewTweetNotif", Password: "test", Email: "testForwardNewTweetNotif@gmail.com", Followers: []string{}, Following: []string{"xxx"}},
		model.User{ID: bson.NewObjectId(), Username: "testForwardFollowNotif", FirstName: "testForwardFollowNotif", Password: "test", Email: "testForwardFollowNotif@gmail.com", Followers: []string{}, Following: []string{}},
		model.User{ID: bson.NewObjectId(), Username: "testClearNotif", FirstName: "testClearNotif", Password: "test", Email: "testClearNotif@gmail.com", Followers: []string{}, Following: []string{}},
		// User for tweet testing
		model.User{ID: bson.NewObjectId(), Username: "testNewTweet", FirstName: "testNewTweet", Password: "test", Email: "testNewTweet@gmail.com", Followers: []string{}, Following: []string{}},
		model.User{ID: bson.NewObjectId(), Username: "testFetchOenTweetsSuccess", FirstName: "testFetchOenTweetsSuccess", Password: "test", Email: "testFetchOenTweetsSuccess@gmail.com", Followers: []string{}, Following: []string{}},
		model.User{ID: bson.NewObjectId(), Username: "testFetchOenTweetsOverPage", FirstName: "testFetchOenTweetsOverPage", Password: "test", Email: "testFetchOenTweetsOverPage@gmail.com", Followers: []string{}, Following: []string{}},
		model.User{ID: bson.NewObjectId(), Username: "testFetchOenTweetsWithNoTweet1", FirstName: "testFetchOenTweetsWithNoTweet1", Password: "test", Email: "testFetchOenTweetsWithNoTweet1@gmail.com", Followers: []string{}, Following: []string{}},
		model.User{ID: bson.NewObjectId(), Username: "testFetchOenTweetsWithNoTweet2", FirstName: "testFetchOenTweetsWithNoTweet2", Password: "test", Email: "testFetchOenTweetsWithNoTweet2@gmail.com", Followers: []string{}, Following: []string{}},
		model.User{ID: bson.NewObjectId(), Username: "TestFetchTweetTimeLineSuccess", FirstName: "TestFetchTweetTimeLineSuccess", Password: "test", Email: "TestFetchTweetTimeLineSuccess@gmail.com", Followers: []string{}, Following: []string{"TestFetchTweetTimeLineFollowing1", "TestFetchTweetTimeLineFollowing2"}},
		model.User{ID: bson.NewObjectId(), Username: "TestFetchTweetTimeLineOverPage", FirstName: "TestFetchTweetTimeLineOverPage", Password: "test", Email: "TestFetchTweetTimeLineOverPage@gmail.com", Followers: []string{}, Following: []string{}},
		model.User{ID: bson.NewObjectId(), Username: "TestFetchTweetTimeLineNoTweet1", FirstName: "TestFetchTweetTimeLineNoTweet1", Password: "test", Email: "TestFetchTweetTimeLineNoTweet1@gmail.com", Followers: []string{}, Following: []string{}},
		model.User{ID: bson.NewObjectId(), Username: "TestFetchTweetTimeLineNoTweet2", FirstName: "TestFetchTweetTimeLineNoTweet2", Password: "test", Email: "TestFetchTweetTimeLineNoTweet2@gmail.com", Followers: []string{}, Following: []string{}},
		model.User{ID: bson.NewObjectId(), Username: "TestFetchTweetTimeLineFollowing1", FirstName: "TestFetchTweetTimeLineFollowing1", Password: "test", Email: "TestFetchTweetTimeLineFollowing1@gmail.com", Followers: []string{}, Following: []string{}},
		model.User{ID: bson.NewObjectId(), Username: "TestFetchTweetTimeLineFollowing2", FirstName: "TestFetchTweetTimeLineFollowing2", Password: "test", Email: "TestFetchTweetTimeLineFollowing2@gmail.com", Followers: []string{}, Following: []string{}},
		// User for comment testing
		model.User{ID: bson.NewObjectId(), Username: "TestNewCommentSuccess", FirstName: "TestNewCommentSuccess", Password: "test", Email: "TestNewCommentSuccess@gmail.com", Followers: []string{}, Following: []string{}},
		model.User{ID: bson.NewObjectId(), Username: "TestNewCommentEmpty", FirstName: "TestNewCommentEmpty", Password: "test", Email: "TestNewCommentEmpty@gmail.com", Followers: []string{}, Following: []string{}},
		model.User{ID: bson.NewObjectId(), Username: "TestNewCommentInvalidTweetID", FirstName: "TestNewCommentInvalidTweetID", Password: "test", Email: "TestNewCommentInvalidTweetID@gmail.com", Followers: []string{}, Following: []string{}},
		model.User{ID: bson.NewObjectId(), Username: "TestFetchCommentSuccess", FirstName: "TestFetchCommentSuccess", Password: "test", Email: "TestFetchCommentSuccess@gmail.com", Followers: []string{}, Following: []string{}},
		model.User{ID: bson.NewObjectId(), Username: "TestFetchCommentSuccess", FirstName: "TestFetchCommentInvalidTweetID", Password: "test", Email: "TestFetchCommentSuccess@gmail.com", Followers: []string{}, Following: []string{}},
	}

	for _, u := range users {
		err := userC.Insert(u)
		if err != nil {
			panic(err)
		}
	}

	tweetC := h.DB.DB(h.DBName).C(model.TweetCollection)
	
		tweets := []model.Tweet {
			// testFetchOenTweetsSuccess
			model.Tweet{ID: bson.NewObjectId(), Owner: "testFetchOenTweetsSuccess", Message: "testFetchOenTweetsSuccess, testFetchOenTweetsSuccess, testFetchOenTweetsSuccess, testFetchOenTweetsSuccess", Timestamp: time.Now()}, 
			model.Tweet{ID: bson.NewObjectId(), Owner: "testFetchOenTweetsSuccess", Message: "testFetchOenTweetsSuccess, testFetchOenTweetsSuccess, testFetchOenTweetsSuccess, testFetchOenTweetsSuccess", Timestamp: time.Now()}, 
			model.Tweet{ID: bson.NewObjectId(), Owner: "testFetchOenTweetsSuccess", Message: "testFetchOenTweetsSuccess, testFetchOenTweetsSuccess, testFetchOenTweetsSuccess, testFetchOenTweetsSuccess", Timestamp: time.Now()}, 
			model.Tweet{ID: bson.NewObjectId(), Owner: "testFetchOenTweetsSuccess", Message: "testFetchOenTweetsSuccess, testFetchOenTweetsSuccess, testFetchOenTweetsSuccess, testFetchOenTweetsSuccess", Timestamp: time.Now()}, 
			model.Tweet{ID: bson.NewObjectId(), Owner: "testFetchOenTweetsSuccess", Message: "testFetchOenTweetsSuccess, testFetchOenTweetsSuccess, testFetchOenTweetsSuccess, testFetchOenTweetsSuccess", Timestamp: time.Now()}, 
			model.Tweet{ID: bson.NewObjectId(), Owner: "testFetchOenTweetsSuccess", Message: "testFetchOenTweetsSuccess, testFetchOenTweetsSuccess, testFetchOenTweetsSuccess, testFetchOenTweetsSuccess", Timestamp: time.Now()}, 
			model.Tweet{ID: bson.NewObjectId(), Owner: "testFetchOenTweetsSuccess", Message: "testFetchOenTweetsSuccess, testFetchOenTweetsSuccess, testFetchOenTweetsSuccess, testFetchOenTweetsSuccess", Timestamp: time.Now()}, 
			model.Tweet{ID: bson.NewObjectId(), Owner: "testFetchOenTweetsSuccess", Message: "testFetchOenTweetsSuccess, testFetchOenTweetsSuccess, testFetchOenTweetsSuccess, testFetchOenTweetsSuccess", Timestamp: time.Now()}, 
			model.Tweet{ID: bson.NewObjectId(), Owner: "testFetchOenTweetsSuccess", Message: "testFetchOenTweetsSuccess, testFetchOenTweetsSuccess, testFetchOenTweetsSuccess, testFetchOenTweetsSuccess", Timestamp: time.Now()}, 
			model.Tweet{ID: bson.NewObjectId(), Owner: "testFetchOenTweetsSuccess", Message: "testFetchOenTweetsSuccess, testFetchOenTweetsSuccess, testFetchOenTweetsSuccess, testFetchOenTweetsSuccess", Timestamp: time.Now()}, 
			// testFetchOenTweetsOverPage
			model.Tweet{ID: bson.NewObjectId(), Owner: "testFetchOenTweetsOverPage", Message: "testFetchOenTweetsSuccess, testFetchOenTweetsSuccess, testFetchOenTweetsSuccess, testFetchOenTweetsSuccess", Timestamp: time.Now()}, 
			// TestFetchTweetTimeLine
			model.Tweet{ID: bson.NewObjectId(), Owner: "TestFetchTweetTimeLineFollowing1", Message: "testFetchOenTweetsSuccess, testFetchOenTweetsSuccess, testFetchOenTweetsSuccess, testFetchOenTweetsSuccess", Timestamp: time.Now()}, 
			model.Tweet{ID: bson.NewObjectId(), Owner: "TestFetchTweetTimeLineFollowing1", Message: "testFetchOenTweetsSuccess, testFetchOenTweetsSuccess, testFetchOenTweetsSuccess, testFetchOenTweetsSuccess", Timestamp: time.Now()}, 
			model.Tweet{ID: bson.NewObjectId(), Owner: "TestFetchTweetTimeLineFollowing1", Message: "testFetchOenTweetsSuccess, testFetchOenTweetsSuccess, testFetchOenTweetsSuccess, testFetchOenTweetsSuccess", Timestamp: time.Now()}, 
			model.Tweet{ID: bson.NewObjectId(), Owner: "TestFetchTweetTimeLineFollowing1", Message: "testFetchOenTweetsSuccess, testFetchOenTweetsSuccess, testFetchOenTweetsSuccess, testFetchOenTweetsSuccess", Timestamp: time.Now()}, 
			model.Tweet{ID: bson.NewObjectId(), Owner: "TestFetchTweetTimeLineFollowing2", Message: "testFetchOenTweetsSuccess, testFetchOenTweetsSuccess, testFetchOenTweetsSuccess, testFetchOenTweetsSuccess", Timestamp: time.Now()}, 
			model.Tweet{ID: bson.NewObjectId(), Owner: "TestFetchTweetTimeLineFollowing2", Message: "testFetchOenTweetsSuccess, testFetchOenTweetsSuccess, testFetchOenTweetsSuccess, testFetchOenTweetsSuccess", Timestamp: time.Now()}, 
			model.Tweet{ID: bson.NewObjectId(), Owner: "TestFetchTweetTimeLineFollowing2", Message: "testFetchOenTweetsSuccess, testFetchOenTweetsSuccess, testFetchOenTweetsSuccess, testFetchOenTweetsSuccess", Timestamp: time.Now()}, 
			model.Tweet{ID: bson.NewObjectId(), Owner: "TestFetchTweetTimeLineFollowing2", Message: "testFetchOenTweetsSuccess, testFetchOenTweetsSuccess, testFetchOenTweetsSuccess, testFetchOenTweetsSuccess", Timestamp: time.Now()}, 
			// TestFetchComment
			model.Tweet{ID: bson.NewObjectId(), Owner: "TestFetchCommentSuccess", Message: "testFetchOenTweetsSuccess, testFetchOenTweetsSuccess, testFetchOenTweetsSuccess, testFetchOenTweetsSuccess", Timestamp: time.Now()}, 
		}
		for _, t := range tweets {
			err := tweetC.Insert(t)
			if err != nil {
				panic(err)
			}
		}

		CommentC := h.DB.DB(h.DBName).C(model.CommentCollection)
		tempTweet := []model.Tweet{}
		tweetC.Find(nil).Sort("timestamp").All(&tempTweet)
		fromTweetIDForTesting := tempTweet[0].ID.Hex()
			comments := []model.Comment {
				// testFetchOenTweetsSuccess
				model.Comment{ID: bson.NewObjectId(), FromTweetID: fromTweetIDForTesting, FromUsername: "TestFetchCommentSuccess", Message: "testFetchOenTweetsSuccess, testFetchOenTweetsSuccess, testFetchOenTweetsSuccess, testFetchOenTweetsSuccess", Timestamp: time.Now()}, 
				model.Comment{ID: bson.NewObjectId(), FromTweetID: fromTweetIDForTesting, FromUsername: "TestFetchCommentSuccess", Message: "testFetchOenTweetsSuccess, testFetchOenTweetsSuccess, testFetchOenTweetsSuccess, testFetchOenTweetsSuccess", Timestamp: time.Now()}, 
				model.Comment{ID: bson.NewObjectId(), FromTweetID: fromTweetIDForTesting, FromUsername: "TestFetchCommentSuccess", Message: "testFetchOenTweetsSuccess, testFetchOenTweetsSuccess, testFetchOenTweetsSuccess, testFetchOenTweetsSuccess", Timestamp: time.Now()}, 
				model.Comment{ID: bson.NewObjectId(), FromTweetID: fromTweetIDForTesting, FromUsername: "TestFetchCommentSuccess", Message: "testFetchOenTweetsSuccess, testFetchOenTweetsSuccess, testFetchOenTweetsSuccess, testFetchOenTweetsSuccess", Timestamp: time.Now()}, 
				model.Comment{ID: bson.NewObjectId(), FromTweetID: fromTweetIDForTesting, FromUsername: "TestFetchCommentSuccess", Message: "testFetchOenTweetsSuccess, testFetchOenTweetsSuccess, testFetchOenTweetsSuccess, testFetchOenTweetsSuccess", Timestamp: time.Now()}, 
				model.Comment{ID: bson.NewObjectId(), FromTweetID: fromTweetIDForTesting, FromUsername: "TestFetchCommentSuccess", Message: "testFetchOenTweetsSuccess, testFetchOenTweetsSuccess, testFetchOenTweetsSuccess, testFetchOenTweetsSuccess", Timestamp: time.Now()}, 
				model.Comment{ID: bson.NewObjectId(), FromTweetID: fromTweetIDForTesting, FromUsername: "TestFetchCommentSuccess", Message: "testFetchOenTweetsSuccess, testFetchOenTweetsSuccess, testFetchOenTweetsSuccess, testFetchOenTweetsSuccess", Timestamp: time.Now()}, 
				model.Comment{ID: bson.NewObjectId(), FromTweetID: fromTweetIDForTesting, FromUsername: "TestFetchCommentSuccess", Message: "testFetchOenTweetsSuccess, testFetchOenTweetsSuccess, testFetchOenTweetsSuccess, testFetchOenTweetsSuccess", Timestamp: time.Now()}, 
				model.Comment{ID: bson.NewObjectId(), FromTweetID: fromTweetIDForTesting, FromUsername: "TestFetchCommentSuccess", Message: "testFetchOenTweetsSuccess, testFetchOenTweetsSuccess, testFetchOenTweetsSuccess, testFetchOenTweetsSuccess", Timestamp: time.Now()}, 
				model.Comment{ID: bson.NewObjectId(), FromTweetID: fromTweetIDForTesting, FromUsername: "TestFetchCommentSuccess", Message: "testFetchOenTweetsSuccess, testFetchOenTweetsSuccess, testFetchOenTweetsSuccess, testFetchOenTweetsSuccess", Timestamp: time.Now()}, 
				model.Comment{ID: bson.NewObjectId(), FromTweetID: fromTweetIDForTesting, FromUsername: "TestFetchCommentSuccess", Message: "testFetchOenTweetsSuccess, testFetchOenTweetsSuccess, testFetchOenTweetsSuccess, testFetchOenTweetsSuccess", Timestamp: time.Now()}, 
				model.Comment{ID: bson.NewObjectId(), FromTweetID: fromTweetIDForTesting, FromUsername: "TestFetchCommentSuccess", Message: "testFetchOenTweetsSuccess, testFetchOenTweetsSuccess, testFetchOenTweetsSuccess, testFetchOenTweetsSuccess", Timestamp: time.Now()}, 
			}
			for _, t := range comments {
				err := CommentC.Insert(t)
				if err != nil {
					panic(err)
				}
			}

	notificationC := h.DB.DB(h.DBName).C(model.NotificationCollection)

	notifications := []model.Individual {
		model.Individual{ID: bson.NewObjectId(), Username: "testSignup", Notifications: make([]model.Notification, 0)},
		model.Individual{ID: bson.NewObjectId(), Username: "testSignup_dup", Notifications: make([]model.Notification, 0)},
		model.Individual{ID: bson.NewObjectId(), Username: "testLogin", Notifications: make([]model.Notification, 0)},
		model.Individual{ID: bson.NewObjectId(), Username: "testUserInfo_1", Notifications: make([]model.Notification, 0)},
		model.Individual{ID: bson.NewObjectId(), Username: "testUserInfo_2", Notifications: make([]model.Notification, 0)},
		model.Individual{ID: bson.NewObjectId(), Username: "testShowFollower_1", Notifications: make([]model.Notification, 0)},
		model.Individual{ID: bson.NewObjectId(), Username: "testShowFollower_2", Notifications: make([]model.Notification, 0)},
		model.Individual{ID: bson.NewObjectId(), Username: "testShowFollower_3", Notifications: make([]model.Notification, 0)},
		model.Individual{ID: bson.NewObjectId(), Username: "testShowFollowing_1", Notifications: make([]model.Notification, 0)},
		model.Individual{ID: bson.NewObjectId(), Username: "testShowFollowing_2", Notifications: make([]model.Notification, 0)},
		model.Individual{ID: bson.NewObjectId(), Username: "testShowFollowing_3", Notifications: make([]model.Notification, 0)},
		model.Individual{ID: bson.NewObjectId(), Username: "testFollow", Notifications: make([]model.Notification, 0)},
		model.Individual{ID: bson.NewObjectId(), Username: "testFollow_1", Notifications: make([]model.Notification, 0)},
		model.Individual{ID: bson.NewObjectId(), Username: "testFollow_2", Notifications: make([]model.Notification, 0)},
		model.Individual{ID: bson.NewObjectId(), Username: "testUnfollow", Notifications: make([]model.Notification, 0)},
		model.Individual{ID: bson.NewObjectId(), Username: "testUnfollow_1", Notifications: make([]model.Notification, 0)},
		model.Individual{ID: bson.NewObjectId(), Username: "testUnfollow_2", Notifications: make([]model.Notification, 0)},
		model.Individual{ID: bson.NewObjectId(), Username: "testUpdateUserInfo", Notifications: make([]model.Notification, 0)},
		model.Individual{ID: bson.NewObjectId(), Username: "testUpdateUserInfo_empty_firstname", Notifications: make([]model.Notification, 0)},
		model.Individual{ID: bson.NewObjectId(), Username: "testUpdateProfilePicture", Notifications: make([]model.Notification, 0)},
		model.Individual{ID: bson.NewObjectId(), Username: "testForwardNewTweetNotif_1", Notifications: make([]model.Notification, 0)},
		model.Individual{ID: bson.NewObjectId(), Username: "testForwardNewTweetNotif_2", Notifications: make([]model.Notification, 0)},
		model.Individual{ID: bson.NewObjectId(), Username: "testForwardFollowNotif", Notifications: make([]model.Notification, 0)},
	}

	const timestampForm = "Jan 2, 2006 at 3:04pm (MST)"
	t := time.Time{}
	ns := []model.Notification{}
	
	t, _ = time.Parse(timestampForm, "Nov 2, 2017 at 3:00pm (MST)")
	ns = append(ns, model.Notification{Timestamp: t, Type: model.FollowType, Detail: model.FollowNotif{Followee: "testFlushNotif", Follower: "AAA"}})
	t, _ = time.Parse(timestampForm, "Nov 3, 2017 at 2:00pm (MST)")
	ns = append(ns, model.Notification{Timestamp: t, Type: model.FollowType, Detail: model.FollowNotif{Followee: "testFlushNotif", Follower: "BBB"}})
	t, _ = time.Parse(timestampForm, "Dec 25, 2017 at 2:00pm (MST)")
	ns = append(ns, model.Notification{Timestamp: t, Type: model.FollowType, Detail: model.FollowNotif{Followee: "testFlushNotif", Follower: "CCC"}})
	notifications = append(notifications, model.Individual{ID: bson.NewObjectId(), Username: "testFlushNotif", Notifications: ns})
	ns = []model.Notification{}

	t, _ = time.Parse(timestampForm, "Nov 2, 2017 at 3:00pm (MST)")
	ns = append(ns, model.Notification{Timestamp: t, Type: model.FollowType, Detail: model.FollowNotif{}})
	t, _ = time.Parse(timestampForm, "Nov 3, 2017 at 2:00pm (MST)")
	ns = append(ns, model.Notification{Timestamp: t, Type: model.FollowType, Detail: model.FollowNotif{}})
	t, _ = time.Parse(timestampForm, "Dec 25, 2017 at 2:00pm (MST)")
	ns = append(ns, model.Notification{Timestamp: t, Type: model.FollowType, Detail: model.FollowNotif{}})
	notifications = append(notifications, model.Individual{ID: bson.NewObjectId(), Username: "testClearNotif", Notifications: ns})
	ns = []model.Notification{}

	for _, c := range notifications {
		err := notificationC.Insert(c)
		if err != nil {
			panic(err)
		}
	}

	h.DB.Close()
}