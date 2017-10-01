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
	session, err := mgo.Dial("mongodb://SEavenger:SEavenger@ds149324.mlab.com:49324/se_avengers")
	if err != nil {
		panic(err)
	}

	// Setup
	e := echo.New()
	h := &Handler{session}

	// test empty request parameter
	//c.SetParamValues("")

	// test cases
	requestParam := []string {
		"MarsLee",}
		//"JasonHe", 
		//"MarsLee", 
		//"DianeLin", 
		//"TomRiddle"}

	expectedJSON := []string {
		`[{\"id\":\"59ca5bd5a54d750d432fa294\",\"owner\":\"\",\"from\":\"JasonHe\",\"message\":\"Hi, I am Jason He. Weather sucks.\",\"timestamp\":\"0001-01-01T00:00:00Z\"},{\"id\":\"59ca5bd5a54d750d432fa295\",\"owner\":\"\",\"from\":\"JasonHe\",\"message\":\"Hello from Jason He.\",\"timestamp\":\"0001-01-01T00:00:00Z\"}]`, }
		//`{"firstname":"Jason","lastname":"He","bio":"Hi everyone, this is Jason He.","tweets":[{"content":"Hi, I am Jason He. Weather sucks.","timestamp":"2017-9-26"},{"content":"Hello from Jason He.","timestamp":"2017-9-26"}]}`, 
		//`{"firstname":"Chih-Yin","lastname":"Lee","bio":"Hi everyone, this is Mars Lee.","tweets":[{"content":"Hi, I am Chih-Yin Lee. Weather sucks.","timestamp":"2017-9-26"},{"content":"Hello from Chih-Yin Lee.","timestamp":"2017-9-26"}]}`, 
		//`{"firstname":"Diane","lastname":"Lin","bio":"Hi everyone, this is Diane Lin.","tweets":[{"content":"Hi, I am Diane Lin. Weather sucks.","timestamp":"2017-9-26"},{"content":"Hello from Diane Lin.","timestamp":"2017-9-26"}]}`, 
		//`{"firstname":"Tom","lastname":"Riddle","bio":"Hi everyone, this is Lord Voldemort."}`}

	// Run
	for i, rp := range requestParam {
		// Setup
		req := httptest.NewRequest(echo.GET, "/?page=12&perpage=10", nil)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetPath("api/v1/tweetlist")
		c.SetParamNames("user")
		c.SetParamValues(rp)
		

		// Assertion
		if assert.NoError(t, h.FetchTweets(c)) {
			assert.Equal(t, http.StatusOK, rec.Code)
			assert.Equal(t, expectedJSON[i], rec.Body.String())
		}
	}
}

func TestNewTweet (t *testing.T) {
	session, err := mgo.Dial("mongodb://SEavenger:SEavenger@ds149324.mlab.com:49324/se_avengers")
	if err != nil {
		panic(err)
	}

	// Setup
	e := echo.New()
	h := &Handler{session}

	// test empty request parameter
	//c.SetParamValues("")

	// test cases
	requestParam := []string {
		"MarsLee",
	}

	expectedJSON := []string {
		"JasonHe",
	}
	
	// Run
	for i, rp := range requestParam {
		// Setup
		req := httptest.NewRequest(echo.GET, "/", nil)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetPath("api/v1/newTweet")
		c.SetParamNames("user")
		c.SetParamValues(rp)

		// Assertion
		if assert.NoError(t, h.NewTweet(c)) {
			assert.Equal(t, http.StatusOK, rec.Code)
			assert.Equal(t, expectedJSON[i], rec.Body.String())
		}
	}
}

func TestDeleteTweet (t *testing.T) {
	session, err := mgo.Dial("mongodb://SEavenger:SEavenger@ds149324.mlab.com:49324/se_avengers")
	if err != nil {
		panic(err)
	}

	// Setup
	e := echo.New()
	h := &Handler{session}

	// test empty request parameter
	//c.SetParamValues("")

	// test cases
	requestParam := []string {}

	// Run
	for _, rp := range requestParam {
		// Setup
		req := httptest.NewRequest(echo.GET, "/", nil)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetPath("api/v1/deleteTweet")
		c.SetParamNames("tweet")
		c.SetParamValues(rp)

		// Assertion
		if assert.NoError(t, h.FetchTweets(c)) {
			assert.Equal(t, http.StatusNoContent, rec.Code)
		}
	}
}
