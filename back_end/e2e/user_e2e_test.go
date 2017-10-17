package e2e

import (
	"server"
	"handler"
	"testing"
	"strings"
	"net/http"
	"net/http/httptest"
	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

func Test_Signup_Login(t *testing.T) {
	dbURL := "mongodb://SEavenger:SEavenger@ds121225.mlab.com:21225/se_avengers_test"
	h := handler.NewHandler(dbURL)
	e := server.NewServer(h)

	type response struct {
		code 		int
		message 	string
	}

	testSuccessSignup := `{"username":"testSignupLogin","firstname":"test","lastname":"SignupLogin","password":"test","email":"testSignupLogin@gmail.com","bio":"testtest"}`

	testSuccessLogin := `{"email":"testSignupLogin@gmail.com","password":"test"}`

	testSuccessExpected := response{code: http.StatusOK, message: `{"username":"testSignupLogin","firstname":"test","lastname":"SignupLogin","email":"testSignupLogin@gmail.com","bio":"testtest"`}

	err := h.DB.DB(h.DBName).C("users").Remove(bson.M{"username": "testSignup"})
	if err != nil {
		if err != mgo.ErrNotFound {
			t.Fatal(err)
		}
	}

	req := httptest.NewRequest(echo.POST, "/api/v1/signup", strings.NewReader(testSuccessSignup))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()

	e.ServeHTTP(rec, req)
	
	req = httptest.NewRequest(echo.POST, "/api/v1/login", strings.NewReader(testSuccessLogin))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec = httptest.NewRecorder()

	e.ServeHTTP(rec, req)

	assert.Equal(t, testSuccessExpected.code, rec.Code, testSuccessSignup)
	assert.Equal(t, testSuccessExpected.message, strings.Split(rec.Body.String(), `,"token`)[0], testSuccessSignup)

	h.DB.Close()
}