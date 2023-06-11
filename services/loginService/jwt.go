package loginService

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

type ServiceJWT interface {
	GenerateToken(userId int) (string, error)
	ValidateToken(token string) (string, bool)
}

type JwtService struct {
	ServiceJWT
}

func NewJWTService() ServiceJWT {
	return &JwtService{}
}

func (service *JwtService) GenerateToken(userId int) (string, error) {
	key := []byte(os.Getenv("secretkey"))

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.RegisteredClaims{
		Subject:   strconv.Itoa(userId),
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
	})

	result, err := token.SignedString(key)
	if err != nil {
		return "", err
	}

	return result, nil
}

func (service *JwtService) ValidateToken(token string) (string, bool) {

	result, _ := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {

			return nil, fmt.Errorf("Unexpected signing method: %v", t.Header["alg"])
		}
		return []byte(os.Getenv("secretkey")), nil
	})

	if claims, ok := result.Claims.(jwt.MapClaims); ok && result.Valid {
		userId := claims["sub"].(string)
		return userId, true
	} else {
		return "", false
	}

}

func IsAuthorized(ctx *gin.Context) {
	BEARER_SCHEMA := "Bearer"

	authHeader := ctx.GetHeader("Authorization")
	tokenString := authHeader[len(BEARER_SCHEMA)+1:]

	service := NewJWTService()

	id, isValid := service.ValidateToken(tokenString)

	if !isValid {
		ctx.AbortWithStatus(401)
		return
	}
	ctx.Set("userId", id)
	ctx.Next()
}
