package app

import "blog/controller/users"

func mapUrls() {
	route.POST("/api/register", users.Register)

	route.POST("/api/login", users.Login)

	route.GET("/api/user", users.Get)

	route.POST("/api/logout", users.Logout)
}
