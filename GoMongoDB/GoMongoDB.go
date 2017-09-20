package main

import (
	"fmt"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"time"
)
//"mongodb://SEavanger:SEavanger@ds139964.mlab.com:39964/se_avangers"

type Person struct {
	ID        bson.ObjectId `bson:"_id,omitempty"`
	Firstname      string
	Lastname      string
	Email     string
	Psw	string
	Age int
	Timestamp time.Time
}

func main() {
	session, err := mgo.Dial("mongodb://SEavanger:SEavanger@ds139964.mlab.com:39964/se_avangers")
	if err != nil {
		panic(err)
	}

	defer session.Close()

	session.SetMode(mgo.Monotonic, true)

	// Drop Database
	/*
	err = session.DB("se_avangers").DropDatabase()
	if err != nil {
		panic(err)
	}
	*/

	// Collection People
	c := session.DB("se_avangers").C("user")

	// Index
	index := mgo.Index{
		Key:        []string{"id"},
		Unique:     true,
		DropDups:   true,
		Background: true,
		Sparse:     true,
	}

	err = c.EnsureIndex(index)
	if err != nil {
		panic(err)
	}

	// Insert Datas
	err = c.Insert(
		&Person{Firstname: "Ale", Lastname: "Wong", Email: "ccjdis@gmail.com", Psw:"2234", Age:23, Timestamp: time.Now()},
		&Person{Firstname: "Boyyu", Lastname: "Wong", Email: "73737827@gmail.com", Psw:"343232", Age:53, Timestamp: time.Now()})

	if err != nil {
		panic(err)
	}

	// Query All
	var results []Person
	err = c.Find(bson.M{"firstname": "Ale"}).Sort("-timestamp").All(&results)

	if err != nil {
		panic(err)
	}
	fmt.Println("Results All: ", results)

	// Update
	/*
	colQuerier := bson.M{"name": "Ale"}
	change := bson.M{"$set": bson.M{"phone": "+86 99 8888 7777", "timestamp": time.Now()}}
	err = c.Update(colQuerier, change)
	if err != nil {
		panic(err)
	}
	*/

}