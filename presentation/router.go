package presentation

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func SetRoute(r *gin.Engine) *gin.Engine {
	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "OK",
		})
	})
	return r
}
