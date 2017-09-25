package handler

import (
	"fmt"
	"net/http"
	"gopkg.in/mgo.v2/bson"
	"github.com/labstack/echo"
	"model"
)

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

	userID := c.Param("user")

	// Retrieve tweets from database
	tweets := []*model.Tweet{}
	db := h.DB.Clone()

	err = db.DB("se_avangers").C("tweets").Find(bson.M{"from": userID}).All(&tweets)
	if err != nil {
		return
	}
	defer db.Close()
	
	res := []tweetContainer{}
	for _, t := range tweets {
		time := t.Timestamp
		res = append(res, tweetContainer{Content: t.Message, Timestamp: fmt.Sprintf("%d-%d-%d", time.Year(), time.Month(), time.Day())})
	}
	
	return c.JSON(http.StatusOK, res)
}