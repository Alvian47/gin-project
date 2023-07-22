package controller

import (
	"diary_api/helper"
	"diary_api/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AddEntry(ctx *gin.Context) {
	var input models.Entry
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": gin.H{
				"errorBindJsonEntry": err.Error(),
			},
		})
	}

	user, err := helper.CurrentUser(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": gin.H{
				"errorValidateUser": err.Error(),
			},
		})
	}

	input.UserID = user.ID
	
	savedEntry, err := input.Save()

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": gin.H{
				"errorSaveEntry": err.Error(),
			},
		})
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"data": savedEntry,
	})
}

func GetAllEntries(ctx *gin.Context){
	user, err := helper.CurrentUser(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": gin.H{
				"errorValidateUser": err.Error(),
			},
		})
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data": user.Entries,
	})
}