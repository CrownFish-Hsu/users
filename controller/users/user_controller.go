package users

import (
	"blog/domain/users"
	"blog/services"
	"blog/utils/errors"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"net/http"
	"strconv"
	"time"
)

func Register(c *gin.Context) {
	var user users.User

	if err := c.ShouldBindJSON(&user); err != nil {
		er := errors.NewBadRequestError("invalid json body")
		c.JSON(er.Status, er)
		return
	}

	res, err := services.CreateUser(user)
	if err != nil {
		c.JSON(err.Status, err)
		return
	}

	c.JSON(http.StatusOK, res)
}

func Login(c *gin.Context) {
	var user users.User

	if err := c.ShouldBindJSON(&user); err != nil {
		er := errors.NewBadRequestError("invalid json")
		c.JSON(er.Status, er)
		return
	}

	result, err := services.GetUser(user)
	if err != nil {
		c.JSON(err.Status, err)
		return
	}

	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.RegisteredClaims{
		Issuer:    strconv.Itoa(int(result.Id)),
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 72)),
	})

	token, ero := claims.SignedString([]byte("qwerty"))
	if ero != nil {
		er := errors.NewInternalServeError("login failed")
		er.Message = ero.Error()
		c.JSON(er.Status, er)
		return
	}

	c.SetCookie("jwt", token, 3600, "/", "localhost", false, true)
	c.JSON(http.StatusOK, result)
}

func Get(c *gin.Context) {
	cookie, err := c.Cookie("jwt")
	if err != nil {
		er := errors.NewBadRequestError("get cookie failed")
		c.JSON(er.Status, er)
		return
	}

	token, err := jwt.ParseWithClaims(cookie, &jwt.RegisteredClaims{}, func(*jwt.Token) (interface{}, error) {
		return []byte("qwerty"), nil
	})
	if err != nil {
		er := errors.NewBadRequestError("parse token failed")
		c.JSON(er.Status, er)
		return
	}

	claims := token.Claims.(*jwt.RegisteredClaims)
	userId, err := strconv.ParseInt(claims.Issuer, 10, 64)
	if err != nil {
		er := errors.NewBadRequestError("userid not int10")
		c.JSON(er.Status, er)
		return
	}

	result, ero := services.GetUserByID(userId)
	if ero != nil {
		c.JSON(ero.Status, ero)
		return
	}

	c.JSON(http.StatusOK, result)
}

func Logout(c *gin.Context) {
	c.SetCookie("jwt", "", -1, "", "", false, true)

	c.JSON(http.StatusOK, gin.H{
		"message": "success",
	})
}
