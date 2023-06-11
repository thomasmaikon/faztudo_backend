package repositorys

import (
	"projeto/FazTudo/dto"
	"projeto/FazTudo/entitys"
	"projeto/FazTudo/infrastructure/database"

	"gorm.io/gorm"
)

type IUserRepository interface {
	CreateUser(inputUser *dto.UserDTO, idLogin int) (entitys.User, error)
	FindUser(loginId int) (*entitys.User, error)
}

type userRepository struct {
	Db *gorm.DB
}

func GetUserRepository() IUserRepository {
	return &userRepository{
		Db: database.GetDB(),
	}
}

func (repository *userRepository) CreateUser(inputUser *dto.UserDTO, idLogin int) (entitys.User, error) {
	newUser := entitys.User{
		Cpf:         inputUser.Cpf,
		Name:        inputUser.Name,
		LoginUserId: idLogin,
	}

	err := repository.Db.Create(&newUser)

	return newUser, err.Error
}

func (repository *userRepository) FindUser(loginId int) (*entitys.User, error) {
	var user entitys.User
	err := repository.Db.Table("users").Where("login_user_id = ?", loginId).Scan(&user)

	return &user, err.Error
}
