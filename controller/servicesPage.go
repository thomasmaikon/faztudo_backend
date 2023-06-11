package controller

import (
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

	id := ctx.GetString("userId")

	userId, err := strconv.Atoi(id)

	err = service.CreateService(input, userId)

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

	id := ctx.GetString("userId")
	userId, err := strconv.Atoi(id)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{})
		return
	}

	service := servicesPageServices.NewServicePage()

	result := service.GetAllServicesPage(userId)
	if result != nil {
		ctx.JSON(http.StatusOK, gin.H{"services": result})
	} else {
		ctx.JSON(http.StatusBadRequest, gin.H{})
	}
}
