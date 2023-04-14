package unittest

import (
	"projeto/FazTudo/dto"
	"projeto/FazTudo/services/loginService"
	"testing"
)

func TestValidateJWT(t *testing.T) {
	service := loginService.NewJWTService()

	input := dto.LoginDTO{
		Login:    "simpleExampleTest@hotmail.com",
		Password: "simplePassword123",
	}

	token, err := service.GenerateToken(input)
	if err != nil {
		t.Fatal(err.Error())
	}

	expected, ok := service.ValidateToken(token)

	if expected != input.Login && !ok {
		t.Fatal("Erro when validate token")
	}
}
