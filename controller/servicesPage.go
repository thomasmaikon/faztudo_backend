package controller

import (
	"fmt"
	"net/http"
	"projeto/FazTudo/dto"
	servicesPageServices "projeto/FazTudo/services/servicesPageService"
	"strconv"

	"github.com/gin-gonic/gin"
)

func CreateServicePage(ctx *gin.Context) {
	var input dto.ServicePageInput
	ctx.BindJSON(&input)

	service := servicesPageServices.NewServicePage()

	email, _ := ctx.Params.Get("email")

	err := service.CreateService(input, fmt.Sprintf("%s", email))

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{})
	} else {
		ctx.JSON(http.StatusCreated, gin.H{})
	}

}

func GetAllServicePage(ctx *gin.Context) {
	service := servicesPageServices.NewServicePage()
	index := ctx.Param("index")
	indexPage, _ := strconv.Atoi(index)

	listServicesPage, err := service.GetAllServicesPaginateds(indexPage)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{})
	} else {
		ctx.JSON(http.StatusOK, listServicesPage)
	}

}

func GetMyServicesPage(ctx *gin.Context) {
	service := servicesPageServices.NewServicePage()
	email, ok := ctx.Params.Get("email")
	if !ok {
		ctx.JSON(http.StatusBadGateway, gin.H{})
		return
	}

	result := service.GetAllServicesPageByLogin(email)

	if result != nil {
		ctx.JSON(http.StatusOK, gin.H{"services": result})
	} else {
		ctx.JSON(http.StatusBadRequest, gin.H{})
	}
}
