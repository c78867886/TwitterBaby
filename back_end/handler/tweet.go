package handler

import (
	"fmt"
	"net/http"
	"gopkg.in/mgo.v2/bson"
	"github.com/labstack/echo"
	"model"
)

type owntweetContainer struct {
		FirstName 	string				`json:"firstname" bson:"firstname"`
		LastName	string				`json:"lastname,omitempty"`
		Bio			string				`json:"bio,omitempty"`
		Tweets		[]tweetContainer	`json:"tweets,omitempty"`
}

type tweetContainer struct {
	Content 	string	`json:"content"`
   	Timestamp 	string	`json:"timestamp"`
}


/*func (h *Handler) CreateTweet(c echo.Context) (err error) {
	
}*/

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

	//userID := c.Param("user")
	userID := c.QueryParam("user")

	db := h.DB.Clone()
	
	// Retrieve user info from database
	user := model.User{}

	err = db.DB("se_avangers").C("users").Find(bson.M{"useriddev": userID}).One(&user)
	if err != nil {
		return
	}

	// Retrieve tweets from database
	tweets := []*model.Tweet{}

	err = db.DB("se_avangers").C("tweets").Find(bson.M{"from": userID}).All(&tweets)
	if err != nil {
		return
	}

	defer db.Close()

	res := owntweetContainer{FirstName: user.FirstName, LastName: user.LastName, Bio: user.Bio, Tweets: []tweetContainer{}}
	for _, t := range tweets {
		time := t.Timestamp
		res.Tweets = append(res.Tweets, tweetContainer{Content: t.Message, Timestamp: fmt.Sprintf("%d-%d-%d", time.Year(), time.Month(), time.Day())})
	}
	
	return c.JSON(http.StatusOK, res)
}