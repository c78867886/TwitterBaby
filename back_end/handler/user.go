package handler

import (
	"time"
	"net/http"
	"gopkg.in/mgo.v2"
	"github.com/dgrijalva/jwt-go"
	"gopkg.in/mgo.v2/bson"
	"github.com/labstack/echo"
	"model"
)

func (h *Handler) Signup(c echo.Context) (err error) {
	// Bind
	u := &model.User{}
	if err = c.Bind(u); err != nil {
		return
	}

	u.ID = bson.NewObjectId()
	u.Followers = []string{}
	u.Following = []string{}

	// Validate
	if u.FirstName == "" || u.Email == "" || u.Password == "" {
		return &echo.HTTPError{Code: http.StatusBadRequest, Message: "Invalid firstname, email or password."}
	}

	// Save user
	db := h.DB.Clone()
	defer db.Close()
	
	err = db.DB("se_avengers").C("users").Insert(u);
	if err != nil {
		return
	}

	// Don't send password
	u.Password = ""

	return c.JSON(http.StatusCreated, u)
}

func (h *Handler) Login(c echo.Context) (err error) {
	// Bind
	u := new(model.User)
	if err = c.Bind(u); err != nil {
		return
	}

	// Find user
	db := h.DB.Clone()
	defer db.Close()

	err = db.DB("se_avengers").C("users").Find(bson.M{"email": u.Email, "password": u.Password}).One(u)
	if err != nil {
		if err == mgo.ErrNotFound {
			return &echo.HTTPError{Code: http.StatusUnauthorized, Message: "Invalid email or password."}
		}
		return
	}

	// JWT

	// Create token
	token := jwt.New(jwt.SigningMethodHS256)

	// Set claims
	claims := token.Claims.(jwt.MapClaims)
	claims["id"] = u.ID
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()

	// Generate encoded token and send it as response
	u.Token, err = token.SignedString([]byte(Key))
	if err != nil {
		return err
	}

	// Don't send password
	u.Password = ""

	return c.JSON(http.StatusOK, u)
}

func (h *Handler) FetchUserInfo (c echo.Context) (err error) {
	userID := c.Param("user")

	db := h.DB.Clone()
	defer db.Close()

	// Retrieve user info from database
	user := model.User{}
	err = db.DB("se_avengers").C("users").FindId(bson.ObjectIdHex(userID)).One(&user)
	if err != nil {
		if err == mgo.ErrNotFound {
			return &echo.HTTPError{Code: http.StatusNotFound, Message: "User does not exist."}
		}
		return
	}

	// Don't send password
	user.Password = ""
	
	return c.JSON(http.StatusOK, user)
}

func (h *Handler) UpdateUserInfo (c echo.Context) (err error) {
	userID := userIDFromToken(c)

	user := model.User{}
	if err = c.Bind(user); err != nil {
		return
	}

	if user.ID != "" {
		return &echo.HTTPError{Code: http.StatusUnauthorized, Message: "User ID cannot be modified."}
	}

	db := h.DB.Clone()
	defer db.Close()

	err = db.DB("twitter").C("users").UpdateId(bson.ObjectIdHex(userID), bson.M{"followers": userID})
	if err != nil {
		if err == mgo.ErrNotFound {
			return &echo.HTTPError{Code: http.StatusNotFound, Message: "User does not exist."}
		}
	}

	// Don't send password
	user.Password = ""
	
	//return c.JSON(http.StatusOK, user)
	return c.NoContent(http.StatusNotImplemented)
}

// Follow : Add a specific user to the current user's following set
func (h *Handler) Follow(c echo.Context) (err error) {
	userID := userIDFromToken(c)
	id := c.Param("id")

	db := h.DB.Clone()
	defer db.Close()

	err = db.DB("se_avengers").C("users").UpdateId(bson.ObjectIdHex(userID), bson.M{"$addToSet": bson.M{"following": id}})
	if err != nil {
		if err == mgo.ErrNotFound {
			return &echo.HTTPError{Code: http.StatusNotFound, Message: "User does not exist."}
		}
	}
	
	user := model.User{}
	err = db.DB("se_avengers").C("users").FindId(bson.ObjectIdHex(userID)).One(&user)

	return c.JSON(http.StatusOK, user.Following)
}

func userIDFromToken(c echo.Context) string {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)

	return claims["id"].(string)
}