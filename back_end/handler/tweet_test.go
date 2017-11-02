package handler

import (
	"fmt"
	"strings"
	"time"
	"net/http"
	"net/http/httptest"
	"testing"
	"gopkg.in/mgo.v2"
	//"gopkg.in/mgo.v2/bson"
	"github.com/labstack/echo"
	"github.com/dgrijalva/jwt-go"
	"github.com/stretchr/testify/assert"
)

func TestFetchOwnTweets (t *testing.T) {
	// test cases
	requestParam := []string {
		"TomRiddle",
		"JasonHe", 
		"MarsLee" }

	expectedJSON := []string {
		"{\"page\":\"0\",\"totalpage\":\"0\",\"totaltweets\":\"0\",\"tweetlist\":[]}",
		"{\"page\":\"1\",\"totalpage\":\"9\",\"totaltweets\":\"25\",\"tweetlist\":[{\"id\":\"59de695257a4370860b1eb9a\",\"owner\":\"JasonHe\",\"message\":\"This is a test message from JasonHe!\",\"timestamp\":\"2017-10-11T14:56:18.704-04:00\"},{\"id\":\"59d24577311bc337bfec6d06\",\"owner\":\"JasonHe\",\"message\":\"Hi, I am Jason He. Weather sucks.\",\"timestamp\":\"2017-10-02T09:56:07.892-04:00\"},{\"id\":\"59d24577311bc337bfec6d07\",\"owner\":\"JasonHe\",\"message\":\"Hello from Jason He.\",\"timestamp\":\"2017-10-02T09:56:07.892-04:00\"}]}" ,
		"{\"page\":\"1\",\"totalpage\":\"8\",\"totaltweets\":\"24\",\"tweetlist\":[{\"id\":\"59d24577311bc337bfec6d04\",\"owner\":\"MarsLee\",\"message\":\"Hi, I am Chih-Yin Lee. Weather sucks.\",\"timestamp\":\"2017-10-02T09:56:07.892-04:00\"},{\"id\":\"59d24577311bc337bfec6d05\",\"owner\":\"MarsLee\",\"message\":\"Hello from Chih-Yin Lee.\",\"timestamp\":\"2017-10-02T09:56:07.892-04:00\"},{\"id\":\"59d24577311bc337bfec6d0f\",\"owner\":\"MarsLee\",\"message\":\"Hi, I am Chih-Yin Lee. Weather sucks.\",\"timestamp\":\"2017-10-02T09:56:07.892-04:00\"}]}"}

	session, err := mgo.Dial("mongodb://SEavenger:SEavenger@ds149324.mlab.com:49324/se_avengers")
	if err != nil {
		panic(err)
	}

	// Setup
	e := echo.New()
	h := &Handler{session}

	// test empty request parameter
	//c.SetParamValues("")

	// Run
	for i, rp := range requestParam {
		// Setup
		req := httptest.NewRequest(echo.GET, "/?page=1&perpage=3", nil)
		req.Header.Set(echo.HeaderAuthorization, "Bearer " + getToken(rp))
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetPath("api/v1/tweetlist")
		c.SetParamNames("username")
		c.SetParamValues(rp)
		processJWTToken(c)

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
		`{"message":"BBBBBBBBQ"}`,
	}

	expectedJSON := []string {
		"{\"owner\":\"59d24577311bc337bfec6cf9\",\"message\":\"BBBBBBBBQ\"}",
	}
	
	// Run
	for i, rp := range requestParam {
		// Setup
		req := httptest.NewRequest(echo.POST, "/", strings.NewReader(rp))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		req.Header.Set(echo.HeaderAuthorization, "Bearer " + getIDToken("59d24577311bc337bfec6cf9"))
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetPath("api/v1/newTweet")
		c.SetParamNames("id")
		c.SetParamValues(rp)
		processJWTToken(c)

		// Assertion
		if assert.NoError(t, h.NewTweet(c)) {
			assert.Equal(t, http.StatusOK, rec.Code)
			assert.Equal(t, expectedJSON[i], rec.Body.String())
		}
	}
}

func TestFetchTweetTimeLine (t *testing.T) {
	// test cases
	requestParam := []string {
		"TomRiddle",
		"JasonHe", 
		"MarsLee" }

	expectedJSON := []string {
		"{\"page\":\"0\",\"totalpage\":\"0\",\"totaltweets\":\"0\",\"tweetlist\":[]}",
		"{\"page\":\"1\",\"totalpage\":\"9\",\"totaltweets\":\"25\",\"tweetlist\":[{\"id\":\"59de695257a4370860b1eb9a\",\"owner\":\"JasonHe\",\"message\":\"This is a test message from JasonHe!\",\"timestamp\":\"2017-10-11T14:56:18.704-04:00\"},{\"id\":\"59d24577311bc337bfec6d06\",\"owner\":\"JasonHe\",\"message\":\"Hi, I am Jason He. Weather sucks.\",\"timestamp\":\"2017-10-02T09:56:07.892-04:00\"},{\"id\":\"59d24577311bc337bfec6d07\",\"owner\":\"JasonHe\",\"message\":\"Hello from Jason He.\",\"timestamp\":\"2017-10-02T09:56:07.892-04:00\"}]}" ,
		"{\"page\":\"1\",\"totalpage\":\"8\",\"totaltweets\":\"24\",\"tweetlist\":[{\"id\":\"59d24577311bc337bfec6d04\",\"owner\":\"MarsLee\",\"message\":\"Hi, I am Chih-Yin Lee. Weather sucks.\",\"timestamp\":\"2017-10-02T09:56:07.892-04:00\"},{\"id\":\"59d24577311bc337bfec6d05\",\"owner\":\"MarsLee\",\"message\":\"Hello from Chih-Yin Lee.\",\"timestamp\":\"2017-10-02T09:56:07.892-04:00\"},{\"id\":\"59d24577311bc337bfec6d0f\",\"owner\":\"MarsLee\",\"message\":\"Hi, I am Chih-Yin Lee. Weather sucks.\",\"timestamp\":\"2017-10-02T09:56:07.892-04:00\"}]}"}

	session, err := mgo.Dial("mongodb://SEavenger:SEavenger@ds149324.mlab.com:49324/se_avengers")
	if err != nil {
		panic(err)
	}

	// Setup
	e := echo.New()
	h := &Handler{session}

	// test empty request parameter
	//c.SetParamValues("")

	// Run
	for i, rp := range requestParam {
		// Setup
		req := httptest.NewRequest(echo.GET, "/?page=1&perpage=3", nil)
		req.Header.Set(echo.HeaderAuthorization, "Bearer " + getToken(rp))
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetPath("api/v1/tweetlist")
		c.SetParamNames("username")
		c.SetParamValues(rp)
		processJWTToken(c)

		// Assertion
		if assert.NoError(t, h.FetchTweets(c)) {
			assert.Equal(t, http.StatusOK, rec.Code)
			assert.Equal(t, expectedJSON[i], rec.Body.String())
		}
	}
	
}


func getIDToken(id string) string {
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["id"] = id
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()

	Token, err := token.SignedString([]byte(Key))
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