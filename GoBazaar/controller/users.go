package controller

import (
	"GoBazaar/structs"

	"github.com/gin-gonic/gin"
)

func UserRegister(in *gin.Context) structs.User {
	/*
		newUser := structs.User{
			FirstName:     "Nikhilesh",
			LastName:      "Gupta",
			Email:         "nikhilesh20072001om@gmail.com",
			Contact:       "9162227847",
			City:          "Isri",
			WalletBalance: 12000.00,
			ID:            2,
		}
	*/

	var newUser structs.User
	if err := in.BindJSON(&newUser); err != nil {
		return structs.User{}
	}
	return newUser
}
