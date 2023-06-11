package unittest

import (
	"projeto/FazTudo/services/loginService"
	"strconv"
	"testing"
)

func TestValidateJWT(t *testing.T) {
	service := loginService.NewJWTService()

	simpleId := 10

	token, err := service.GenerateToken(simpleId)
	if err != nil {
		t.Fatal(err.Error())
	}

	received, ok := service.ValidateToken(token)

	output, _ := strconv.Atoi(received)

	if output != simpleId && !ok {
		t.Fatal("Erro when validate token")
	}
}
