package main

import (
	"fmt"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"time"
)

type Person struct {
	ID        bson.ObjectId `bson:"_id,omitempty"`
	firstname      string
	lastname      string
	email     string
	psw	string
	age int
	timestamp time.Time

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

	// Collection user
	c := session.DB("se_avangers").C("user")

	// Index
	index := mgo.Index{
		Key:        []string{"name", "phone"},
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
		&Person{firstname: "Ale", lastname: "Wong", email: "ccjdis@gmail.com", psw:"2234", age:23, timestamp: time.Now()},
		&Person{firstname: "Boyyu", lastname: "Wong", email: "73737827@gmail.com", psw:"343232", age:53, timestamp: time.Now()})

	if err != nil {
		panic(err)
	}

	// Query
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
