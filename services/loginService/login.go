package loginService

import (
	"projeto/FazTudo/dto"
	"projeto/FazTudo/repositorys"
	"projeto/FazTudo/services"
)

type LoginService interface {
	CredentialIsValid(input dto.LoginDTO) string
	CreateCredential(input dto.LoginDTO) (string, error)
	CreateJWT(userId int) (string, error)
	ValidateJWT(token string) (string, bool)
	//	GetIdByLogin(userId int) (int, error)
}

type loginService struct {
	UserService services.IUserService
	repositorys.RepositoryLogin
	JwtService
}

func NewLoginService() LoginService {
	return &loginService{
		RepositoryLogin: repositorys.NewLoginRepository(),
		JwtService:      JwtService{},
		UserService:     services.GetUserService(),
	}
}

func (service *loginService) CredentialIsValid(input dto.LoginDTO) string {

	login, err := service.RepositoryLogin.FindLogin(input)

	if err == nil && login != nil {

		userId, _ := service.UserService.FindUserId(login.Id)
		token, _ := service.GenerateToken(userId)

		return token
	}

	return ""
}

func (service *loginService) CreateCredential(input dto.LoginDTO) (string, error) {

	login, err := service.RepositoryLogin.CreateLogin(input)
	if err != nil {
		return "", err
	}

	userId, err := service.UserService.CreateUser(&input.User, login.Id)
	if err != nil {
		return "", err
	}

	return service.JwtService.GenerateToken(userId)
}

func (service *loginService) CreateJWT(userId int) (string, error) {
	return service.JwtService.GenerateToken(userId)
}

func (service *loginService) ValidateJWT(token string) (string, bool) {
	return service.JwtService.ValidateToken(token)
}

/*
func (service *loginService) GetIdByLogin(userId int) (int, error) {
	return service.RepositoryLogin.GetIdByLogin(userId)
}
*/
