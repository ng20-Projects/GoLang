package handler

import (
	"GoBazaar/controller"
	"net/http"

	"github.com/gin-gonic/gin"
)

func HomeHandler(in *gin.Context) {
	message := controller.Welcome(in)
	in.IndentedJSON(http.StatusOK, message)
}

func RegisterUser(in *gin.Context) {
	user := controller.UserRegister(in)
	in.IndentedJSON(http.StatusCreated, user)
}
