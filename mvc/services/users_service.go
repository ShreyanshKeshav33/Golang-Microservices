package services

import (
	"github.com/ShreyanshKeshav33/Golang-Microservices/mvc/domain"
	"github.com/ShreyanshKeshav33/Golang-Microservices/mvc/utils"
)

func GetUser(userId int64) (*domain.User, *utils.ApplicationError) {
	return domain.GetUser(userId)
}
