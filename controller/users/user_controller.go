package users

import (
	"blog/domain/users"
	"blog/services"
	"blog/utils/errors"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Register(c *gin.Context)  {
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
