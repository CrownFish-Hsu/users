package app

import "blog/controller/users"

func mapUrls()  {
	route.POST("/api/register", users.Register)

	route.POST("/api/login", users.Login)

	//route.GET("/api/user", users.Get)

	//route.GET("/api/logout", users.logout)
}