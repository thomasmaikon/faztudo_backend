package services

import (
	"projeto/FazTudo/dto"
	"projeto/FazTudo/repositorys"
)

type IUserService interface {
	CreateUser(inputUser *dto.UserDTO, idLogin int) (int, error)
	FindUserId(loginId int) (int, error)
}

type userService struct {
	UserRepository repositorys.IUserRepository
}

func GetUserService() IUserService {
	return &userService{
		UserRepository: repositorys.GetUserRepository(),
	}
}

func (service *userService) CreateUser(inputUser *dto.UserDTO, idLogin int) (int, error) {
	user, err := service.UserRepository.CreateUser(inputUser, idLogin)

	return user.Id, err
}

func (service *userService) FindUserId(loginId int) (int, error) {
	user, err := service.UserRepository.FindUser(loginId)

	return user.Id, err
}
