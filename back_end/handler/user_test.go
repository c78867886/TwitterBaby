package handler

import (
	"model"
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

type testCase struct {
	positive	bool
	input		input
	expected	expected
}

type input struct {
	user	string
	target	string
}

type expected struct {
	code 		int
	message 	string
}

func TestSignup(t *testing.T) {
	// Setup
	e := echo.New()
	h := NewHandler("mongodb://SEavenger:SEavenger@ds121225.mlab.com:21225/se_avengers_test")

	// Test cases
	testCases := []testCase {
		// test success
		testCase{
			true,
			input{"", `{"username":"testSignup","firstname":"test","lastname":"signup","password":"test","email":"testSignup@gmail.com","bio":"testtest"}`},
			expected{http.StatusCreated, `{"username":"testSignup","firstname":"test","lastname":"signup","email":"testSignup@gmail.com","bio":"testtest"}`},
		},
		testCase{
			true,
			input{"", `{"username":"testSignup","firstname":"test","lastname":"signup","password":"test","email":"testSignup@gmail.com"}`,},
			expected{http.StatusCreated, `{"username":"testSignup","firstname":"test","lastname":"signup","email":"testSignup@gmail.com"}`},
		},
		testCase{
			true,
			input{"", `{"username":"testSignup","firstname":"test","password":"test","email":"testSignup@gmail.com","bio":"testtest"}`},
			expected{http.StatusCreated, `{"username":"testSignup","firstname":"test","email":"testSignup@gmail.com","bio":"testtest"}`},
		},
		// test empty
		testCase{
			false,
			input{"", `{"firstname":"test","lastname":"signup","password":"test","email":"testSignup@gmail.com"}`},
			expected{http.StatusBadRequest, "Invalid username, firstname, email or password."},
		},
		testCase{
			false,
			input{"", `{"username":"testSignup","lastname":"signup","password":"test","email":"testSignup@gmail.com"}`,},
			expected{http.StatusBadRequest, "Invalid username, firstname, email or password."},
		},
		testCase{
			false,
			input{"", `{"username":"testSignup","firstname":"test","lastname":"signup","password":"test"}`},
			expected{http.StatusBadRequest, "Invalid username, firstname, email or password."},
		},
		testCase{
			false,
			input{"", `{"username":"testSignup","firstname":"test","lastname":"signup","email":"testSignup@gmail.com","bio":"testtest"}`},
			expected{http.StatusBadRequest, "Invalid username, firstname, email or password."},
		},
		// test duplicate
		testCase{
			false,
			input{"", `{"username":"testSignup_dup","firstname":"test","lastname":"signup_dup","password":"test","email":"testSignup_dup@gmail.com","bio":"testtest"}`},
			expected{http.StatusBadRequest, "Username or email already used."},
		},
	}

	// Run
	for _, tc := range testCases {
		// Delete DB entry
		if tc.positive {
			err := h.DB.DB(h.DBName).C("users").Remove(bson.M{"username": "testSignup"})
			if err != nil {
				if err != mgo.ErrNotFound {
					t.Fatal(err)
				}
			}
		}

		// Setup
		req := httptest.NewRequest(echo.POST, "/", strings.NewReader(tc.input.target))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		// Set the registered path for the handler.
		c.SetPath("/api/v1/signup")

		// Assertion
		if tc.positive {
			if assert.NoError(t, h.Signup(c), tc.input.target) {
				assert.Equal(t, tc.expected.code, rec.Code, tc.input.target)
				assert.Equal(t, tc.expected.message, rec.Body.String(), tc.input.target)
			}
		} else {
			if err := h.Signup(c); assert.Error(t, err, tc.input.target) {
				assert.Equal(t, tc.expected.code, err.(*echo.HTTPError).Code, tc.input.target)
				assert.Equal(t, tc.expected.message, err.(*echo.HTTPError).Message, tc.input.target)
			}
		}
	}
	h.DB.Close()
}

func TestLogin(t *testing.T) {
	e := echo.New()
	h := NewHandler("mongodb://SEavenger:SEavenger@ds121225.mlab.com:21225/se_avengers_test")

	testCases := []testCase {
		// test success
		testCase{
			true,
			input{"", `{"email":"testLogin@gmail.com","password":"test"}`},
			expected{http.StatusOK, `{"username":"testLogin","firstname":"testLogin","email":"testLogin@gmail.com"`},
		},
		// tes fail
		testCase{
			false,
			input{"", `{"email":"testLogin@gmail.com","password":""}`},
			expected{http.StatusUnauthorized, "Invalid email or password."},
		},
		testCase{
			false,
			input{"", `{"email":"","password":"test"}`},
			expected{http.StatusUnauthorized, "Invalid email or password."},
		},
		testCase{
			false,
			input{"", `{"email":"testLogin","password":"test"}`},
			expected{http.StatusUnauthorized, "Invalid email or password."},
		},
	}

	for _, tc := range testCases {
		req := httptest.NewRequest(echo.POST, "/", strings.NewReader(tc.input.target))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetPath("/api/v1/login")

		if tc.positive {
			if assert.NoError(t, h.Login(c), tc.input.target) {
				assert.Equal(t, tc.expected.code, rec.Code, tc.input.target)
				assert.Equal(t, tc.expected.message, strings.Split(rec.Body.String(), `,"token`)[0], tc.input.target)
			}
		} else {
			if err := h.Login(c); assert.Error(t, err, tc.input.target) {
				assert.Equal(t, tc.expected.code, err.(*echo.HTTPError).Code, tc.input.target)
				assert.Equal(t, tc.expected.message, err.(*echo.HTTPError).Message, tc.input.target)
			}
		}
	}
	h.DB.Close()
}

func TestFetchUserInfo(t *testing.T) {
	e := echo.New()
	h := NewHandler("mongodb://SEavenger:SEavenger@ds121225.mlab.com:21225/se_avengers_test")

	testCases := []testCase {
		// test success
		testCase{
			true,
			input{"testUserInfo_1", "testUserInfo_2"},
			expected{http.StatusOK, `{"userinfo":{"username":"testUserInfo_2","firstname":"testUserInfo_2","email":"testUserInfo_2@gmail.com"},"followercount":1,"followingcount":0,"followed":true}`},
		},
		testCase{
			true,
			input{"testUserInfo_2", "testUserInfo_1"},
			expected{http.StatusOK, `{"userinfo":{"username":"testUserInfo_1","firstname":"testUserInfo_1","email":"testUserInfo_1@gmail.com"},"followercount":0,"followingcount":1,"followed":false}`},
		},
		// tes fail
		testCase{
			false,
			input{"testUserInfo_1", "userNotExist"},
			expected{http.StatusNotFound, "User does not exist."},
		},
	}

	for _, tc := range testCases {
		req := httptest.NewRequest(echo.GET, "/", nil)
		req.Header.Set(echo.HeaderAuthorization, "Bearer " + getToken(tc.input.user))
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetPath("/api/v1/userInfo")
		c.SetParamNames("username")
		c.SetParamValues(tc.input.target)
		processJWTToken(c)

		if tc.positive {
			if assert.NoError(t, h.FetchUserInfo(c), tc.input.user + " fetching " + tc.input.target) {
				assert.Equal(t, tc.expected.code, rec.Code, tc.input.user + " fetching " + tc.input.target)
				assert.Equal(t, tc.expected.message, rec.Body.String(), tc.input.user + " fetching " + tc.input.target)
			}
		} else {
			if err := h.FetchUserInfo(c); assert.Error(t, err, tc.input.user + " fetching " + tc.input.target) {
				assert.Equal(t, tc.expected.code, err.(*echo.HTTPError).Code, tc.input.user + " fetching " + tc.input.target)
				assert.Equal(t, tc.expected.message, err.(*echo.HTTPError).Message, tc.input.user + " fetching " + tc.input.target)
			}
		}
	}
	h.DB.Close()
}

func TestFollow(t *testing.T) {
	e := echo.New()
	h := NewHandler("mongodb://SEavenger:SEavenger@ds121225.mlab.com:21225/se_avengers_test")

	testCases := []testCase {
		// test success
		testCase{
			true,
			input{"testFollow", "testFollow_2"},
			expected{http.StatusOK, `["testFollow_1","testFollow_2"]`},
		},
		// test fail
		testCase{
			false,
			input{"testFollow", "testFollow_not_exist"},
			expected{http.StatusNotFound, "User does not exist."},
		},
	}

	for _, tc := range testCases {
		req := httptest.NewRequest(echo.POST, "/", nil)
		req.Header.Set(echo.HeaderAuthorization, "Bearer " + getToken(tc.input.user))
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetPath("/api/v1/follow")
		c.SetParamNames("username")
		c.SetParamValues(tc.input.target)
		processJWTToken(c)

		if tc.positive {
			if assert.NoError(t, h.Follow(c), tc.input.user + " follows " + tc.input.target) {
				assert.Equal(t, tc.expected.code, rec.Code, tc.input.user + " follows " + tc.input.target)
				assert.Equal(t, tc.expected.message, rec.Body.String(), tc.input.user + " follows " + tc.input.target)

				user := model.User{}
				h.DB.DB(h.DBName).C("users").Find(bson.M{"username": tc.input.target}).One(&user)
				assert.Equal(t, []string{tc.input.user}, user.Followers, tc.input.user + " follows " + tc.input.target)
			}
		} else {
			if err := h.Follow(c); assert.Error(t, err, tc.input.user + " follows " + tc.input.target) {
				assert.Equal(t, tc.expected.code, err.(*echo.HTTPError).Code, tc.input.user + " follows " + tc.input.target)
				assert.Equal(t, tc.expected.message, err.(*echo.HTTPError).Message, tc.input.user + " follows " + tc.input.target)
			}
		}
	}
	h.DB.Close()
}

func TestShowFollower(t *testing.T) {
	e := echo.New()
	h := NewHandler("mongodb://SEavenger:SEavenger@ds121225.mlab.com:21225/se_avengers_test")

	testCases := []testCase {
		// test success
		testCase{
			true,
			input{"", "testShowFollower_1"},
			expected{http.StatusOK, "[]"},
		},
		testCase{
			true,
			input{"", "testShowFollower_2"},
			expected{http.StatusOK, `[{"username":"testShowFollower_1","firstname":"testShowFollower_1"},{"username":"testShowFollower_3","firstname":"testShowFollower_3"}]`},
		},
		testCase{
			true,
			input{"", "testShowFollower_3"},
			expected{http.StatusOK, `[{"username":"testShowFollower_2","firstname":"test","lastname":"ShowFollower_2","bio":"testtest"}]`},
		},
		// test fail
		testCase{
			false,
			input{"", "testShowFollower_not_exist"},
			expected{http.StatusNotFound, "User does not exist."},
		},
	}

	for _, tc := range testCases {
		req := httptest.NewRequest(echo.GET, "/", nil)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetPath("/api/v1/showFollower")
		c.SetParamNames("username")
		c.SetParamValues(tc.input.target)

		if tc.positive {
			if assert.NoError(t, h.ShowFollower(c), tc.input.target) {
				assert.Equal(t, tc.expected.code, rec.Code, tc.input.target)
				assert.Equal(t, tc.expected.message, rec.Body.String(), tc.input.target)
			}
		} else {
			if err := h.ShowFollower(c); assert.Error(t, err, tc.input.target) {
				assert.Equal(t, tc.expected.code, err.(*echo.HTTPError).Code, tc.input.target)
				assert.Equal(t, tc.expected.message, err.(*echo.HTTPError).Message, tc.input.target)
			}
		}
	}
	h.DB.Close()
}

func TestShowFollowing(t *testing.T) {
	e := echo.New()
	h := NewHandler("mongodb://SEavenger:SEavenger@ds121225.mlab.com:21225/se_avengers_test")

	testCases := []testCase {
		// test success
		testCase{
			true,
			input{"", "testShowFollowing_1"},
			expected{http.StatusOK, "[]"},
		},
		testCase{
			true,
			input{"", "testShowFollowing_2"},
			expected{http.StatusOK, `[{"username":"testShowFollowing_1","firstname":"testShowFollowing_1"},{"username":"testShowFollowing_3","firstname":"testShowFollowing_3"}]`},
		},
		testCase{
			true,
			input{"", "testShowFollowing_3"},
			expected{http.StatusOK, `[{"username":"testShowFollowing_2","firstname":"test","lastname":"ShowFollowing_2","bio":"testtest"}]`},
		},
		// test fail
		testCase{
			false,
			input{"", "testShowFollowing_not_exist"},
			expected{http.StatusNotFound, "User does not exist."},
		},
	}

	for _, tc := range testCases {
		req := httptest.NewRequest(echo.GET, "/", nil)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetPath("/api/v1/showFollowing")
		c.SetParamNames("username")
		c.SetParamValues(tc.input.target)

		if tc.positive {
			if assert.NoError(t, h.ShowFollowing(c), tc.input.target) {
				assert.Equal(t, tc.expected.code, rec.Code, tc.input.target)
				assert.Equal(t, tc.expected.message, rec.Body.String(), tc.input.target)
			}
		} else {
			if err := h.ShowFollowing(c); assert.Error(t, err, tc.input.target) {
				assert.Equal(t, tc.expected.code, err.(*echo.HTTPError).Code, tc.input.target)
				assert.Equal(t, tc.expected.message, err.(*echo.HTTPError).Message, tc.input.target)
			}
		}
	}
	h.DB.Close()
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