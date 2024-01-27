package controllers

import (
	"net/http"
	"strconv" //string conversion

	"github.com/ShreyanshKeshav33/Golang-Microservices/mvc/services"
	"github.com/ShreyanshKeshav33/Golang-Microservices/mvc/utils"
	"github.com/gin-gonic/gin"
)

func GetUser(c *gin.Context) {
	userId, err := strconv.ParseInt(c.Param("user_id"), 10, 64)
	//log.Printf("About to process user_id %v", userId)
	if err != nil {
		apiErr := &utils.ApplicationError{
			Message:    "user_id must be a number",
			StatusCode: http.StatusBadRequest,
			Code:       "bad request",
		}
		c.JSON(apiErr.StatusCode, apiErr)
		return
	}
	user, apiErr := services.GetUser(userId)
	if apiErr != nil {
		c.JSON(apiErr.StatusCode, apiErr)
		//Handle the error and return to the client
		return
	}
	// return user to client
	c.JSON(http.StatusOK, user)
}
