package loginService

import (
	"fmt"
	"os"
	"projeto/FazTudo/dto"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

func generateToken(input dto.LoginDTO) (string, error) {
	key := []byte(os.Getenv("secretkey"))

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.RegisteredClaims{
		Issuer:    input.Login,
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
	})

	result, err := token.SignedString(key)
	if err != nil {
		return "", err
	}

	return result, nil
}

func validateToken(token string) (string, bool) {

	result, _ := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {

			return nil, fmt.Errorf("Unexpected signing method: %v", t.Header["alg"])
		}
		return []byte(os.Getenv("secretkey")), nil
	})

	if claims, ok := result.Claims.(jwt.MapClaims); ok && result.Valid {
		data := claims["iss"].(string)
		return data, true
	} else {
		return "", false
	}

}

func IsAuthorized(ctx *gin.Context) {
	BEARER_SCHEMA := "Bearer"

	authHeader := ctx.GetHeader("Authorization")
	tokenString := authHeader[len(BEARER_SCHEMA)+1:]

	email, isValid := validateToken(tokenString)

	if !isValid {
		ctx.AbortWithStatus(401)
	}
	ctx.AddParam("email", email)
}
