package app

import "github.com/gin-gonic/gin"

var (
	route = gin.Default()
)

func StartApplication()  {
	mapUrls()
	route.Run(":8080")
}
