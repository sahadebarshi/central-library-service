package controller

import (
	"net/http"
	"rest/cls/handler"
	middlewares "rest/cls/middleware"
	"sync/atomic"

	"github.com/gin-gonic/gin"
)

func RouteController(server *gin.Engine) {

	authenticate := server.Group("/")

	authenticate.Use(middlewares.Authenticated)

	authenticate.POST("/users/details", handler.GetBookUser)

	//server.Run(":8085")

}

func HealthController(server *gin.Engine, mtlsServerStatus *atomic.Bool) {
	server.GET("/health", func(c *gin.Context) {
		if mtlsServerStatus.Load() {
			c.JSON(http.StatusOK, gin.H{"status": "ok"})
		} else {
			c.JSON(http.StatusServiceUnavailable, gin.H{"error": "mTLS server is down"})
		}
	})
}
