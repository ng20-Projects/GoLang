package controller

import (
	"GoBazaar/structs"

	"github.com/gin-gonic/gin"
)

func Welcome(in *gin.Context) structs.WelcomeMsg {
	msg := structs.WelcomeMsg{
		Msg: "Welcome to GoBazzar!! Your own desi shop",
	}
	return msg
}
