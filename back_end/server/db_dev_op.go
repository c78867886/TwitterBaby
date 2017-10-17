package server

import (
	"fmt"
	"time"
	"gopkg.in/mgo.v2/bson"
	"gopkg.in/mgo.v2"
	"model"
)

func dbDrop(db *mgo.Session) {
	err := db.DB("se_avengers").DropDatabase()
	if err != nil {
		panic(err)
	}

	db.Close()
}

func dbReinsert(db *mgo.Session) {
	dbDrop(db.Clone())

	userC := db.DB("se_avengers").C("users")

	users := []model.User {
		model.User{Username: "JasonHo", FirstName: "Jason", LastName: "Ho", Password: "test1", Email: "hojason117@gmail.com", Followers: []string{}, Following: []string{}, 
			Bio: "Hi everyone, this is Jason Ho.", Token: ""},
		model.User{Username: "MarsLee", FirstName: "Chih-Yin", LastName: "Lee", Password: "test2", Email: "c788678867886@gmail.com", Followers: []string{}, Following: []string{}, 
			Bio: "Hi everyone, this is Mars Lee.", Token: ""},
		model.User{Username: "JasonHe", FirstName: "Jason", LastName: "He", Password: "test3", Email: "hexing_h@hotmail.com", Followers: []string{}, Following: []string{}, 
			Bio: "Hi everyone, this is Jason He.", Token: ""},
		model.User{Username: "DianeLin", FirstName: "Diane", LastName: "Lin", Password: "test4", Email: "diane@gmail.com", Followers: []string{}, Following: []string{}, 
			Bio: "Hi everyone, this is Diane Lin.", Token: ""}, 
		model.User{Username: "TomRiddle", FirstName: "Tom", LastName: "Riddle", Password: "test5", Email: "triddle@gmail.com", Followers: []string{}, Following: []string{}, 
			Bio: "Hi everyone, this is Lord Voldemort.", Token: ""}, 
		model.User{Username: "JS", FirstName: "Jon", Password: "pass", Email: "json@gmail.com", Followers: []string{}, Following: []string{}, 
			Bio: "Hi everyone, this is Lord Voldemort.", Token: ""}}
	
	userIDs := []bson.ObjectId{}

	for _, u:= range users { 
		u.ID = bson.NewObjectId()
		userIDs = append(userIDs, u.ID)
		err := userC.Insert(u)
		if err != nil {
			panic(err)
		}
	}

	tweetC := db.DB("se_avengers").C("tweets")

	tweets := []model.Tweet {
		model.Tweet{ID: bson.NewObjectId(), Owner: "JasonHo", Message: "Hi, I am Jason Ho. Weather sucks.", Timestamp: time.Now()}, 
		model.Tweet{ID: bson.NewObjectId(), Owner: "JasonHo", Message: "Hello from Jason Ho.", Timestamp: time.Now()}, 
		model.Tweet{ID: bson.NewObjectId(), Owner: userIDs[0].Hex(), Message: "Hello world!", Timestamp: time.Now()}, 
		model.Tweet{ID: bson.NewObjectId(), Owner: userIDs[0].Hex(), Message: "Harvest Moon: All you need to know It's the full moon closest to Sept equinox - coming up Oct 5.", Timestamp: time.Now()}, 
		model.Tweet{ID: bson.NewObjectId(), Owner: userIDs[0].Hex(), Message: "Hey yo!!", Timestamp: time.Now()}, 
		model.Tweet{ID: bson.NewObjectId(), Owner: userIDs[1].Hex(), Message: "Hi, I am Chih-Yin Lee. Weather sucks.", Timestamp: time.Now()}, 
		model.Tweet{ID: bson.NewObjectId(), Owner: userIDs[1].Hex(), Message: "Hello from Chih-Yin Lee.", Timestamp: time.Now()}, 
		model.Tweet{ID: bson.NewObjectId(), Owner: userIDs[2].Hex(), Message: "Hi, I am Jason He. Weather sucks.", Timestamp: time.Now()}, 
		model.Tweet{ID: bson.NewObjectId(), Owner: userIDs[2].Hex(), Message: "Hello from Jason He.", Timestamp: time.Now()}, 
		model.Tweet{ID: bson.NewObjectId(), Owner: userIDs[3].Hex(), Message: "Hi, I am Diane Lin. Weather sucks.", Timestamp: time.Now()}, 
		model.Tweet{ID: bson.NewObjectId(), Owner: userIDs[3].Hex(), Message: "Hello from Diane Lin.", Timestamp: time.Now()},
		model.Tweet{ID: bson.NewObjectId(), Owner: "JasonHo", Message: "Hi, I am Jason Ho. Weather sucks.", Timestamp: time.Now()}, 
		model.Tweet{ID: bson.NewObjectId(), Owner: "JasonHo", Message: "Hello from Jason Ho.", Timestamp: time.Now()}, 
		model.Tweet{ID: bson.NewObjectId(), Owner: userIDs[0].Hex(), Message: "Hello world!", Timestamp: time.Now()}, 
		model.Tweet{ID: bson.NewObjectId(), Owner: userIDs[0].Hex(), Message: "Harvest Moon: All you need to know It's the full moon closest to Sept equinox - coming up Oct 5.", Timestamp: time.Now()}, 
		model.Tweet{ID: bson.NewObjectId(), Owner: userIDs[0].Hex(), Message: "Hey yo!!", Timestamp: time.Now()}, 
		model.Tweet{ID: bson.NewObjectId(), Owner: userIDs[1].Hex(), Message: "Hi, I am Chih-Yin Lee. Weather sucks.", Timestamp: time.Now()}, 
		model.Tweet{ID: bson.NewObjectId(), Owner: userIDs[1].Hex(), Message: "Hello from Chih-Yin Lee.", Timestamp: time.Now()}, 
		model.Tweet{ID: bson.NewObjectId(), Owner: userIDs[2].Hex(), Message: "Hi, I am Jason He. Weather sucks.", Timestamp: time.Now()}, 
		model.Tweet{ID: bson.NewObjectId(), Owner: userIDs[2].Hex(), Message: "Hello from Jason He.", Timestamp: time.Now()}, 
		model.Tweet{ID: bson.NewObjectId(), Owner: userIDs[3].Hex(), Message: "Hi, I am Diane Lin. Weather sucks.", Timestamp: time.Now()}, 
		model.Tweet{ID: bson.NewObjectId(), Owner: userIDs[3].Hex(), Message: "Hello from Diane Lin.", Timestamp: time.Now()},
		model.Tweet{ID: bson.NewObjectId(), Owner: "JasonHo", Message: "Hi, I am Jason Ho. Weather sucks.", Timestamp: time.Now()}, 
		model.Tweet{ID: bson.NewObjectId(), Owner: "JasonHo", Message: "Hello from Jason Ho.", Timestamp: time.Now()}, 
		model.Tweet{ID: bson.NewObjectId(), Owner: userIDs[0].Hex(), Message: "Hello world!", Timestamp: time.Now()}, 
		model.Tweet{ID: bson.NewObjectId(), Owner: userIDs[0].Hex(), Message: "Harvest Moon: All you need to know It's the full moon closest to Sept equinox - coming up Oct 5.", Timestamp: time.Now()}, 
		model.Tweet{ID: bson.NewObjectId(), Owner: userIDs[0].Hex(), Message: "Hey yo!!", Timestamp: time.Now()}, 
		model.Tweet{ID: bson.NewObjectId(), Owner: userIDs[1].Hex(), Message: "Hi, I am Chih-Yin Lee. Weather sucks.", Timestamp: time.Now()}, 
		model.Tweet{ID: bson.NewObjectId(), Owner: userIDs[1].Hex(), Message: "Hello from Chih-Yin Lee.", Timestamp: time.Now()}, 
		model.Tweet{ID: bson.NewObjectId(), Owner: userIDs[2].Hex(), Message: "Hi, I am Jason He. Weather sucks.", Timestamp: time.Now()}, 
		model.Tweet{ID: bson.NewObjectId(), Owner: userIDs[2].Hex(), Message: "Hello from Jason He.", Timestamp: time.Now()}, 
		model.Tweet{ID: bson.NewObjectId(), Owner: userIDs[3].Hex(), Message: "Hi, I am Diane Lin. Weather sucks.", Timestamp: time.Now()}, 
		model.Tweet{ID: bson.NewObjectId(), Owner: userIDs[3].Hex(), Message: "Hello from Diane Lin.", Timestamp: time.Now()},
		model.Tweet{ID: bson.NewObjectId(), Owner: "JasonHo", Message: "Hi, I am Jason Ho. Weather sucks.", Timestamp: time.Now()}, 
		model.Tweet{ID: bson.NewObjectId(), Owner: "JasonHo", Message: "Hello from Jason Ho.", Timestamp: time.Now()}, 
		model.Tweet{ID: bson.NewObjectId(), Owner: userIDs[0].Hex(), Message: "Hello world!", Timestamp: time.Now()}, 
		model.Tweet{ID: bson.NewObjectId(), Owner: userIDs[0].Hex(), Message: "Harvest Moon: All you need to know It's the full moon closest to Sept equinox - coming up Oct 5.", Timestamp: time.Now()}, 
		model.Tweet{ID: bson.NewObjectId(), Owner: userIDs[0].Hex(), Message: "Hey yo!!", Timestamp: time.Now()}, 
		model.Tweet{ID: bson.NewObjectId(), Owner: userIDs[1].Hex(), Message: "Hi, I am Chih-Yin Lee. Weather sucks.", Timestamp: time.Now()}, 
		model.Tweet{ID: bson.NewObjectId(), Owner: userIDs[1].Hex(), Message: "Hello from Chih-Yin Lee.", Timestamp: time.Now()}, 
		model.Tweet{ID: bson.NewObjectId(), Owner: userIDs[2].Hex(), Message: "Hi, I am Jason He. Weather sucks.", Timestamp: time.Now()}, 
		model.Tweet{ID: bson.NewObjectId(), Owner: userIDs[2].Hex(), Message: "Hello from Jason He.", Timestamp: time.Now()}, 
		model.Tweet{ID: bson.NewObjectId(), Owner: userIDs[3].Hex(), Message: "Hi, I am Diane Lin. Weather sucks.", Timestamp: time.Now()}, 
		model.Tweet{ID: bson.NewObjectId(), Owner: userIDs[3].Hex(), Message: "Hello from Diane Lin.", Timestamp: time.Now()},
		model.Tweet{ID: bson.NewObjectId(), Owner: "JasonHo", Message: "Hi, I am Jason Ho. Weather sucks.", Timestamp: time.Now()}, 
		model.Tweet{ID: bson.NewObjectId(), Owner: "JasonHo", Message: "Hello from Jason Ho.", Timestamp: time.Now()}, 
		model.Tweet{ID: bson.NewObjectId(), Owner: userIDs[0].Hex(), Message: "Hello world!", Timestamp: time.Now()}, 
		model.Tweet{ID: bson.NewObjectId(), Owner: userIDs[0].Hex(), Message: "Harvest Moon: All you need to know It's the full moon closest to Sept equinox - coming up Oct 5.", Timestamp: time.Now()}, 
		model.Tweet{ID: bson.NewObjectId(), Owner: userIDs[0].Hex(), Message: "Hey yo!!", Timestamp: time.Now()}, 
		model.Tweet{ID: bson.NewObjectId(), Owner: userIDs[1].Hex(), Message: "Hi, I am Chih-Yin Lee. Weather sucks.", Timestamp: time.Now()}, 
		model.Tweet{ID: bson.NewObjectId(), Owner: userIDs[1].Hex(), Message: "Hello from Chih-Yin Lee.", Timestamp: time.Now()}, 
		model.Tweet{ID: bson.NewObjectId(), Owner: userIDs[2].Hex(), Message: "Hi, I am Jason He. Weather sucks.", Timestamp: time.Now()}, 
		model.Tweet{ID: bson.NewObjectId(), Owner: userIDs[2].Hex(), Message: "Hello from Jason He.", Timestamp: time.Now()}, 
		model.Tweet{ID: bson.NewObjectId(), Owner: userIDs[3].Hex(), Message: "Hi, I am Diane Lin. Weather sucks.", Timestamp: time.Now()}, 
		model.Tweet{ID: bson.NewObjectId(), Owner: userIDs[3].Hex(), Message: "Hello from Diane Lin.", Timestamp: time.Now()},
		model.Tweet{ID: bson.NewObjectId(), Owner: "JasonHo", Message: "Hi, I am Jason Ho. Weather sucks.", Timestamp: time.Now()}, 
		model.Tweet{ID: bson.NewObjectId(), Owner: "JasonHo", Message: "Hello from Jason Ho.", Timestamp: time.Now()}, 
		model.Tweet{ID: bson.NewObjectId(), Owner: userIDs[0].Hex(), Message: "Hello world!", Timestamp: time.Now()}, 
		model.Tweet{ID: bson.NewObjectId(), Owner: userIDs[0].Hex(), Message: "Harvest Moon: All you need to know It's the full moon closest to Sept equinox - coming up Oct 5.", Timestamp: time.Now()}, 
		model.Tweet{ID: bson.NewObjectId(), Owner: userIDs[0].Hex(), Message: "Hey yo!!", Timestamp: time.Now()}, 
		model.Tweet{ID: bson.NewObjectId(), Owner: userIDs[1].Hex(), Message: "Hi, I am Chih-Yin Lee. Weather sucks.", Timestamp: time.Now()}, 
		model.Tweet{ID: bson.NewObjectId(), Owner: userIDs[1].Hex(), Message: "Hello from Chih-Yin Lee.", Timestamp: time.Now()}, 
		model.Tweet{ID: bson.NewObjectId(), Owner: userIDs[2].Hex(), Message: "Hi, I am Jason He. Weather sucks.", Timestamp: time.Now()}, 
		model.Tweet{ID: bson.NewObjectId(), Owner: userIDs[2].Hex(), Message: "Hello from Jason He.", Timestamp: time.Now()}, 
		model.Tweet{ID: bson.NewObjectId(), Owner: userIDs[3].Hex(), Message: "Hi, I am Diane Lin. Weather sucks.", Timestamp: time.Now()}, 
		model.Tweet{ID: bson.NewObjectId(), Owner: userIDs[3].Hex(), Message: "Hello from Diane Lin.", Timestamp: time.Now()},
		model.Tweet{ID: bson.NewObjectId(), Owner: "JasonHo", Message: "Hi, I am Jason Ho. Weather sucks.", Timestamp: time.Now()}, 
		model.Tweet{ID: bson.NewObjectId(), Owner: "JasonHo", Message: "Hello from Jason Ho.", Timestamp: time.Now()}, 
		model.Tweet{ID: bson.NewObjectId(), Owner: userIDs[0].Hex(), Message: "Hello world!", Timestamp: time.Now()}, 
		model.Tweet{ID: bson.NewObjectId(), Owner: userIDs[0].Hex(), Message: "Harvest Moon: All you need to know It's the full moon closest to Sept equinox - coming up Oct 5.", Timestamp: time.Now()}, 
		model.Tweet{ID: bson.NewObjectId(), Owner: userIDs[0].Hex(), Message: "Hey yo!!", Timestamp: time.Now()}, 
		model.Tweet{ID: bson.NewObjectId(), Owner: userIDs[1].Hex(), Message: "Hi, I am Chih-Yin Lee. Weather sucks.", Timestamp: time.Now()}, 
		model.Tweet{ID: bson.NewObjectId(), Owner: userIDs[1].Hex(), Message: "Hello from Chih-Yin Lee.", Timestamp: time.Now()}, 
		model.Tweet{ID: bson.NewObjectId(), Owner: userIDs[2].Hex(), Message: "Hi, I am Jason He. Weather sucks.", Timestamp: time.Now()}, 
		model.Tweet{ID: bson.NewObjectId(), Owner: userIDs[2].Hex(), Message: "Hello from Jason He.", Timestamp: time.Now()}, 
		model.Tweet{ID: bson.NewObjectId(), Owner: userIDs[3].Hex(), Message: "Hi, I am Diane Lin. Weather sucks.", Timestamp: time.Now()}, 
		model.Tweet{ID: bson.NewObjectId(), Owner: userIDs[3].Hex(), Message: "Hello from Diane Lin.", Timestamp: time.Now()},
		model.Tweet{ID: bson.NewObjectId(), Owner: "JasonHo", Message: "Hi, I am Jason Ho. Weather sucks.", Timestamp: time.Now()}, 
		model.Tweet{ID: bson.NewObjectId(), Owner: "JasonHo", Message: "Hello from Jason Ho.", Timestamp: time.Now()}, 
		model.Tweet{ID: bson.NewObjectId(), Owner: userIDs[0].Hex(), Message: "Hello world!", Timestamp: time.Now()}, 
		model.Tweet{ID: bson.NewObjectId(), Owner: userIDs[0].Hex(), Message: "Harvest Moon: All you need to know It's the full moon closest to Sept equinox - coming up Oct 5.", Timestamp: time.Now()}, 
		model.Tweet{ID: bson.NewObjectId(), Owner: userIDs[0].Hex(), Message: "Hey yo!!", Timestamp: time.Now()}, 
		model.Tweet{ID: bson.NewObjectId(), Owner: userIDs[1].Hex(), Message: "Hi, I am Chih-Yin Lee. Weather sucks.", Timestamp: time.Now()}, 
		model.Tweet{ID: bson.NewObjectId(), Owner: userIDs[1].Hex(), Message: "Hello from Chih-Yin Lee.", Timestamp: time.Now()}, 
		model.Tweet{ID: bson.NewObjectId(), Owner: userIDs[2].Hex(), Message: "Hi, I am Jason He. Weather sucks.", Timestamp: time.Now()}, 
		model.Tweet{ID: bson.NewObjectId(), Owner: userIDs[2].Hex(), Message: "Hello from Jason He.", Timestamp: time.Now()}, 
		model.Tweet{ID: bson.NewObjectId(), Owner: userIDs[3].Hex(), Message: "Hi, I am Diane Lin. Weather sucks.", Timestamp: time.Now()}, 
		model.Tweet{ID: bson.NewObjectId(), Owner: userIDs[3].Hex(), Message: "Hello from Diane Lin.", Timestamp: time.Now()},
		model.Tweet{ID: bson.NewObjectId(), Owner: "JasonHo", Message: "Hi, I am Jason Ho. Weather sucks.", Timestamp: time.Now()}, 
		model.Tweet{ID: bson.NewObjectId(), Owner: "JasonHo", Message: "Hello from Jason Ho.", Timestamp: time.Now()}, 
		model.Tweet{ID: bson.NewObjectId(), Owner: userIDs[0].Hex(), Message: "Hello world!", Timestamp: time.Now()}, 
		model.Tweet{ID: bson.NewObjectId(), Owner: userIDs[0].Hex(), Message: "Harvest Moon: All you need to know It's the full moon closest to Sept equinox - coming up Oct 5.", Timestamp: time.Now()}, 
		model.Tweet{ID: bson.NewObjectId(), Owner: userIDs[0].Hex(), Message: "Hey yo!!", Timestamp: time.Now()}, 
		model.Tweet{ID: bson.NewObjectId(), Owner: userIDs[1].Hex(), Message: "Hi, I am Chih-Yin Lee. Weather sucks.", Timestamp: time.Now()}, 
		model.Tweet{ID: bson.NewObjectId(), Owner: userIDs[1].Hex(), Message: "Hello from Chih-Yin Lee.", Timestamp: time.Now()}, 
		model.Tweet{ID: bson.NewObjectId(), Owner: userIDs[2].Hex(), Message: "Hi, I am Jason He. Weather sucks.", Timestamp: time.Now()}, 
		model.Tweet{ID: bson.NewObjectId(), Owner: userIDs[2].Hex(), Message: "Hello from Jason He.", Timestamp: time.Now()}, 
		model.Tweet{ID: bson.NewObjectId(), Owner: userIDs[3].Hex(), Message: "Hi, I am Diane Lin. Weather sucks.", Timestamp: time.Now()}, 
		model.Tweet{ID: bson.NewObjectId(), Owner: userIDs[3].Hex(), Message: "Hello from Diane Lin.", Timestamp: time.Now()},
		model.Tweet{ID: bson.NewObjectId(), Owner: "JasonHo", Message: "Hi, I am Jason Ho. Weather sucks.", Timestamp: time.Now()}, 
		model.Tweet{ID: bson.NewObjectId(), Owner: "JasonHo", Message: "Hello from Jason Ho.", Timestamp: time.Now()}, 
		model.Tweet{ID: bson.NewObjectId(), Owner: userIDs[0].Hex(), Message: "Hello world!", Timestamp: time.Now()}, 
		model.Tweet{ID: bson.NewObjectId(), Owner: userIDs[0].Hex(), Message: "Harvest Moon: All you need to know It's the full moon closest to Sept equinox - coming up Oct 5.", Timestamp: time.Now()}, 
		model.Tweet{ID: bson.NewObjectId(), Owner: userIDs[0].Hex(), Message: "Hey yo!!", Timestamp: time.Now()}, 
		model.Tweet{ID: bson.NewObjectId(), Owner: userIDs[1].Hex(), Message: "Hi, I am Chih-Yin Lee. Weather sucks.", Timestamp: time.Now()}, 
		model.Tweet{ID: bson.NewObjectId(), Owner: userIDs[1].Hex(), Message: "Hello from Chih-Yin Lee.", Timestamp: time.Now()}, 
		model.Tweet{ID: bson.NewObjectId(), Owner: userIDs[2].Hex(), Message: "Hi, I am Jason He. Weather sucks.", Timestamp: time.Now()}, 
		model.Tweet{ID: bson.NewObjectId(), Owner: userIDs[2].Hex(), Message: "Hello from Jason He.", Timestamp: time.Now()}, 
		model.Tweet{ID: bson.NewObjectId(), Owner: userIDs[3].Hex(), Message: "Hi, I am Diane Lin. Weather sucks.", Timestamp: time.Now()}, 
		model.Tweet{ID: bson.NewObjectId(), Owner: userIDs[3].Hex(), Message: "Hello from Diane Lin.", Timestamp: time.Now()},
		model.Tweet{ID: bson.NewObjectId(), Owner: "JasonHo", Message: "Hi, I am Jason Ho. Weather sucks.", Timestamp: time.Now()}, 
		model.Tweet{ID: bson.NewObjectId(), Owner: "JasonHo", Message: "Hello from Jason Ho.", Timestamp: time.Now()}, 
		model.Tweet{ID: bson.NewObjectId(), Owner: userIDs[0].Hex(), Message: "Hello world!", Timestamp: time.Now()}, 
		model.Tweet{ID: bson.NewObjectId(), Owner: userIDs[0].Hex(), Message: "Harvest Moon: All you need to know It's the full moon closest to Sept equinox - coming up Oct 5.", Timestamp: time.Now()}, 
		model.Tweet{ID: bson.NewObjectId(), Owner: userIDs[0].Hex(), Message: "Hey yo!!", Timestamp: time.Now()}, 
		model.Tweet{ID: bson.NewObjectId(), Owner: userIDs[1].Hex(), Message: "Hi, I am Chih-Yin Lee. Weather sucks.", Timestamp: time.Now()}, 
		model.Tweet{ID: bson.NewObjectId(), Owner: userIDs[1].Hex(), Message: "Hello from Chih-Yin Lee.", Timestamp: time.Now()}, 
		model.Tweet{ID: bson.NewObjectId(), Owner: userIDs[2].Hex(), Message: "Hi, I am Jason He. Weather sucks.", Timestamp: time.Now()}, 
		model.Tweet{ID: bson.NewObjectId(), Owner: userIDs[2].Hex(), Message: "Hello from Jason He.", Timestamp: time.Now()}, 
		model.Tweet{ID: bson.NewObjectId(), Owner: userIDs[3].Hex(), Message: "Hi, I am Diane Lin. Weather sucks.", Timestamp: time.Now()}, 
		model.Tweet{ID: bson.NewObjectId(), Owner: userIDs[3].Hex(), Message: "Hello from Diane Lin.", Timestamp: time.Now()},
		model.Tweet{ID: bson.NewObjectId(), Owner: "JasonHo", Message: "Hi, I am Jason Ho. Weather sucks.", Timestamp: time.Now()}, 
		model.Tweet{ID: bson.NewObjectId(), Owner: "JasonHo", Message: "Hello from Jason Ho.", Timestamp: time.Now()}, 
		model.Tweet{ID: bson.NewObjectId(), Owner: userIDs[0].Hex(), Message: "Hello world!", Timestamp: time.Now()}, 
		model.Tweet{ID: bson.NewObjectId(), Owner: userIDs[0].Hex(), Message: "Harvest Moon: All you need to know It's the full moon closest to Sept equinox - coming up Oct 5.", Timestamp: time.Now()}, 
		model.Tweet{ID: bson.NewObjectId(), Owner: userIDs[0].Hex(), Message: "Hey yo!!", Timestamp: time.Now()}, 
		model.Tweet{ID: bson.NewObjectId(), Owner: userIDs[1].Hex(), Message: "Hi, I am Chih-Yin Lee. Weather sucks.", Timestamp: time.Now()}, 
		model.Tweet{ID: bson.NewObjectId(), Owner: userIDs[1].Hex(), Message: "Hello from Chih-Yin Lee.", Timestamp: time.Now()}, 
		model.Tweet{ID: bson.NewObjectId(), Owner: userIDs[2].Hex(), Message: "Hi, I am Jason He. Weather sucks.", Timestamp: time.Now()}, 
		model.Tweet{ID: bson.NewObjectId(), Owner: userIDs[2].Hex(), Message: "Hello from Jason He.", Timestamp: time.Now()}, 
		model.Tweet{ID: bson.NewObjectId(), Owner: userIDs[3].Hex(), Message: "Hi, I am Diane Lin. Weather sucks.", Timestamp: time.Now()}, 
		model.Tweet{ID: bson.NewObjectId(), Owner: userIDs[3].Hex(), Message: "Hello from Diane Lin.", Timestamp: time.Now()}}
		

	for _, t := range tweets {
		err := tweetC.Insert(t)
		if err != nil {
			panic(err)
		}
	}

	db.Close()
}

