package handler

import (
	"gopkg.in/mgo.v2"
	"time"
	"net/http"
	"gopkg.in/mgo.v2/bson"
	"github.com/labstack/echo"
	"model"
	"github.com/dgrijalva/jwt-go"
	"math"
	"strconv"
)

// FetchTweets : Handle requests asking for a list of tweets posted by a specific user.
//				 URL: "/api/v1/tweetlist/:user"
//				 Method: GET
//				 Return 200 OK on success.
//				 Return 404 Not Found if the user is not in the database.
func (h *Handler) FetchTweets (c echo.Context) (err error) {
	username := userNameFromToken(c)
	username = c.Param("username")
	/*
	t := new(model.Tweet)
	if err = c.Bind(t); err != nil {
		return
	}
	*/
	page, err := strconv.Atoi(c.QueryParam("page"))
	perpage, err := strconv.Atoi(c.QueryParam("perpage"))

	db := h.DB.Clone()
	defer db.Close()

	// Retrieve user info from database by username
	user := model.User{}
	err = db.DB("se_avengers").C("users").Find(bson.M{"username": username}).One(&user)
	if err != nil {
		if err == mgo.ErrNotFound {
			return &echo.HTTPError{Code: http.StatusNotFound, Message: "User "+username+" does not exist."}
		}
		return
	}
	id := user.ID.Hex()

	// Retrieve tweets from database
	tweets := []model.Tweet{}
	err = db.DB("se_avengers").C("tweets").Find(bson.M{"owner": id}).Sort("-timestamp").All(&tweets)

	if err != nil {
		if err == mgo.ErrNotFound {
			return &echo.HTTPError{Code: http.StatusNotFound, Message: "Can not find any Tweet from this user."}
		}
		return 
	}

	totalTweets := len(tweets)
	totalPage := int(math.Ceil(float64(totalTweets)/float64(perpage)))

	var tweetList [] model.Tweet
	if page == totalPage{
		tweetList = tweets[perpage*(page-1):]
	}else{
		tweetList = tweets[perpage*(page-1):perpage*page]
	}

	// Change id to username
	for i := range tweetList {
		tweetList[i].Owner = user.Username
	}

	var container struct {
		Page	string	`json:"page"`
		TotalPage	string	`json:"totalpage"`
		TotalTweets	string	`json:"totaltweets"`
		TweetList []model.Tweet `json:"tweetlist"`
	}
	container.Page = strconv.Itoa(page)
	container.TotalPage = strconv.Itoa(totalPage)
	container.TotalTweets = strconv.Itoa(totalTweets)
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
	userID := userIDFromToken(c)

	db := h.DB.Clone()
	defer db.Close()

	tweet := &model.Tweet{ID: bson.NewObjectId(), Owner: string(userID), Timestamp: time.Now()}
	if err = c.Bind(tweet); err != nil {
		return
	}
	
	
	// Validation
	if tweet.Message == "" {
		return &echo.HTTPError{Code: http.StatusBadRequest, Message: "Message cannot be empty."}
	}
	
	// Save tweet in database
	err = db.DB("se_avengers").C("tweets").Insert(tweet)
	if err != nil {
		return
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

func (h *Handler) FetchTweetTimeLine (c echo.Context) (err error) {
	username := userNameFromToken(c)
	username = c.Param("username")
	
	page, err := strconv.Atoi(c.QueryParam("page"))
	perpage, err := strconv.Atoi(c.QueryParam("perpage"))

	db := h.DB.Clone()
	defer db.Close()

	// Retrieve user info from database by username
	user := model.User{}
	err = db.DB("se_avengers").C("users").Find(bson.M{"username": username}).One(&user)
	if err != nil {
		if err == mgo.ErrNotFound {
			return &echo.HTTPError{Code: http.StatusNotFound, Message: "User "+username+" does not exist."}
		}
		return
	}
	// Get the user list of tweet timeline.
	following := user.Following
	timeLineUserList := append(following, user.ID.Hex())

	// Map between ID and username
	mapIDandUsername := make(map[string]string)
	for i := range timeLineUserList{
		tempUser := model.User{}
		err = db.DB("se_avengers").C("users").FindId(bson.ObjectIdHex(timeLineUserList[i])).One(&tempUser)
		mapIDandUsername[tempUser.ID.Hex()] = tempUser.Username
	}

	// Retrieve tweets from database
	tweets := []model.Tweet{}
	err = db.DB("se_avengers").C("tweets").Find(bson.M{"owner": bson.M{"$in": timeLineUserList}}).Sort("-timestamp").All(&tweets)
	if err != nil {
		if err == mgo.ErrNotFound {
			return &echo.HTTPError{Code: http.StatusNotFound, Message: "Can not find any Tweet from this user."}
		}
		return 
	}

	totalTweets := len(tweets)
	totalPage := int(math.Ceil(float64(totalTweets)/float64(perpage)))

	var tweetList [] model.Tweet
	if page == totalPage{
		tweetList = tweets[perpage*(page-1):]
	}else{
		tweetList = tweets[perpage*(page-1):perpage*page]
	}

	// Change id to username
	for i := range tweetList {
		tweetList[i].Owner = mapIDandUsername[tweetList[i].Owner]
	}

	var container struct {
		Page	string	`json:"page"`
		TotalPage	string	`json:"totalpage"`
		TotalTweets	string	`json:"totaltweets"`
		TweetList []model.Tweet `json:"tweetlist"`
	}
	container.Page = strconv.Itoa(page)
	container.TotalPage = strconv.Itoa(totalPage)
	container.TotalTweets = strconv.Itoa(totalTweets)
	container.TweetList = tweetList
	

	return c.JSON(http.StatusOK, container)
	//return c.JSON(http.StatusOK, &tweets)
}

func userNameFromToken(c echo.Context) string {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)

	return claims["username"].(string)
}
