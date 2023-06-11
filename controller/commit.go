package controller

import (
	"net/http"
	"projeto/FazTudo/dto"
	"projeto/FazTudo/services/commitService"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetCommitsFromServicePage(ctx *gin.Context) {
	servicePageId := ctx.Param("id")
	pageId, _ := strconv.Atoi(servicePageId)

	service := commitService.NewCommitService()

	output, err := service.GetCommitByServicePage(pageId)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{})
	} else {
		ctx.JSON(http.StatusOK, gin.H{"commits": output})
	}
}

func CreateCommitInservicePage(ctx *gin.Context) {
	servicePageId := ctx.Param("id")
	pageId, err := strconv.Atoi(servicePageId)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{})
		return
	}

	id := ctx.GetString("userId")
	userId, err := strconv.Atoi(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"info": err.Error()})
		return
	}

	service := commitService.NewCommitService()

	var commit dto.SimpleCommitInput
	ctx.BindJSON(&commit)

	resutl := service.CreateCommit(userId, pageId, commit)

	if resutl != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{})
	} else {
		ctx.JSON(http.StatusCreated, gin.H{})
	}
}
