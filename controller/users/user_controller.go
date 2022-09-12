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
