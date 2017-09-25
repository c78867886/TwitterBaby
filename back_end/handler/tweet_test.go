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
		`[{"content":"Hi, I am Jason Ho.","timestamp":"2017-9-25"},{"content":"Hello from Jason Ho.","timestamp":"2017-9-25"},{"content":"Hello world!","timestamp":"2017-9-25"}]`, 
		`[{"content":"Hi, I am Jason He.","timestamp":"2017-9-25"},{"content":"Hello from Jason He.","timestamp":"2017-9-25"}]`, 
		`[{"content":"Hi, I am Chih-Yin Lee.","timestamp":"2017-9-25"},{"content":"Hello from Chih-Yin Lee.","timestamp":"2017-9-25"}]`, 
		`[{"content":"Hi, I am Diane Lin.","timestamp":"2017-9-25"},{"content":"Hello from Diane Lin.","timestamp":"2017-9-25"}]`}

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
