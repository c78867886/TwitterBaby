package comment

import (
	"model"
	//"notification"
	"fmt"
	"strings"
	"time"
	"net/http"
	"net/http/httptest"
	"testing"
	//"gopkg.in/mgo.v2"
	//"gopkg.in/mgo.v2/bson"
	"github.com/labstack/echo"
	"github.com/dgrijalva/jwt-go"
	"github.com/stretchr/testify/assert"
	"encoding/json"
)

type (
	restHTTPTestCase struct {
		positive	bool
		isinvalidtweetid	bool
		input		restHTTPInput
		expected	restHTTPExpected
	}

	restHTTPInput struct {
		user	string
		target	string
	}

	restHTTPExpected struct {
		code 		int
		message 	string
	}
)

const key  = "UYrtPaa0Pky7QZyVrkIfnouatG7LjTKystf0FGdOuDiXCZyCSuVz8YdK7OBeSrC"

func TestNewComment (t *testing.T) {
	// Setup
	e := echo.New()
	//nh := notification.NewHandler("mongodb://SEavenger:SEavenger@ds121225.mlab.com:21225/se_avengers_test")
	h := NewHandler("mongodb://SEavenger:SEavenger@ds121225.mlab.com:21225/se_avengers_test", key) //, nh.Manager.Operator)

	/*
	const illegalSize int = 10485760+1
	oversizedImage := [illegalSize]byte{}
	for i := 0; i < illegalSize; i++ {
		oversizedImage[i] = 'A'
	}
	*/

	// Get the tweetID for testing
	tempTweet := []model.Tweet{}
	h.db.DB(h.dbName).C(model.TweetCollection).Find(nil).Sort("timestamp").All(&tempTweet)
	fromTweetIDForTesting := tempTweet[1].ID.Hex()

	// Test cases
	testCases := []restHTTPTestCase {
		// test success
		restHTTPTestCase{
			true,
			false,
			restHTTPInput{"TestNewCommentSuccess", `{"message":"testMessage"}`},
			restHTTPExpected{http.StatusOK, `{"fromtweetid":"`+fromTweetIDForTesting+`","fromusername":"TestNewCommentSuccess","message":"testMessage"}`},
		},
		
		// test empty
		restHTTPTestCase{
			false,
			false,
			restHTTPInput{"TestNewCommentEmpty", `{"message":""}`},
			restHTTPExpected{http.StatusBadRequest, "Message cannot be empty."},
		},
		// test illegal tweetId
		restHTTPTestCase{
			false,
			true,
			restHTTPInput{"TestNewCommentInvalidTweetID", `{"message":"testMessagetestMessagetestMessagetestMessagetestMessagetestMessagetestMessagetestMessagetestMessagetestMessagetestMessagetestMessagetestMessagetestMessagetestMessagetestMessagetestMessagetestMessagetestMessagetestMessagetestMessagetestMessagetestMessagetestMessagetestMessagetestMessagetestMessagetestMessagetestMessagetestMessagetestMessagetestMessagetestMessagetestMessagetestMessagetestMessagetestMessagetestMessagetestMessagetestMessagetestMessagetestMessagetestMessagetestMessagetestMessagetestMessagetestMessagetestMessagetestMessagetestMessagetestMessagetestMessagetestMessagetestMessagetestMessagetestMessagetestMessagetestMessagetestMessagetestMessagetestMessagetestMessagetestMessagetestMessagetestMessagetestMessagetestMessagetestMessagetestMessagetestMessagetestMessagetestMessagetestMessagetestMessagetestMessagetestMessagetestMessagetestMessagetestMessagetestMessage"}`},
			restHTTPExpected{http.StatusNotFound, "Invalid tweet ID"},
		},

		// test illegal picture
		/*
		restHTTPTestCase{
			false,
			restHTTPInput{"testSignup", `{"message":"testMessage", "picture":"` + string(oversizedImage[:]) + `"}`},
			restHTTPExpected{http.StatusBadRequest, "Image must be smaller than 10 MB."},
		},
		*/
	}
	
	// Run
	for _, tc := range testCases {
		// Setup
		req := httptest.NewRequest(echo.POST, "/", strings.NewReader(tc.input.target))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		req.Header.Set(echo.HeaderAuthorization, "Bearer " + getToken(tc.input.user))
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetPath("/api/v1/newcomment")
		c.SetParamNames("tweet")
		if tc.isinvalidtweetid{
			c.SetParamValues("5a007fb6311bc348635124")
		}else{
			c.SetParamValues(fromTweetIDForTesting)
		}
		
		processJWTToken(c)

		// Assertion
		if tc.positive {
			if assert.NoError(t, h.NewComment(c), tc.input.target) {
				assert.Equal(t, tc.expected.code, rec.Code, tc.input.target)
				assert.Equal(t, tc.expected.message, rec.Body.String(), tc.input.target)
			}
		} else {
			if err := h.NewComment(c); assert.Error(t, err, tc.input.target) {
				assert.Equal(t, tc.expected.code, err.(*echo.HTTPError).Code, tc.input.target)
				assert.Equal(t, tc.expected.message, err.(*echo.HTTPError).Message, tc.input.target)
			}
		}
	}
	h.Shutdown()
}

