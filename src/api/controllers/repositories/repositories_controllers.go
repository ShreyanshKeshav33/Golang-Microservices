package repositories

import (
	"net/http"

	"github.com/ShreyanshKeshav33/Golang-Microservices/src/api/domain/repositories"
	"github.com/ShreyanshKeshav33/Golang-Microservices/src/api/services"
	"github.com/ShreyanshKeshav33/Golang-Microservices/src/api/utils/errors"
	"github.com/gin-gonic/gin"
)

func CreateRepo(c *gin.Context) {
	var request repositories.CreateRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		apiErr := errors.NewBadRequestError("invalid json body")
		c.JSON(apiErr.Status(), apiErr)
		return
	}
	result, err := services.RepositoryService.CreateRepo(request)
	if err != nil {
		c.JSON(err.Status(), err)
		return
	}
	c.JSON(http.StatusCreated, result)
}
