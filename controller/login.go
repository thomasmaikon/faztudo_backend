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

	//exist := loginService.NewLoginService().CredentialIsValid(login)

	/* 	if exist {
	ctx.JSON(http.StatusConflict, gin.H{}) */
	/* } else { */
	token, err := loginService.NewLoginService().CreateCredential(login)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"Info": err.Error()})
	} else {
		var bearer string = "Bearer " + token
		ctx.Writer.Header().Set("Authorization", bearer)
		ctx.JSON(http.StatusCreated, gin.H{})
	}
	//}
}

func ValidateCrendential(ctx *gin.Context) {
	var login dto.LoginDTO
	ctx.BindJSON(&login)

	token := loginService.NewLoginService().CredentialIsValid(login)

	if token != "" {
		var bearer string = "Bearer " + token
		ctx.Writer.Header().Set("Authorization", bearer)
		ctx.JSON(http.StatusAccepted, gin.H{"token": token})
	} else {
		ctx.JSON(http.StatusForbidden, gin.H{})
	}
}

func Logout(ctx *gin.Context) {

}
