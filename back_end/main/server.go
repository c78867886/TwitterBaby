package main

import (
	"server"
	"handler"
	"notification"
)

func main() {
	// Instantiate server
	dbURL := "mongodb://SEavenger:SEavenger@ds149324.mlab.com:49324/se_avengers"
	srvAddr := "localhost:1323"
	h := handler.NewHandler(dbURL)
	nh := notification.NewHandler()
	e := server.NewServer(h, nh)

	// Initiate parallel server control
	go server.TerminalControl(e, h, srvAddr)

	// Start server
	e.Logger.Fatal(e.Start(srvAddr))
}