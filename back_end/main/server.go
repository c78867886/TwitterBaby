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
	e.Use(middleware.JWTWithConfig(middleware.JWTConfig{
		SigningKey: []byte(handler.Key),
		Skipper: func(c echo.Context) bool {
			// Skip authentication for and signup login requests
			if c.Request().Method == "OPTIONS" || c.Path() == "/api/v1/login" || c.Path() == "/api/v1/signup" {
				return true
			}
			return false
		},
	}))

	// Database connection
	session, err := mgo.Dial("mongodb://SEavenger:SEavenger@ds149324.mlab.com:49324/se_avengers")
	if err != nil {
		e.Logger.Fatal(err)
	}
	defer session.Close()

	// Initialize handler
	h := &handler.Handler{DB: session}

	// CORS config
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"http://localhost:4200"},
		AllowMethods: []string{echo.GET, echo.PUT, echo.POST, echo.DELETE},
		AllowCredentials: true, 
		AllowHeaders: []string{echo.HeaderAuthorization, echo.HeaderContentType}, 
	}))

	// Routes
	e.POST("/api/v1/signup", h.Signup)
	e.POST("/api/v1/login", h.Login)
	e.POST("/api/v1/follow/:id", h.Follow)
	e.GET("/api/v1/tweetlist/:username", h.FetchTweets)
	e.GET("/api/v1/userInfo/:id", h.FetchUserInfo)
	e.GET("/api/v1/userInfo", h.FetchUserInfo)
	e.POST("/api/v1/newTweet/:id", h.NewTweet)
	e.DELETE("/api/v1/deleteTweet/:tweet", h.DeleteTweet)
	e.GET("/api/v1/tweettimeline/:username", h.FetchTweetTimeLine)
	e.POST("/api/v1/updateUserInfo", h.UpdateUserInfo)
	e.GET("/api/v1/showFollower/:id", h.ShowFollower)
	e.GET("/api/v1/showFollowing/:id", h.ShowFollowing)

	// Set server address
	srvAddr := "localhost:1323"
	//srvAddr := "127.0.0.1:1323"
	//srvAddr := "192.168.1.2:1323"

	// Initiate parallel server control
	go serverControl(e, session, srvAddr)

	// Start server
	e.HideBanner = true
	e.Logger.Fatal(e.Start(srvAddr))

	// c.Path() == "/" || c.Path() == "/index.html" || c.Path() == "/favicon.ico" || c.Path() == "/inline.bundle.js" || c.Path() == "/inline.bundle.js.map" 
	// || c.Path() == "/main.bundle.js.map" || c.Path() == "/polyfills.bundle.js" || c.Path() == "/polyfills.bundle.js.map" || c.Path() == "/styles.bundle.js" 
	// || c.Path() == "/styles.bundle.js.map" || c.Path() == "/vendor.bundle.js" || c.Path() == "/vendor.bundle.js.map" || c.Path() == "/main.bundle.js"

	//e.Use(middleware.StaticWithConfig(middleware.StaticConfig{Root: "../../bin/dist", Browse: true}))
	//e.Static("/", "../../bin/dist/assets")
	//e.File("/index.html", "../../bin/dist/index.html")
	//e.File("/main.bundle.js", "../../bin/dist/assets/main.bundle.js")
}

func serverControl(e *echo.Echo, session *mgo.Session, srvAddr string) {
	var op string
	
	for {
		fmt.Printf("Listening on %s\n", srvAddr)
		fmt.Print("Option('h' for help): ")
		fmt.Scanln(&op)
		if op == "q" {
			fmt.Println("Shutting down server.")
			shutdownServer(e)
			break
		} else if op == "h" {
			fmt.Println("'h' for help")
			fmt.Println("'q' to shutdown server")
			fmt.Println("'d' to drop database")
			fmt.Println("'i' to reconstruct database to default (w/ some initial collections)")
		} else if op == "d" {
			dbDrop(session.Clone())
			fmt.Println("Dropped database.")
		} else if op == "i" {
			dbReinsert(session.Clone())
			fmt.Println("Database reconstructed.")
		}
	}
}

func shutdownServer(e *echo.Echo) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := e.Shutdown(ctx); err != nil {
		e.Logger.Fatal(err)
	}
}