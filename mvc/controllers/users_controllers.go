package controllers

import (
	"encoding/json"
	"net/http"
	"strconv" //string conversion

	"github.com/ShreyanshKeshav33/Golang-Microservices/mvc/services"
	"github.com/ShreyanshKeshav33/Golang-Microservices/mvc/utils"
)

func GetUser(resp http.ResponseWriter, req *http.Request) {
	userId, err := strconv.ParseInt(req.URL.Query().Get("user_id"), 10, 64)
	//log.Printf("About to process user_id %v", userId)
	if err != nil {
		apiErr := &utils.ApplicationError{
			Message:    "user_id must be a number",
			StatusCode: http.StatusBadRequest,
			Code:       "bad request",
		}
		jsonValue, _ := json.Marshal(apiErr)
		resp.WriteHeader(apiErr.StatusCode)
		resp.Write(jsonValue)
		//if error, then just return Bad Request to the client
		return
	}
	user, apiErr := services.GetUser(userId)
	if apiErr != nil {
		resp.WriteHeader(http.StatusNotFound)
		resp.Write([]byte(apiErr.Message))
		//Handle the error and return to the client
		return
	}
	// return user to client
	jsonValue, _ := json.Marshal(user)
	resp.Write(jsonValue)
}
