package handler

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"gopkg.in/mgo.v2"
	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
)

func TestFetchUserInfo (t *testing.T) {
	session, err := mgo.Dial("mongodb://SEavanger:SEavanger@ds139964.mlab.com:39964/se_avangers")
	if err != nil {
		panic(err)
	}

	// Setup
	e := echo.New()
	h := &Handler{session}

	// test cases
	requestParam := []string {
		"JasonHo", 
		"MarsLee",
		"JasonHe",  
		"DianeLin"}

	expectedJSON := []string {
		`[{"firstname":"Jason","lastname":"Ho","password":"test1","email":"hojason117@gmail.com"}]`, 
		`[{"firstname":"Chih-Yin","lastname":"Lee","password":"test2","email":"c788678867886@gmail.com"}]`, 
		`[{"firstname":"Jason","lastname":"He","password":"test3","email":"hexing_h@hotmail.com"}]`, 
		`[{"firstname":"Diane","lastname":"Lin","password":"test4","email":"diane@gmail.com"}]`,
	}

	// Run
	for i, rp := range requestParam {
		// Setup
		req := httptest.NewRequest(echo.GET, "/", nil)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		// Set the registered path for the handler.
		c.SetPath("/api/v1/UserInfolist")
		// Set path parameter names.
		c.SetParamNames("user")
		// Set path parameter values.
		c.SetParamValues(rp)

		// Assertion
		if assert.NoError(t, h.FetchUserInfo(c)) {
			assert.Equal(t, http.StatusOK, rec.Code)
			assert.Equal(t, expectedJSON[i], rec.Body.String())
		}
	}
}
