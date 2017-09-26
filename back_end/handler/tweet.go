package handler

import (
	"gopkg.in/mgo.v2"
	"fmt"
	"net/http"
	"gopkg.in/mgo.v2/bson"
	"github.com/labstack/echo"
	"model"
)

type ownTweetContainer struct {
		FirstName 	string				`json:"firstname" bson:"firstname"`
		LastName	string				`json:"lastname,omitempty"`
		Bio			string				`json:"bio,omitempty"`
		Tweets		[]tweetContainer	`json:"tweets,omitempty"`
}

type tweetContainer struct {
	Content 	string	`json:"content"`
   	Timestamp 	string	`json:"timestamp"`
}

// FetchOwnTweets : Handle requests asking for a list of tweets posted by a specific user, and respond with that list along with some user info.
func (h *Handler) FetchOwnTweets (c echo.Context) (err error) {
	//userID := userIDFromToken(c)

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
	if !bson.IsObjectIdHex(userID) {
		//return c.JSON(http.StatusBadRequest, "Malformed user ID.")
	}

	db := h.DB.Clone()
	
	// Retrieve user info from database
	user := model.User{}
	//err = db.DB("se_avengers").C("users").Find(bson.M{"_id": bson.ObjectIdHex(userID)}).One(&user)
	err = db.DB("se_avengers").C("users").Find(bson.M{"useriddev": userID}).One(&user)
	if err != nil {
		if err == mgo.ErrNotFound {
			return c.JSON(http.StatusNotFound, "User does not exist.")
		}
	}

	// Retrieve tweets from database
	tweets := []*model.Tweet{}
	//err = db.DB("se_avengers").C("tweets").Find(bson.M{"owner": bson.ObjectIdHex(userID)}).All(&tweets)
	err = db.DB("se_avengers").C("tweets").Find(bson.M{"from": userID}).All(&tweets)
	if err != nil {
		return
	}

	defer db.Close()

	res := ownTweetContainer{FirstName: user.FirstName, LastName: user.LastName, Bio: user.Bio, Tweets: []tweetContainer{}}
	for _, t := range tweets {
		time := t.Timestamp
		res.Tweets = append(res.Tweets, tweetContainer{Content: t.Message, Timestamp: fmt.Sprintf("%d-%d-%d", time.Year(), time.Month(), time.Day())})
	}
	
	return c.JSON(http.StatusOK, res)
}

// NewTweet : Add one tweet for a specific user.
func (h *Handler) NewTweet(c echo.Context) (err error) {
	userID := c.Param("user")
	if !bson.IsObjectIdHex(userID) {
		return c.JSON(http.StatusBadRequest, "Malformed user ID.")
	}

	db := h.DB.Clone()



	defer db.Close()



	//return c.JSON(http.StatusOK, )
	return c.NoContent(http.StatusNotImplemented)
}

// DeleteTweet : Delete a specific tweet.
func (h *Handler) DeleteTweet(c echo.Context) (err error) {
	tweetID := c.Param("tweet")
	if !bson.IsObjectIdHex(tweetID) {
		return c.JSON(http.StatusBadRequest, "Malformed tweet ID.")
	}

	db := h.DB.Clone()

	err = db.DB("se_avengers").C("tweets").Remove(bson.M{"_id": bson.ObjectIdHex(tweetID)})
	if err != nil {
		if err == mgo.ErrNotFound {
			return c.JSON(http.StatusNotFound, "Tweet does not exist.")
		}
	}

	defer db.Close()

	return c.NoContent(http.StatusNoContent)
}
