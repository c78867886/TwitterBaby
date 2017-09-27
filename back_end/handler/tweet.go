package handler

import (
	"gopkg.in/mgo.v2"
	"time"
	"net/http"
	"gopkg.in/mgo.v2/bson"
	"github.com/labstack/echo"
	"model"
)

// FetchOwnTweets : Handle requests asking for a list of tweets posted by a specific user, and respond with that list along with some user info.
//					Return 200 OK on success.
// 					Return 400 Bad Request if user ID is malformed.
//					Return 404 Not Found if the user is not in the database.
//					Request format: GET "api/v1/tweetlist/:user"
func (h *Handler) FetchTweets (c echo.Context) (err error) {
	/*page, _ := strconv.Atoi(c.QueryParam("page"))
	limit, _ := strconv.Atoi(c.QueryParam("limit"))

	// Defaults
	if page == 0 {
		page = 1
	}
	if limit == 0 {
		limit = 100
	}*/

	userID := c.Param("user")

	db := h.DB.Clone()
	defer db.Close()
	
	// Retrieve user info from database
	user := model.User{}
	err = db.DB("se_avengers").C("users").FindId(bson.ObjectIdHex(userID)).One(&user)
	if err != nil {
		if err == mgo.ErrNotFound {
			return &echo.HTTPError{Code: http.StatusNotFound, Message: "User does not exist."}
		}
		return
	}

	// Retrieve tweets from database
	tweets := []model.Tweet{}
	err = db.DB("se_avengers").C("tweets").Find(bson.M{"owner": userID}).All(&tweets)
	if err != nil {
		return
	}
	
	return c.JSON(http.StatusOK, &tweets)
}

// NewTweet : Add one tweet for a specific user.
//			  Return 201 Created on success
func (h *Handler) NewTweet(c echo.Context) (err error) {
	userID := userIDFromToken(c)

	db := h.DB.Clone()
	defer db.Close()

	// Retrieve user info from database
	user := model.User{}
	err = db.DB("se_avengers").C("users").FindId(bson.ObjectIdHex(userID)).One(&user)
	if err != nil {
		if err == mgo.ErrNotFound {
			return &echo.HTTPError{Code: http.StatusNotFound, Message: "User does not exist."}
		}
		return
	}

	tweet := &model.Tweet{ID: bson.NewObjectId(), Owner: userID, Timestamp: time.Now()}
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
//				 Return 204 No Content on success.
// 				 Return 400 Bad Request if tweet ID is malformed.
//				 Return 404 Not Found if the tweet is not in the database.
//				 Request format: DELETE "api/v1/deleteTweet/:tweet"
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
