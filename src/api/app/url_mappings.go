package app

import (
	"github.com/ShreyanshKeshav33/Golang-Microservices/src/api/controllers/repositories"
	"github.com/ShreyanshKeshav33/Golang-Microservices/src/api/controllers/singlerouter"
)

func mapUrls() {
	router.GET("/singroute", singlerouter.SingleRouter)
	router.POST("/repositories", repositories.CreateRepo)
}
