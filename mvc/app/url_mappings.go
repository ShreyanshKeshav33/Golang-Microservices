package app

import (
	"github.com/ShreyanshKeshav33/Golang-Microservices/mvc/controllers"
)

func mapUrls() {
	router.GET("/users/:user_id", controllers.GetUser)
}
