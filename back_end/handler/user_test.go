package handler

import (
	"fmt"
	"strings"
	"time"
	"net/http"
	"net/http/httptest"
	"testing"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"github.com/labstack/echo"
	"github.com/dgrijalva/jwt-go"
	"github.com/stretchr/testify/assert"
)

func TestSignup(t *testing.T) {
	// Setup
	e := echo.New()
	h := NewHandler("mongodb://SEavenger:SEavenger@ds121225.mlab.com:21225/se_avengers_test")

	// Test cases
	type response struct {
		code 		int
		message 	string
	}

	testSuccess := []string {
		`{"username":"testSignup","firstname":"test","lastname":"signup","password":"test","email":"testSignup@gmail.com","bio":"testtest"}`,
		`{"username":"testSignup","firstname":"test","lastname":"signup","password":"test","email":"testSignup@gmail.com"}`,
		`{"username":"testSignup","firstname":"test","password":"test","email":"testSignup@gmail.com","bio":"testtest"}`,
	}

	testSuccessExpected := []response {
		response{code: http.StatusCreated, message: `{"username":"testSignup","firstname":"test","lastname":"signup","email":"testSignup@gmail.com","bio":"testtest"}`},
		response{code: http.StatusCreated, message: `{"username":"testSignup","firstname":"test","lastname":"signup","email":"testSignup@gmail.com"}`},
		response{code: http.StatusCreated, message: `{"username":"testSignup","firstname":"test","email":"testSignup@gmail.com","bio":"testtest"}`},
	}

	// Test successful
	for i, tc := range testSuccess {
		// Delete DB entry
		err := h.DB.DB(h.DBName).C("users").Remove(bson.M{"username": "testSignup"})
		if err != nil {
			if err != mgo.ErrNotFound {
				t.Fatal(err)
			}
		}

		// Setup
		req := httptest.NewRequest(echo.POST, "/", strings.NewReader(tc))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		// Set the registered path for the handler.
		c.SetPath("/api/v1/signup")

		// Assertion
		if assert.NoError(t, h.Signup(c), tc) {
			assert.Equal(t, testSuccessExpected[i].code, rec.Code, tc)
			assert.Equal(t, testSuccessExpected[i].message, rec.Body.String(), tc)
		}
	}

	testEmpty := []string {
		`{"firstname":"test","lastname":"signup","password":"test","email":"testSignup@gmail.com"}`,
		`{"username":"testSignup","lastname":"signup","password":"test","email":"testSignup@gmail.com"}`,
		`{"username":"testSignup","firstname":"test","lastname":"signup","password":"test"}`,
		`{"username":"testSignup","firstname":"test","lastname":"signup","email":"testSignup@gmail.com","bio":"testtest"}`,
	}

	// Test empty fail
	for _, tc := range testEmpty {
		// Setup
		req := httptest.NewRequest(echo.POST, "/", strings.NewReader(tc))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		// Set the registered path for the handler.
		c.SetPath("/api/v1/signup")

		// Assertion
		assert.Error(t, h.Signup(c), tc)
	}

	testDuplicate := `{"username":"testSignup","firstname":"test","lastname":"signup","password":"test","email":"testSignup@gmail.com","bio":"testtest"}`

	// Test duplicate fail
	// Setup
	req := httptest.NewRequest(echo.POST, "/", strings.NewReader(testDuplicate))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	// Set the registered path for the handler.
	c.SetPath("/api/v1/signup")

	// Assertion
	assert.Error(t, h.Signup(c), testDuplicate)
}

func TestLogin(t *testing.T) {
	e := echo.New()
	h := NewHandler("mongodb://SEavenger:SEavenger@ds121225.mlab.com:21225/se_avengers_test")

	type response struct {
		code 		int
		message 	string
	}

	testSuccess := []string {
		`{"email":"testLogin@gmail.com","password":"test"}`,
	}

	testSuccessExpected := []response {
		response{code: http.StatusOK, message: `{"username":"testLogin","firstname":"testLogin","email":"testLogin@gmail.com"`},
	}

	for i, tc := range testSuccess {
		req := httptest.NewRequest(echo.POST, "/", strings.NewReader(tc))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetPath("/api/v1/login")

		if assert.NoError(t, h.Login(c)) {
			assert.Equal(t, testSuccessExpected[i].code, rec.Code, tc)
			assert.Equal(t, testSuccessExpected[i].message, strings.Split(rec.Body.String(), `,"token`)[0], tc)
		}
	}

	testFail := []string {
		`{"email":"testLogin@gmail.com","password":""}`,
		`{"email":"","password":"test"}`,
		`{"email":"testLogin","password":"test"}`,
	}

	for _, tc := range testFail {
		req := httptest.NewRequest(echo.POST, "/", strings.NewReader(tc))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetPath("/api/v1/login")

		assert.Error(t, h.Login(c), tc)
	}
}

