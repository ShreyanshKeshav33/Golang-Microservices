package app

import (
	"net/http"

	"github.com/ShreyanshKeshav33/Golang-Microservices/mvc/controllers"
)

func StartApp() {
	http.HandleFunc("/users", controllers.GetUser)

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}
