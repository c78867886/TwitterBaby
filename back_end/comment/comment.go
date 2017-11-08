package comment

import (
	"strings"
	"time"
	"net/http"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"github.com/labstack/echo"
	"model"
	"github.com/dgrijalva/jwt-go"
	//"math"
	"strconv"
)

type (
	// Handler : Handler for comment related requests.
	Handler struct {
		db 				*mgo.Session
		dbName 			string
		key				string
	}
)

// NewHandler : Create a Handler instance.
func NewHandler(dbURL string, key string) (h *Handler) {
	// Database connection
	session, err := mgo.Dial(dbURL)
	if err != nil {
		session.Close()
		panic(err)
	}

	return &Handler{session, strings.Split(dbURL, "/")[3], key}
}

// NewComment : Add one comment for a specific tweet.
//			  URL: "/api/v1/newcomment/:tweet"
//			  Method: POST
//			  Return 200 Created on success, along with the comment data.
//			  Return 404 Not Found if the tweetID is not in the database.
//			  Return 400 Bad Request if the content of the comment is empty.
func (h *Handler) NewComment(c echo.Context) (err error) {
	userName := userNameFromToken(c)
	tweetID := c.Param("tweet")

	db := h.db.Clone()
	defer db.Close()

	// Check tweetID is valid
	if bson.IsObjectIdHex(tweetID){
		validTweet := new(model.Tweet)
		err = db.DB(h.dbName).C("tweets").FindId(bson.ObjectIdHex(tweetID)).One(&validTweet)
		if err != nil {
			if err == mgo.ErrNotFound {
				return &echo.HTTPError{Code: http.StatusNotFound, Message: "Invalid tweet ID"}
			}
			return &echo.HTTPError{Code: http.StatusNotFound, Message: "Invalid tweet ID"}
		}
	}else{
		return &echo.HTTPError{Code: http.StatusNotFound, Message: "Invalid tweet ID"}
	}

	comment := &model.Comment{ID: bson.NewObjectId(), FromTweetID: string(tweetID), FromUsername: string(userName), Timestamp: time.Now()}
	if err = c.Bind(comment); err != nil {
		return
	}
	
	// Validation
	if comment.Message == "" {
		return &echo.HTTPError{Code: http.StatusBadRequest, Message: "Message cannot be empty."}
	}
	
	// Save comment in database
	err = db.DB(h.dbName).C("comments").Insert(comment)
	if err != nil {
		return
	}

	// update comment number of tweet
	tweet := model.Tweet{}
	db.DB(h.dbName).C("tweets").FindId(bson.ObjectIdHex(tweetID)).One(&tweet)
	
	tweetid := tweet.ID.Hex()
	err = db.DB(h.dbName).C("tweets").Update(bson.M{"_id": bson.ObjectIdHex(tweetid)}, bson.M{"$set": bson.M{"numcomment": tweet.Numcomment+1}})


	comment.ID = ""
	var container struct {
		FromTweetID	string	`json:"fromtweetid"`
		FromUsername	string	`json:"fromusername"`
		Message	string	`json:"message"`
	}
	container.FromTweetID = comment.FromTweetID
	container.FromUsername = comment.FromUsername
	container.Message = comment.Message

	//h.NotifHandler.Manager.Operator <- model.Notification{Timestamp: time.Now(), Detail: model.NewTweetNotif{Publisher: userName}}

	return c.JSON(http.StatusOK, container)
}

// FetchComment : Handle requests asking for a list of comment depend on a specific tweet.
//				 URL: "/api/v1/fetchcomment/:tweet"
//				 Method: GET
//				 Return 200 OK on success.
//				 Return 404 Not Found if the tweetID is not in the database.
func (h *Handler) FetchComment (c echo.Context) (err error) {
	//username := userNameFromToken(c)
	tweetID := c.Param("tweet")
	
	db := h.db.Clone()
	defer db.Close()

	// Check tweetID is valid
	if bson.IsObjectIdHex(tweetID){
		validTweet := new(model.Tweet)
		err = db.DB(h.dbName).C("tweets").FindId(bson.ObjectIdHex(tweetID)).One(&validTweet)
		if err != nil {
			if err == mgo.ErrNotFound {
				return &echo.HTTPError{Code: http.StatusNotFound, Message: "Invalid tweet ID"}
			}
			return &echo.HTTPError{Code: http.StatusNotFound, Message: "Invalid tweet ID"}
		}
	}else{
		return &echo.HTTPError{Code: http.StatusNotFound, Message: "Invalid tweet ID"}
	}

	// Retrieve comments from database
	comments := []model.Comment{}
	err = db.DB(h.dbName).C(model.CommentCollection).Find(bson.M{"fromtweetid": tweetID}).Sort("-timestamp").All(&comments)

	if err != nil {
		return
	}

	var container struct {
		TotalComment string `json:"totalcomment"`
		CommentList []model.Comment `json:"commentlist"`
	}
	container.CommentList = comments
	container.TotalComment = strconv.Itoa(len(comments))
	
	return c.JSON(http.StatusOK, container)
}

func userNameFromToken(c echo.Context) string {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)

	return claims["username"].(string)
}

// Shutdown : Gracefully shutdown comment handler.
func (h *Handler) Shutdown() {
	h.db.Close()
}
