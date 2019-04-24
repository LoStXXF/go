package controller

import (
	"github.com/gin-gonic/gin"
)

type RetJSON struct {
	Status int
	Details string
}

func Run() {
	route := gin.Default()
	user := route.Group("/user")
	{
		user.POST("/login", api_login)
		user.POST("/register", api_register)
		user.GET("/get_personal_info/:id", api_update_personal_info)
		user.GET("/get_verification_code/:phone", api_get_verification_code)
	}
	route.Run(":8081") // listen and serve on 0.0.0.0:8080
}