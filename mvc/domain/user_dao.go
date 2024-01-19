package domain

import (
	"fmt"
	"net/http"

	"github.com/ShreyanshKeshav33/Golang-Microservices/mvc/utils"
)

var (
	users = map[int64]*User{
		123: {Id: 1, FirstName: "Shreyansh", LastName: "Keshav", Email: "shreykeshav33@gmail.com"},
	}
)

func GetUser(userId int64) (*User, *utils.ApplicationError) {
	if user := users[userId]; user != nil {
		return user, nil
	}
	return nil, &utils.ApplicationError{
		Message:    fmt.Sprintf("user %vwas not found", userId),
		StatusCode: http.StatusNotFound,
		Code:       "not found",
	}
}
