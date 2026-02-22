package middlewares

import "github.com/gin-gonic/gin"

func Authenticated(context *gin.Context) {

	println("Request method is: ", context.Request.Method)

	context.Next()

}
