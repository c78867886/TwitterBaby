package main

import (
	"fmt"
	"gopkg.in/mgo.v2"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/labstack/gommon/log"
	"context"
	"time"
	"handler"
)

func main() {
	e := echo.New()
	e.Logger.SetLevel(log.ERROR)
	e.Use(middleware.Logger())
	/*e.Use(middleware.JWTWithConfig(middleware.JWTConfig{
		SigningKey: []byte(handler.Key),
		Skipper: func(c echo.Context) bool {
			// Skip authentication for and signup login requests
			if c.Path() == "/login" || c.Path() == "/signup" {
				return true
			}
			return false
		},
	}))*/

	DBDrop()
	DBInsert()

	// Database connection
	session, err := mgo.Dial("mongodb://SEavanger:SEavanger@ds139964.mlab.com:39964/se_avangers")
	if err != nil {
		e.Logger.Fatal(err)
	}
	defer session.Close()

	// Initialize handler
	h := &handler.Handler{DB: session}

	// Routes
	e.GET("/tweets", h.FetchOwnTweets)

	// initiate parallel server control
	go serverControl(e, session)

	// Start server
	e.HideBanner = true
	e.Logger.Fatal(e.Start(":1323"))
}

func serverControl(e *echo.Echo, session *mgo.Session) {
	var op string
	
	for {
		fmt.Print("Option: ")
		fmt.Scanln(&op)
		if op == "q" {
			fmt.Println("Shutting down server.")
			shutdownServer(e, session)
			break
		}
	}
}

func shutdownServer(e *echo.Echo, session *mgo.Session) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := e.Shutdown(ctx); err != nil {
		e.Logger.Fatal(err)
	}

	session.Close()
}