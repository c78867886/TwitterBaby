package handler

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"gopkg.in/mgo.v2"
	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
)

func TestFetchOwnTweets (t *testing.T) {
	session, err := mgo.Dial("mongodb://SEavanger:SEavanger@ds139964.mlab.com:39964/se_avangers")
	if err != nil {
		panic(err)
	}

	// Setup
	e := echo.New()
	req := httptest.NewRequest(echo.GET, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/api/v1/tweetlist")
	c.SetParamNames("user")
	c.SetParamValues("JasonHo")
	h := &Handler{session}

	// Assertions
	var tweetJSON = `[{"id":"59c8402ca54d7515eefc62b8","from":"JasonHo","message":"Hi, I am Jason Ho.","timestamp":"2017.1.1"}]`
	if assert.NoError(t, h.FetchOwnTweets(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, tweetJSON, rec.Body.String())
	}
}

/*func TestFetchOwnTweets_1 (t *testing.T) {
	// Setup
	e := echo.New()
	session, err := mgo.Dial("mongodb://SEavanger:SEavanger@ds139964.mlab.com:39964/se_avangers")
	if err != nil {
		panic(err)
	}
	h := &Handler{DB: session}
	e.GET("/api/v1/tweetlist", h.FetchOwnTweets)
	
	req := httptest.NewRequest(echo.GET, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/api/v1/tweetlist")
	c.SetParamNames("user")
	c.SetParamValues("JasonHo")

	// Assertions
	var tweetJSON = `[{"id":"59c82b80a54d750fd33d5527","from":"JasonHo","message":"Hi, I am Jason Ho."}]`
	if assert.NoError(t, h.FetchOwnTweets(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, tweetJSON, rec.Body.String())
	}
}*/