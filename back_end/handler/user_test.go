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
		testCase{
			true,
			input{"", `{"username":"testSignup","firstname":"test","password":"test","email":"testSignup@gmail.com","bio":"testtest","tag":"Einstein"}`},
			expected{http.StatusCreated, `{"username":"testSignup","firstname":"test","email":"testSignup@gmail.com","bio":"testtest","tag":"Einstein"}`},
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
			err := h.DB.DB(h.DBName).C(model.UserCollection).Remove(bson.M{"username": "testSignup"})
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
				h.DB.DB(h.DBName).C(model.UserCollection).Find(bson.M{"username": tc.input.target}).One(&user)
				assert.Equal(t, []string{tc.input.user}, user.Followers, tc.input.user + " follows " + tc.input.target)

				h.DB.DB(h.DBName).C(model.UserCollection).Update(bson.M{"username": tc.input.target}, bson.M{"$pull": bson.M{"followers": tc.input.user}})
				h.DB.DB(h.DBName).C(model.UserCollection).Update(bson.M{"username": tc.input.user}, bson.M{"$pull": bson.M{"following": tc.input.target}})
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

func TestUnfollow(t *testing.T) {
	e := echo.New()
	h := NewHandler("mongodb://SEavenger:SEavenger@ds121225.mlab.com:21225/se_avengers_test")

	testCases := []testCase {
		// test success
		testCase{
			true,
			input{"testUnfollow", "testUnfollow_2"},
			expected{http.StatusOK, `["testUnfollow_1"]`},
		},
		// test fail
		testCase{
			false,
			input{"testUnfollow", "testUnfollow_not_exist"},
			expected{http.StatusNotFound, "User does not exist."},
		},
	}

	for _, tc := range testCases {
		req := httptest.NewRequest(echo.POST, "/", nil)
		req.Header.Set(echo.HeaderAuthorization, "Bearer " + getToken(tc.input.user))
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetPath("/api/v1/unfollow")
		c.SetParamNames("username")
		c.SetParamValues(tc.input.target)
		processJWTToken(c)

		if tc.positive {
			if assert.NoError(t, h.Unfollow(c), tc.input.user + " unfollows " + tc.input.target) {
				assert.Equal(t, tc.expected.code, rec.Code, tc.input.user + " unfollows " + tc.input.target)
				assert.Equal(t, tc.expected.message, rec.Body.String(), tc.input.user + " unfollows " + tc.input.target)

				user := model.User{}
				h.DB.DB(h.DBName).C(model.UserCollection).Find(bson.M{"username": tc.input.target}).One(&user)
				assert.Equal(t, []string{}, user.Followers, tc.input.user + " unfollows " + tc.input.target)

				h.DB.DB(h.DBName).C(model.UserCollection).Update(bson.M{"username": tc.input.target}, bson.M{"$addToSet": bson.M{"followers": tc.input.user}})
				h.DB.DB(h.DBName).C(model.UserCollection).Update(bson.M{"username": tc.input.user}, bson.M{"$addToSet": bson.M{"following": tc.input.target}})
			}
		} else {
			if err := h.Unfollow(c); assert.Error(t, err, tc.input.user + " unfollows " + tc.input.target) {
				assert.Equal(t, tc.expected.code, err.(*echo.HTTPError).Code, tc.input.user + " unfollows " + tc.input.target)
				assert.Equal(t, tc.expected.message, err.(*echo.HTTPError).Message, tc.input.user + " unfollows " + tc.input.target)
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
	e := echo.New()
	h := NewHandler("mongodb://SEavenger:SEavenger@ds121225.mlab.com:21225/se_avengers_test")

	testCases := []testCase {
		// test success
		testCase{
			true,
			input{"testUpdateUserInfo", `{"firstname":"testUpdateUserInfo_mod","lastname":"mod","bio":"mod","tag":"mod"}`},
			expected{http.StatusOK, `{"username":"testUpdateUserInfo","firstname":"testUpdateUserInfo_mod","lastname":"mod","email":"testUpdateUserInfo@gmail.com","bio":"mod","tag":"mod"}`},
		},
		testCase{
			true,
			input{"testUpdateUserInfo", `{"firstname":"testUpdateUserInfo","lastname":"","bio":"","tag":""}`},
			expected{http.StatusOK, `{"username":"testUpdateUserInfo","firstname":"testUpdateUserInfo","email":"testUpdateUserInfo@gmail.com"}`},
		},
		// test fail
		testCase{
			false,
			input{"testUpdateUserInfo_empty_firstname", `{"firstname":"","lastname":"","bio":"","tag":"testUpdate"}`},
			expected{http.StatusBadRequest, "Firstname cannot be empty."},
		},
		testCase{
			false,
			input{"testUpdateUserInfo_not_exist", `{"firstname":"testUpdateUserInfo_not_exist","lastname":"","bio":"","tag":""}`},
			expected{http.StatusNotFound, "User does not exist."},
		},
	}

	for _, tc := range testCases {
		req := httptest.NewRequest(echo.POST, "/", strings.NewReader(tc.input.target))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		req.Header.Set(echo.HeaderAuthorization, "Bearer " + getToken(tc.input.user))
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetPath("/api/v1/updateUserInfo")
		processJWTToken(c)

		if tc.positive {
			if assert.NoError(t, h.UpdateUserInfo(c), tc.input.target) {
				assert.Equal(t, tc.expected.code, rec.Code, tc.input.target)
				assert.Equal(t, tc.expected.message, rec.Body.String(), tc.input.target)
			}
		} else {
			if err := h.UpdateUserInfo(c); assert.Error(t, err, tc.input.target) {
				assert.Equal(t, tc.expected.code, err.(*echo.HTTPError).Code, tc.input.target)
				assert.Equal(t, tc.expected.message, err.(*echo.HTTPError).Message, tc.input.target)
			}
		}
	}
	h.DB.Close()
}

func TestUpdateProfilePicture(t *testing.T) {
	e := echo.New()
	h := NewHandler("mongodb://SEavenger:SEavenger@ds121225.mlab.com:21225/se_avengers_test")

	oversizedImage := [10485760 + 1]byte{}
	for i := 0; i < 10485760 + 1; i++ {
		oversizedImage[i] = 'A'
	}

	testCases := []testCase {
		// test success
		testCase{
			true,
			input{"testUpdateProfilePicture", `{"picture":"iVBORw0KGgoAAAANSUhEUgAAACAAAAAgCAMAAABEpIrGAAAAjVBMVEXp6en///+8vLwPZ66L1fPY` + 
														  `//9Xo95eqeH4/f8wg8AQaK6U2flirOJkruJ4wutnseTf8/+F0PF0v+p5y+7l9/+CzO9puefw+//2` + 
														  `9vZ8xux/ye3X8f/b29trteWt6f+V0/O48P8YbbCo3/WVrcC1zeHFxcWo5/+h4f/K+/++9P+FwuzM` + 
														  `/f/H6fgjd7hnvekVvPy4AAABSElEQVQ4jXXT23KDIBCAYbY5WaNiqEGJWoOao0ne//G6gBjphH/G` + 
														  `GVk/1ysJQC/DxcdC2QMQ6J/tylP77IEUsm3WnppWFoQsVuuNp/VqQRAMm8jTZhhBjEX4RhS7GRAO` + 
														  `ah7dy7S8/xPRECpwjSmlUamKqFN8taCmcYbP0zeosQmo0wjscg1qasBNnWiWIshsKdPgZkBtwLys` + 
														  `1Dss4JzXuQtyPRwB10CtyN7fUIAbcOSMMZ6rcW7L7np4NIBZkPMpppuDQAPmZsBvgvEAFwc8cfsd` + 
														  `gRAJ0+A1PRIqC/Q9aPBygDDgbACuCGCq0eBswAETgRMEQk1HsMcOD3CFOODUgIsG+xzmBUJNLwbs` + 
														  `8Bavxxw0arKzwLT/mWVGI1juPC01kMl26WmbSAQdCk+J7Agpqk4+vz8mZVcVBIrq9OXpVBX4dwP+` + 
														  `n74KgD9PVzqtRZGhswAAAABJRU5ErkJggg=="}`},
			expected{http.StatusCreated, `{"username":"testUpdateProfilePicture","firstname":"testUpdateProfilePicture","email":"testUpdateProfilePicture@gmail.com",` + 
											   `"picture":"iVBORw0KGgoAAAANSUhEUgAAACAAAAAgCAMAAABEpIrGAAAAjVBMVEXp6en///+8vLwPZ66L1fPY` + 
														  `//9Xo95eqeH4/f8wg8AQaK6U2flirOJkruJ4wutnseTf8/+F0PF0v+p5y+7l9/+CzO9puefw+//2` + 
														  `9vZ8xux/ye3X8f/b29trteWt6f+V0/O48P8YbbCo3/WVrcC1zeHFxcWo5/+h4f/K+/++9P+FwuzM` + 
														  `/f/H6fgjd7hnvekVvPy4AAABSElEQVQ4jXXT23KDIBCAYbY5WaNiqEGJWoOao0ne//G6gBjphH/G` + 
														  `GVk/1ysJQC/DxcdC2QMQ6J/tylP77IEUsm3WnppWFoQsVuuNp/VqQRAMm8jTZhhBjEX4RhS7GRAO` + 
														  `ah7dy7S8/xPRECpwjSmlUamKqFN8taCmcYbP0zeosQmo0wjscg1qasBNnWiWIshsKdPgZkBtwLys` + 
														  `1Dss4JzXuQtyPRwB10CtyN7fUIAbcOSMMZ6rcW7L7np4NIBZkPMpppuDQAPmZsBvgvEAFwc8cfsd` + 
														  `gRAJ0+A1PRIqC/Q9aPBygDDgbACuCGCq0eBswAETgRMEQk1HsMcOD3CFOODUgIsG+xzmBUJNLwbs` + 
														  `8Bavxxw0arKzwLT/mWVGI1juPC01kMl26WmbSAQdCk+J7Agpqk4+vz8mZVcVBIrq9OXpVBX4dwP+` + 
											  			  `n74KgD9PVzqtRZGhswAAAABJRU5ErkJggg=="}`},
		},
		testCase{
			true,
			input{"testUpdateProfilePicture", `{"picture":""}`},
			expected{http.StatusCreated, `{"username":"testUpdateProfilePicture","firstname":"testUpdateProfilePicture","email":"testUpdateProfilePicture@gmail.com"}`},
		},
		// test fail
		testCase{
			false,
			input{"testUpdateProfilePicture_oversized_image", `{"picture":"` + string(oversizedImage[:]) + `"}`},
			expected{http.StatusBadRequest, "Image must be smaller than 10 MB."},
		},
		testCase{
			false,
			input{"testUpdateProfilePicture_not_exist", `{"picture":""}`},
			expected{http.StatusNotFound, "User does not exist."},
		},
	}

	for _, tc := range testCases {
		req := httptest.NewRequest(echo.POST, "/", strings.NewReader(tc.input.target))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		req.Header.Set(echo.HeaderAuthorization, "Bearer " + getToken(tc.input.user))
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetPath("/api/v1/updateProfilePic")
		processJWTToken(c)

		if tc.positive {
			if assert.NoError(t, h.UpdateProfilePicture(c), tc.input.target) {
				assert.Equal(t, tc.expected.code, rec.Code, tc.input.target)
				assert.Equal(t, tc.expected.message, rec.Body.String(), tc.input.target)
			}
		} else {
			if err := h.UpdateProfilePicture(c); assert.Error(t, err, tc.input.target) {
				assert.Equal(t, tc.expected.code, err.(*echo.HTTPError).Code, tc.input.target)
				assert.Equal(t, tc.expected.message, err.(*echo.HTTPError).Message, tc.input.target)
			}
		}
	}
	h.DB.Close()
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