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

	fmt.Println("Receive request.")

	userID := c.Param("user")

	// Retrieve tweets from database
	tweets := []*model.Tweet{}
	db := h.DB.Clone()

	err = db.DB("se_avangers").C("tweets").Find(bson.M{"from": userID}).All(&tweets)
	if err != nil {
		return
	}
	defer db.Close()

	type container struct {
		content 	string
   		timestamp 	string
	} 
	//temp := []container {container{content: tweets[0].Message, timestamp: tweets[0].Timestamp}}
	//temp := []container {container{content: "123", timestamp: tweets[0].Timestamp}}
	//var temp = []*container{}
	//temp = append(temp, &container{content: "123", timestamp: tweets[0].Timestamp})
	//temp[0].content = tweets[0].Message
	//temp[0].timestamp = tweets[0].Timestamp
	//fmt.Println(temp[0])

	//return c.JSON(http.StatusOK, temp)
	
	return c.JSON(http.StatusOK, tweets)
	//return c.JSON(http.StatusOK, tweets[0])
}