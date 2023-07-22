package helper

import (
	"diary_api/models"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

var privateKey = []byte(os.Getenv("JWT_PRIVATE_KEY"))

func GenerateJWT(user models.User) (string, error) {
	// TTL = Token Time to Live
	tokenTTL, _ := strconv.Atoi(os.Getenv("TOKEN_TTL"))
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":  user.ID,
		"iat": time.Now().Unix(),
		"exp": time.Now().Add(time.Second * time.Duration(tokenTTL)).Unix(),
	})

	return token.SignedString(privateKey)
}

func ValidateJWT(ctx *gin.Context) error {
	token, err := getToken(ctx)
	if err != nil {
		return err
	}

	_, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		return nil
	}
	return errors.New("invalid token provided")
}

func CurrentUser(ctx *gin.Context) (models.User, error) {
	err := ValidateJWT(ctx)
	if err != nil {
		return models.User{}, err
	}

	token, _ := getToken(ctx)
	claims, _ := token.Claims.(jwt.MapClaims)
	userId := uint(claims["id"].(float64))

	user, err := models.FindUserById(userId)
	if err != nil {
		return models.User{}, err
	}

	return user, nil
}

func getToken(ctx *gin.Context) (*jwt.Token, error) {
	tokenString := getTokenFromRequest(ctx)
	token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
		}
		return privateKey, nil
	})

	return token, err
}

func getTokenFromRequest(ctx *gin.Context) string {
	bearerToken := ctx.Request.Header.Get("Authorization")
	splitToken := strings.Split(bearerToken, " ")
	if len(splitToken) == 2 {
		return splitToken[1]
	}

	return ""
}
