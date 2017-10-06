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

// Signup : Create an user instance.
//			URL: "/api/v1/signup"
//			Method: POST
//			Return 201 Created on success, along with the user data.
//			Return 400 Bad Request if one of username, firstname, email, password is empty, or username, email already used.
func (h *Handler) Signup(c echo.Context) (err error) {
	// Bind
	u := &model.User{}
	if err = c.Bind(u); err != nil {
		return
	}

	db := h.DB.Clone()
	defer db.Close()

	u.ID = bson.NewObjectId()
	u.Followers = []string{}
	u.Following = []string{}

	// Validate
	if u.Username == "" || u.FirstName == "" || u.Email == "" || u.Password == "" {
		return &echo.HTTPError{Code: http.StatusBadRequest, Message: "Invalid username, firstname, email or password."}
	}

	varify := &model.User{}
	err = db.DB("se_avengers").C("users").Find(bson.M{"$or": []bson.M{bson.M{"username": u.Username}, bson.M{"email": u.Email}}}).One(varify)
	if err == nil {
		return &echo.HTTPError{Code: http.StatusBadRequest, Message: "Username or email already used."}
	}

	// Save user
	err = db.DB("se_avengers").C("users").Insert(u);
	if err != nil {
		return
	}

	// Don't send password
	u.Password = ""

	return c.JSON(http.StatusCreated, u)
}

// Login : Login to an account associated with the email address and the password.
//		   URL: "/api/v1/login"
//		   Method: POST
//		   Return 200 OK on success, along with the user data, which now contains a token.
//		   Return 401 Unauthorized if an account associated with the email address and password is not found.
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
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()
	claims["username"] = u.Username

	// Generate encoded token and send it as response
	u.Token, err = token.SignedString([]byte(Key))
	if err != nil {
		return err
	}

	// Don't send password
	u.Password = ""

	return c.JSON(http.StatusOK, u)
}

// FetchUserInfo : Return user info for a specific user, and whether it is followed by the current user.
//				   # Response does not include the full list of followers and following, only the counts.
//				   URL: "/api/v1/userInfo/:id", "/api/v1/userInfo?username=:username"
//				   Method: GET
//				   Return 200 OK on success.
//				   Return 404 Not Found if the user is not in the database.
func (h *Handler) FetchUserInfo (c echo.Context) (err error) {
	selfID := userIDFromToken(c)

	db := h.DB.Clone()
	defer db.Close()

	// Retrieve user info from database
	user := model.User{}
	if c.Path() == "/api/v1/userInfo/:id" {
		err = db.DB("se_avengers").C("users").FindId(bson.ObjectIdHex(c.Param("id"))).One(&user)
	} else {
		err = db.DB("se_avengers").C("users").Find(bson.M{"username": c.QueryParam("username")}).One(&user)
	}
	if err != nil {
		if err == mgo.ErrNotFound {
			return &echo.HTTPError{Code: http.StatusNotFound, Message: "User does not exist."}
		}
		return
	}

	// Don't send password
	user.Password = ""

	var container struct {
		U				model.User	`json:"userinfo"`
		FollowerCount	int			`json:"followercount"`
		FollowingCount	int			`json:"followingcount"`
		Followed		bool		`json:"followed"`
	}
	container.U = user
	container.FollowerCount = len(user.Followers)
	container.FollowingCount = len(user.Following)

	container.Followed = false
	for i := range user.Followers {
		if user.Followers[i] == selfID {
			container.Followed = true
			break
		}
	}

	container.U.Followers = nil
	container.U.Following = nil

	return c.JSON(http.StatusOK, container)
}

// Follow : Add a specific user ID to the current user's following set, and add current user to that user's follower list.
//			URL: "/api/v1/follow/:id"
//			Method: POST
//			Return 200 OK on success, along with the user's following list.
//			Return 404 Not Found if the user is not in the database.
func (h *Handler) Follow(c echo.Context) (err error) {
	userID := userIDFromToken(c)
	id := c.Param("id")

	db := h.DB.Clone()
	defer db.Close()

	// Add id to self following list
	err = db.DB("se_avengers").C("users").UpdateId(bson.ObjectIdHex(userID), bson.M{"$addToSet": bson.M{"following": id}})
	if err != nil {
		if err == mgo.ErrNotFound {
			return &echo.HTTPError{Code: http.StatusNotFound, Message: "User does not exist."}
		}
	}

	// Add self to id's follower list
	err = db.DB("se_avengers").C("users").UpdateId(bson.ObjectIdHex(id), bson.M{"$addToSet": bson.M{"followers": userID}})
	if err != nil {
		return
	}
	
	user := model.User{}
	err = db.DB("se_avengers").C("users").FindId(bson.ObjectIdHex(userID)).One(&user)

	return c.JSON(http.StatusOK, user.Following)
}

// ShowFollower : Return the follower list for a specific user, along with some followers info.
//				  URL: "/api/v1/showFollower/:id"
//				  Method: GET
//				  Return 200 OK on success.
//				  Return 404 Not Found if the user is not in the database.
func (h *Handler) ShowFollower(c echo.Context) (err error) {
	userID := c.Param("id")

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

	type followerData struct {
		ID			bson.ObjectId	`json:"id" bson:"_id"`
		Username	string			`json:"username" bson:"username"`
		FirstName	string			`json:"firstname" bson:"firstname"`
		LastName	string			`json:"lastname,omitempty" bson:"lastname,omitempty"`
		Bio			string			`json:"bio,omitempty" bson:"bio,omitempty"`
	}
	container := []followerData{}

	for _, f := range user.Followers {
		follower := followerData{}
		err = db.DB("se_avengers").C("users").FindId(bson.ObjectIdHex(f)).One(&follower)
		if err != nil {
			return
		}
		container = append(container, follower)
	}

	return c.JSON(http.StatusOK, &container)
}

// ShowFollowing : Return the following list for a specific user, along with some followings info.
//				   URL: "/api/v1/showFollowing/:id"
//				   Method: GET
//				   Return 200 OK on success.
//				   Return 404 Not Found if the user is not in the database.
func (h *Handler) ShowFollowing(c echo.Context) (err error) {
	userID := c.Param("id")
	
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

	type followingData struct {
		ID			bson.ObjectId	`json:"id" bson:"_id"`
		Username	string			`json:"username" bson:"username"`
		FirstName	string			`json:"firstname" bson:"firstname"`
		LastName	string			`json:"lastname,omitempty" bson:"lastname,omitempty"`
		Bio			string			`json:"bio,omitempty" bson:"bio,omitempty"`
	}
	container := []followingData{}

	for _, f := range user.Following {
		following := followingData{}
		err = db.DB("se_avengers").C("users").FindId(bson.ObjectIdHex(f)).One(&following)
		if err != nil {
			return
		}
		container = append(container, following)
	}

	return c.JSON(http.StatusOK, &container)
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

func userIDFromToken(c echo.Context) string {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)

	return claims["id"].(string)
}
