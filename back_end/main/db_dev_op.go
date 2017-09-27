package main

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
		model.User{FirstName: "Jason", LastName: "Ho", Password: "test1", Email: "hojason117@gmail.com", Followers: []string{}, Following: []string{}, 
			Bio: "Hi everyone, this is Jason Ho.", Token: ""},
		model.User{FirstName: "Chih-Yin", LastName: "Lee", Password: "test2", Email: "c788678867886@gmail.com", Followers: []string{}, Following: []string{}, 
			Bio: "Hi everyone, this is Mars Lee.", Token: ""},
		model.User{FirstName: "Jason", LastName: "He", Password: "test3", Email: "hexing_h@hotmail.com", Followers: []string{}, Following: []string{}, 
			Bio: "Hi everyone, this is Jason He.", Token: ""},
		model.User{FirstName: "Diane", LastName: "Lin", Password: "test4", Email: "diane@gmail.com", Followers: []string{}, Following: []string{}, 
			Bio: "Hi everyone, this is Diane Lin.", Token: ""}, 
		model.User{FirstName: "Tom", LastName: "Riddle", Password: "test5", Email: "triddle@gmail.com", Followers: []string{}, Following: []string{}, 
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
		model.Tweet{ID: bson.NewObjectId(), Owner: userIDs[0].Hex(), Message: "Hi, I am Jason Ho. Weather sucks.", Timestamp: time.Now()}, 
		model.Tweet{ID: bson.NewObjectId(), Owner: userIDs[0].Hex(), Message: "Hello from Jason Ho.", Timestamp: time.Now()}, 
		model.Tweet{ID: bson.NewObjectId(), Owner: userIDs[0].Hex(), Message: "Hello world!", Timestamp: time.Now()}, 
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
