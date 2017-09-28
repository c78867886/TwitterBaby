package handler

import (
	"fmt"
	"gopkg.in/mgo.v2"
	"time"
	"net/http"
	"gopkg.in/mgo.v2/bson"
	"github.com/labstack/echo"
	"model"
	//"github.com/dgrijalva/jwt-go"
	"strconv"
	"math"
	"math/rand"
	//"fmt"
)

// FetchTweets : Handle requests asking for a list of tweets posted by a specific user.
//				 URL: "/api/v1/tweetlist/:user"
//				 Method: GET
//				 Return 200 OK on success.
//				 Return 404 Not Found if the user is not in the database.
func (h *Handler) FetchTweets (c echo.Context) (err error) {
	userID := c.Param("user")
	page, err := strconv.Atoi(c.QueryParam("page"))
	perpage, err := strconv.Atoi(c.QueryParam("perpage"))
	
	db := h.DB.Clone()
	defer db.Close()

	// Retrieve tweets from database
	tweets := []model.Tweet{}
	err = db.DB("se_avangers").C("tweets").Find(bson.M{"owner": userID}).Sort("timestamp").All(&tweets)

	if err != nil {
		if err == mgo.ErrNotFound {
			return &echo.HTTPError{Code: http.StatusNotFound, Message: "Can not find and Tweet from this user."}
		}
		return
	}

	totalTweets := len(tweets)
	fmt.Println(totalTweets)
	totalPage := int(math.Ceil(float64(totalTweets)/float64(perpage)))

	var tweetList [] model.Tweet
	if page == totalPage{
		tweetList = tweets[perpage*(page-1):]
	}else{
		tweetList = tweets[perpage*(page-1):perpage*page]
	}

	var container struct {
		Page	int	`json:"page"`
		TotalPage	int	`json:"totalPage"`
		TotalTweets	int	`json:"totalTweets"`
		TweetList []model.Tweet `json:"tweetList"`
	}
	container.Page = page
	container.TotalPage = totalPage
	container.TotalTweets = totalTweets
	container.TweetList = tweetList
	
	return c.JSON(http.StatusOK, container)
	//return c.JSON(http.StatusOK, &tweets)
	
}

// NewTweet : Add one tweet for a specific user.
//			  URL: "/api/v1/newTweet"
//			  Method: POST
//			  Return 201 Created on success, along with the tweet data.
//			  Return 404 Not Found if the user is not in the database.
//			  Return 400 Bad Request if the content of the tweet is empty.
func (h *Handler) NewTweet(c echo.Context) (err error) {
	//userID := userIDFromToken(c)
	userID := c.Param("user")

	db := h.DB.Clone()
	defer db.Close()

	// Retrieve user info from database
	/*
	user := model.User{}
	err = db.DB("se_avengers").C("users").FindId(bson.ObjectIdHex(userID)).One(&user)
	if err != nil {
		if err == mgo.ErrNotFound {
			return &echo.HTTPError{Code: http.StatusNotFound, Message: "User does not exist."}
		}
		return
	}*/

	var tweet *model.Tweet
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	for i := 0; i < 124; i++{
		tweet := &model.Tweet{ID: bson.NewObjectId(), Owner: userID, Timestamp: time.Now(), Message: "Rock the world lol! #"+strconv.Itoa(r1.Intn(1000))}
		if err = c.Bind(tweet); err != nil {
			return
		}
		
		// Validation
		if tweet.Message == "" {
			return &echo.HTTPError{Code: http.StatusBadRequest, Message: "Message cannot be empty."}
		}
	
		// Save tweet in database
		err = db.DB("se_avangers").C("tweets").Insert(tweet)
		if err != nil {
			return
		}
	}

	return c.JSON(http.StatusCreated, tweet)
}

// DeleteTweet : Delete a specific tweet.
//				 URL: "/api/v1/deleteTweet/:tweet"
//				 Method: DELETE
//				 Return 204 No Content on success.
// 				 Return 400 Bad Request if tweet ID is malformed.
//				 Return 404 Not Found if the tweet is not in the database.
func (h *Handler) DeleteTweet(c echo.Context) (err error) {
	tweetID := c.Param("tweet")

	if !bson.IsObjectIdHex(tweetID) {
		return &echo.HTTPError{Code: http.StatusBadRequest, Message: "Malformed tweet ID."}
	}

	db := h.DB.Clone()
	defer db.Close()

	err = db.DB("se_avengers").C("tweets").RemoveId(bson.ObjectIdHex(tweetID))
	if err != nil {
		if err == mgo.ErrNotFound {
			return &echo.HTTPError{Code: http.StatusNotFound, Message: "Tweet does not exist."}
		}
		return
	}

	return c.NoContent(http.StatusNoContent)
}

