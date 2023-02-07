package controller

import (
	"net/http"
	"projeto/FazTudo/dto"
	"projeto/FazTudo/services/loginService"

	"github.com/gin-gonic/gin"
)

func CreateLogin(ctx *gin.Context) {
	var login dto.LoginDTO
	ctx.BindJSON(&login)

	exist := loginService.NewLoginSerice().ValidateCredential(login)

	if exist {
		ctx.JSON(http.StatusConflict, gin.H{})
	} else {
		token, err := loginService.NewLoginSerice().CreateCredential(login)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{})
		} else {
			var bearer string = "Bearer " + token
			ctx.Writer.Header().Set("Authorization", bearer)
			ctx.JSON(http.StatusCreated, gin.H{})
		}
	}
}

func ValidateCrendential(ctx *gin.Context) {
	var login dto.LoginDTO
	ctx.BindJSON(&login)

	isValid := loginService.NewLoginSerice().ValidateCredential(login)

	if isValid {

		token, err := loginService.NewLoginSerice().CreateJWT(login)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{})
		}

		var bearer string = "Bearer " + token
		ctx.Writer.Header().Set("Authorization", bearer)
		ctx.JSON(http.StatusAccepted, gin.H{})
	} else {
		ctx.JSON(http.StatusForbidden, gin.H{})
	}
}

func Logout(ctx *gin.Context) {

}
