package controller

import (
	"diary_api/helper"
	"diary_api/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Register(ctx *gin.Context) {
	var input models.AuthenticationInput

	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"errorBindJSON": err.Error(),
		})
		return
	}

	user := models.User{
		Username: input.Username,
		Password: input.Password,
	}

	savedUser, err := user.Save()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"errorCreatedUser": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"user": savedUser})
}

func Login(ctx *gin.Context) {
	var input models.AuthenticationInput

	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"errorBindJSON": err.Error(),
		})
		return
	}

	user, err := models.FindUserByUsername(input.Username)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"errorFindUser": err.Error(),
		})
		return
	}

	err = user.ValidatePassword(input.Password)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"errorValidPass": err.Error(),
		})
		return
	}

	jwt, err := helper.GenerateJWT(user)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"errorJwt": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"jwt": jwt,
	})
}
