package singlerouter

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

const (
	singRoute = "singlerouter"
)

func SingleRouter(c *gin.Context) {
	c.String(http.StatusOK, singRoute)
}
