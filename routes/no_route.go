package routes

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func NoRoute(c *gin.Context) {
	c.String(http.StatusNotFound, "Not Found")
}