func TestFetchComment (t *testing.T) {
	e := echo.New()
	h := NewHandler("mongodb://SEavenger:SEavenger@ds121225.mlab.com:21225/se_avengers_test", key)//, make(chan model.Notification))

	// Get the tweetID for testing
	tempTweet := []model.Tweet{}
	h.db.DB(h.dbName).C(model.TweetCollection).Find(nil).Sort("timestamp").All(&tempTweet)
	fromTweetIDForTesting := tempTweet[0].ID.Hex()

	testCases := []restHTTPTestCase {
		// test success
		restHTTPTestCase{
			true,
			false,
			restHTTPInput{"TestFetchCommentSuccess", ""},
			restHTTPExpected{http.StatusOK, `{"commentlist":"[]","totalcomment":"12"}`},
		},
		// test invalid tweetID
		restHTTPTestCase{
			false,
			true,
			restHTTPInput{"TestFetchCommentInvalidTweetID", ""},
			restHTTPExpected{http.StatusNotFound, "Invalid tweet ID"},
		},
	}
	// Run
	for _, tc := range testCases {
		// Setup
		req := httptest.NewRequest(echo.GET, "/", nil)
		req.Header.Set(echo.HeaderAuthorization, "Bearer " + getToken(tc.input.user))
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetPath("api/v1/fetchcomment")
		c.SetParamNames("tweet")
		if tc.isinvalidtweetid{
			c.SetParamValues("5a007fb6311bc348635124")
		}else{
			c.SetParamValues(fromTweetIDForTesting)
		}
		processJWTToken(c)

		// Assertion
		if tc.positive {
			if assert.NoError(t, h.FetchComment(c), tc.input.user + " fetching its comment") {
				assert.Equal(t, tc.expected.code, rec.Code, tc.input.user + " fetching its comment")
				assert.Equal(t, tc.expected.message, removeCommentList(rec.Body.String()), tc.input.user + " fetching its comment")
			}
		} else {
			if err := h.FetchComment(c); assert.Error(t, err, tc.input.user + " fetching its comment") {
				assert.Equal(t, tc.expected.code, err.(*echo.HTTPError).Code, tc.input.user + " fetching its comment")
				assert.Equal(t, tc.expected.message, err.(*echo.HTTPError).Message, tc.input.user + " fetching its comment")
			}
		}
	}
}

func getIDToken(id string) string {
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["id"] = id
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()

	Token, err := token.SignedString([]byte(key))
	if err != nil {
		panic(err)
	}

	return Token
}

func getToken(username string) string {
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["username"] = username
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()

	Token, err := token.SignedString([]byte(key))
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
		SigningKey: 	[]byte(key),
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

func removeCommentList(recorderString string) string {
	in := []byte(recorderString)
	var raw map[string]string
	json.Unmarshal(in, &raw)
	raw["commentlist"] = "[]"
	out, _ := json.Marshal(raw)
	return string(out)
}
