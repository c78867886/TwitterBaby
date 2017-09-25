package handler

import (
	"net/http"
	"gopkg.in/mgo.v2/bson"
	"github.com/labstack/echo"
	"model"
)

/*func userIDFromToken(c echo.Context) string {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	return claims["id"].(string)
}*/
type userContainer struct {
	FirstName 	string			`json:"firstname"`
	LastName	string			`json:"lastname"`
	Password 	string			`json:"password"`
	Email 		string			`json:"email"`
}

func (h *Handler) FetchUserInfo (c echo.Context) (err error) {
	userID := c.Param("user")

	// Retrieve user infomation from database
	userInfo := []*model.User{}
	db := h.DB.Clone()

	err = db.DB("se_avangers").C("users").Find(bson.M{"useriddev": userID}).All(&userInfo)
	if err != nil {
		return
	}
	defer db.Close()
	
	res := []userContainer{}
	for _, u := range userInfo {
		firstname := u.FirstName
		lastname := u.LastName
		password := u.Password
		email := u.Email
		res = append(res, userContainer{
			FirstName:firstname, LastName:lastname, Password:password , Email:email })
	}
	
	return c.JSON(http.StatusOK, res)
}