func dbFind(db *mgo.Session) {
	collect := db.DB("se_avengers").C("tweets")

	from := "JasonHo"
	var result model.User
	err := collect.Find(bson.M{"from": from}).One(&result)
	if err != nil {
		panic(err)
	}

	fmt.Println(result.ID)

	db.Close()
}

func reconstructTestDB() {
	session, err := mgo.Dial("mongodb://SEavenger:SEavenger@ds121225.mlab.com:21225/se_avengers_test")
	if err != nil {
		panic(err)
	}
	defer session.Close()

	err = session.DB("se_avengers_test").DropDatabase()
	if err != nil {
		panic(err)
	}

	userC := session.DB("se_avengers_test").C("users")

	users := []model.User {
		model.User{ID: bson.NewObjectId(), Username: "testSignup", FirstName: "test", LastName: "signup", Password: "test", Email: "testSignup@gmail.com", Followers: []string{}, Following: []string{}, 
			Bio: "testtest", Token: ""},
		model.User{ID: bson.NewObjectId(), Username: "testLogin", FirstName: "testLogin", Password: "test", Email: "testLogin@gmail.com", Followers: []string{}, Following: []string{}},
		model.User{ID: bson.NewObjectId(), Username: "testUserInfo_1", FirstName: "testUserInfo_1", Password: "test", Email: "testUserInfo_1@gmail.com", Followers: []string{}, Following: []string{"testUserInfo_2"}},
		model.User{ID: bson.NewObjectId(), Username: "testUserInfo_2", FirstName: "testUserInfo_2", Password: "test", Email: "testUserInfo_2@gmail.com", Followers: []string{"testUserInfo_1"}, Following: []string{}},
	}

	for _, u:= range users {
		err := userC.Insert(u)
		if err != nil {
			panic(err)
		}
	}
}