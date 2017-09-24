package handler

import (
	"fmt"
	"net/http"
	//"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"github.com/labstack/echo"
	"model"
)

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

	userID := c.Get("user")
	fmt.Println(userID)
	// Retrieve tweets from database
	tweets := []*model.Tweet{}
	db := h.DB.Clone()

	err = db.DB("se_avangers").C("tweets").Find(bson.M{"from": userID}).All(&tweets)
	if err != nil {
		return
	}
	defer db.Close()
	fmt.Println("success")
	fmt.Println(tweets[0].Message)
	return c.JSON(http.StatusOK, tweets)
}