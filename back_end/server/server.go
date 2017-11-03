package server

import (
	"handler"
	"notification"
	"fmt"
	"time"
	"context"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/labstack/gommon/log"
)

// NewServer : Instantiate a server
func NewServer(h *handler.Handler, nh *notification.Handler) (e *echo.Echo) {
	e = echo.New()
	e.HideBanner = true
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
	
	// CORS config
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"http://localhost:4200"},
		AllowMethods: []string{echo.GET, echo.PUT, echo.POST, echo.DELETE},
		AllowCredentials: true, 
		AllowHeaders: []string{echo.HeaderAuthorization, echo.HeaderContentType}, 
	}))

	// Routes
	e.GET("/api/v1/ws", nh.GetConnection)
	e.POST("/api/v1/signup", h.Signup)
	e.POST("/api/v1/login", h.Login)
	e.POST("/api/v1/follow/:username", h.Follow)
	e.POST("/api/v1/unfollow/:username", h.Unfollow)
	e.GET("/api/v1/userInfo/:username", h.FetchUserInfo)
	e.POST("/api/v1/updateUserInfo", h.UpdateUserInfo)
	e.POST("/api/v1/updateProfilePic", h.UpdateProfilePicture)
	e.GET("/api/v1/showFollower/:username", h.ShowFollower)
	e.GET("/api/v1/showFollowing/:username", h.ShowFollowing)
	e.GET("/api/v1/tweetlist/:username", h.FetchTweets)
	e.POST("/api/v1/newTweet/:id", h.NewTweet)
	e.DELETE("/api/v1/deleteTweet/:tweet", h.DeleteTweet)
	e.GET("/api/v1/tweettimeline/:username", h.FetchTweetTimeLine)
	
	// c.Path() == "/" || c.Path() == "/index.html" || c.Path() == "/favicon.ico" || c.Path() == "/inline.bundle.js" || c.Path() == "/inline.bundle.js.map" 
	// || c.Path() == "/main.bundle.js.map" || c.Path() == "/polyfills.bundle.js" || c.Path() == "/polyfills.bundle.js.map" || c.Path() == "/styles.bundle.js" 
	// || c.Path() == "/styles.bundle.js.map" || c.Path() == "/vendor.bundle.js" || c.Path() == "/vendor.bundle.js.map" || c.Path() == "/main.bundle.js"

	//e.Use(middleware.StaticWithConfig(middleware.StaticConfig{Root: "../../bin/dist", Browse: true}))
	//e.Static("/", "../../bin/dist/assets")
	//e.File("/index.html", "../../bin/dist/index.html")
	//e.File("/main.bundle.js", "../../bin/dist/assets/main.bundle.js")

	return e;
}

// TerminalControl : Thread for terminal control
func TerminalControl(e *echo.Echo, h *handler.Handler, srvAddr string) {
	var op string
	
	for {
		fmt.Println("Listening on " + srvAddr)
		fmt.Print("Option('h' for help): ")
		fmt.Scanln(&op)
		if op == "q" {
			fmt.Println("Shutting down server.")
			ShutdownServer(e, h)
			break
		} else if op == "h" {
			fmt.Println("'h' for help")
			fmt.Println("'q' to shutdown server")
			fmt.Println("'d' to drop database")
			fmt.Println("'i' to reconstruct database to default (w/ some initial collections)")
			fmt.Println("'r' to reconstruct testing database")
		} else if op == "d" {
			dbDrop(h.DB.Clone())
			fmt.Println("Dropped database.")
		} else if op == "i" {
			dbReinsert(h.DB.Clone())
			fmt.Println("Database reconstructed.")
		} else if op == "r" {
			reconstructTestDB()
			fmt.Println("Testing database reconstructed.")
		}
	}
}

// ShutdownServer : Shutdown the server
func ShutdownServer(e *echo.Echo, h *handler.Handler) {
	h.DB.Close()

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := e.Shutdown(ctx); err != nil {
		e.Logger.Fatal(err)
	}
}