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
	h := &Handler{session}

	// test empty request parameter
	//c.SetParamValues("")

	// test cases
	requestParam := []string {
		"JasonHo", 
		"JasonHe", 
		"MarsLee", 
		"DianeLin"}

	expectedJSON := []string {
		`{"firstname":"Jason","lastname":"Ho","bio":"Hi everyone, this is Jason Ho.","tweets":[{"content":"Hi, I am Jason Ho. Weather sucks.","timestamp":"2017-9-25"},{"content":"Hello from Jason Ho.","timestamp":"2017-9-25"},{"content":"Hello world!","timestamp":"2017-9-25"}]}`, 
		`{"firstname":"Jason","lastname":"He","bio":"Hi everyone, this is Jason He.","tweets":[{"content":"Hi, I am Jason He. Weather sucks.","timestamp":"2017-9-25"},{"content":"Hello from Jason He.","timestamp":"2017-9-25"}]}`, 
		`{"firstname":"Chih-Yin","lastname":"Lee","bio":"Hi everyone, this is Mars Lee.","tweets":[{"content":"Hi, I am Chih-Yin Lee. Weather sucks.","timestamp":"2017-9-25"},{"content":"Hello from Chih-Yin Lee.","timestamp":"2017-9-25"}]}`, 
		`{"firstname":"Diane","lastname":"Lin","bio":"Hi everyone, this is Diane Lin.","tweets":[{"content":"Hi, I am Diane Lin. Weather sucks.","timestamp":"2017-9-25"},{"content":"Hello from Diane Lin.","timestamp":"2017-9-25"}]}`}

	// Run
	for i, rp := range requestParam {
		// Setup
		req := httptest.NewRequest(echo.GET, "/", nil)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetPath("api/v1/tweetlist")
		c.SetParamNames("user")
		c.SetParamValues(rp)

		// Assertion
		if assert.NoError(t, h.FetchOwnTweets(c)) {
			assert.Equal(t, http.StatusOK, rec.Code)
			assert.Equal(t, expectedJSON[i], rec.Body.String())
		}
	}
}
