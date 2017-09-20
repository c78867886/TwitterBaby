package main

import (
	"time"

	"gopkg.in/mgo.v2"
)

func Insert() {
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
		&Person{Firstname: "Ale", Lastname: "Wong", Email: "ccjdis@gmail.com", Psw: "2234", Age: 23, Timestamp: time.Now()},
		&Person{Firstname: "Boyyu", Lastname: "Wong", Email: "73737827@gmail.com", Psw: "343232", Age: 53, Timestamp: time.Now()})

	if err != nil {
		panic(err)
	}
}
