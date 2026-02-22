package handler

import (
	"net/http"
	model "rest/cls/Model"

	"github.com/gin-gonic/gin"
)

type BookUserInterface interface {
	Save() error
}

type BookUserWrapper struct {
	model.BookUser
}

func (buw *BookUserWrapper) Save() error {
	return nil
}

/*func HandaleBookUser(b BookUserInterface) {

}*/

func GetBookUser(context *gin.Context) {
	var userData BookUserWrapper
	err := context.ShouldBindJSON(&userData.BookUser)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"error_message": err.Error(),
		})
		return
	}

	context.JSON(http.StatusCreated, userData.BookUser)
}
