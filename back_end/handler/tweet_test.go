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
	c.SetPath("/tweets/:user")
	c.SetParamNames("user")
	c.SetParamValues("JasonHo")
	h := &Handler{session}

	// Assertions
	var tweetJSON = `{"id":"","from":"JasonHo","message":"Hi, I am Jason Ho."}`
	
	if assert.NoError(t, h.FetchOwnTweets(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, tweetJSON, rec.Body.String())
	}
}