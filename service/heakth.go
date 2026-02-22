package service

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetHealth(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{
		"message": "Health is fine",
	})
}
