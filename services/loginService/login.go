package loginService

import (
	"projeto/FazTudo/dto"
	"projeto/FazTudo/repositorys"
)

type LoginService interface {
	CredentialIsValid(input dto.LoginDTO) bool
	CreateCredential(input dto.LoginDTO) (string, error)
	CreateJWT(input dto.LoginDTO) (string, error)
	ValidateJWT(token string) (string, bool)
	GetIdByLogin(login string) (int, error)
}

type loginService struct {
	repositorys.RepositoryLogin
	JwtService
}

func NewLoginService() LoginService {
	return &loginService{
		RepositoryLogin: repositorys.NewLoginRepository(),
		JwtService:      JwtService{},
	}
}

func (service *loginService) CredentialIsValid(input dto.LoginDTO) bool {

	value, err := service.RepositoryLogin.FindLogin(input)

	if err == nil && value == input {
		return true
	}

	return false
}

func (service *loginService) CreateCredential(input dto.LoginDTO) (string, error) {

	err := service.RepositoryLogin.CreateLogin(input)
	if err != nil {
		return "", err
	}

	return service.JwtService.GenerateToken(input)
}

func (service *loginService) CreateJWT(input dto.LoginDTO) (string, error) {
	return service.JwtService.GenerateToken(input)
}

func (service *loginService) ValidateJWT(token string) (string, bool) {
	return service.JwtService.ValidateToken(token)
}

func (service *loginService) GetIdByLogin(login string) (int, error) {
	return service.RepositoryLogin.GetIdByLogin(login)
}