func TestFetchUserInfo(t *testing.T) {
	e := echo.New()
	h := NewHandler("mongodb://SEavenger:SEavenger@ds121225.mlab.com:21225/se_avengers_test")
	
	type response struct {
		code 		int
		message 	string
	}

	type request struct {
		user	string
		target	string
	}

	testParamSuccess := []request {
		request {"testUserInfo_1", "testUserInfo_2"},
		request {"testUserInfo_2", "testUserInfo_1"},
	}

	testSuccessExpected := []response {
		response{code: http.StatusOK, message: `{"userinfo":{"username":"testUserInfo_2","firstname":"testUserInfo_2","email":"testUserInfo_2@gmail.com"},"followercount":1,"followingcount":0,"followed":true}`},
		response{code: http.StatusOK, message: `{"userinfo":{"username":"testUserInfo_1","firstname":"testUserInfo_1","email":"testUserInfo_1@gmail.com"},"followercount":0,"followingcount":1,"followed":false}`},
	}
	
	for i, tc := range testParamSuccess {
		req := httptest.NewRequest(echo.GET, "/", nil)
		req.Header.Set(echo.HeaderAuthorization, "Bearer " + getToken(tc.user))
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetPath("/api/v1/UserInfo")
		c.SetParamNames("username")
		c.SetParamValues(tc.target)
		processJWTToken(c)

		if assert.NoError(t, h.FetchUserInfo(c)) {
			assert.Equal(t, testSuccessExpected[i].code, rec.Code)
			assert.Equal(t, testSuccessExpected[i].message, rec.Body.String())
		}
	}
}

func TestFollow(t *testing.T) {

}

func TestShowFollower(t *testing.T) {

}

func TestShowFollowing(t *testing.T) {

}

func TestUpdateUserInfo(t *testing.T) {

}





func getToken(username string) string {
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["username"] = username
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()

	Token, err := token.SignedString([]byte(Key))
	if err != nil {
		panic(err)
	}

	return Token
}

func processJWTToken(c echo.Context) {
	type JWTConfig struct {
		SigningKey 		interface{}
		SigningMethod	string
		ContextKey		string
		TokenLookup		string
		AuthScheme 		string
		Claims 			jwt.Claims
		keyFunc 		jwt.Keyfunc
	}
	var DefaultJWTConfig = JWTConfig {
		SigningKey: 	[]byte(Key),
		SigningMethod: 	"HS256",
		ContextKey:    	"user",
		TokenLookup:   	"header:" + echo.HeaderAuthorization,
		AuthScheme:    	"Bearer",
		Claims:       	jwt.MapClaims{},
	}
	DefaultJWTConfig.keyFunc = func(t *jwt.Token) (interface{}, error) {
		if t.Method.Alg() != DefaultJWTConfig.SigningMethod {
			return nil, fmt.Errorf("Unexpected jwt signing method=%v", t.Header["alg"])
		}
		return DefaultJWTConfig.SigningKey, nil
	}

	parts := strings.Split(DefaultJWTConfig.TokenLookup, ":")
	extractor := jwtFromHeader(parts[1], DefaultJWTConfig.AuthScheme)

	auth, err := extractor(c)
	token := new(jwt.Token)
	
	if _, ok := DefaultJWTConfig.Claims.(jwt.MapClaims); ok {
		token, err = jwt.Parse(auth, DefaultJWTConfig.keyFunc)
	}

	if err == nil && token.Valid {
		c.Set(DefaultJWTConfig.ContextKey, token)
	}
}

var ErrJWTMissing = echo.NewHTTPError(http.StatusBadRequest, "Missing or malformed jwt")
type jwtExtractor func(echo.Context) (string, error)

func jwtFromHeader(header string, authScheme string) jwtExtractor {
	return func(c echo.Context) (string, error) {
		auth := c.Request().Header.Get(header)
		l := len(authScheme)
		if len(auth) > l+1 && auth[:l] == authScheme {
			return auth[l+1:], nil
		}
		return "", ErrJWTMissing
	}
}