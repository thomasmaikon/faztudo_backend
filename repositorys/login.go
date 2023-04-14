package repositorys

import (
	"context"
	"fmt"
	"projeto/FazTudo/consts"
	"projeto/FazTudo/dto"
	"projeto/FazTudo/entitys"
	"projeto/FazTudo/infrastructure/database"

	"gorm.io/gorm"
)

type RepositoryLogin interface {
	FindLogin(inputDTO dto.LoginDTO) (dto.LoginDTO, error)
	CreateLogin(inputDTO dto.LoginDTO) error
	GetIdByLogin(login string) (int, error)
}

type loginRepository struct {
	db *gorm.DB
}

func NewLoginRepository() RepositoryLogin {
	return &loginRepository{database.GetDB()}
}

func (repository *loginRepository) FindLogin(inputDTO dto.LoginDTO) (dto.LoginDTO, error) {

	var user entitys.LoginUser

	ctx, cancel := context.WithTimeout(context.Background(), consts.QueryTimeoutShort)
	defer cancel()

	query := fmt.Sprintf("SELECT login, password FROM credentials WHERE credentials.login like '%s' and credentials.password like '%s'", inputDTO.Login, inputDTO.Password)

	err := repository.db.WithContext(ctx).Raw(query).Scan(&user)
	if err.Error != nil {
		return dto.LoginDTO{}, err.Error
	}

	return dto.LoginDTO{user.Login, user.Password}, nil
}

func (repository *loginRepository) CreateLogin(inputDTO dto.LoginDTO) error {

	ctx, cancel := context.WithTimeout(context.Background(), consts.QueryTimeoutShort)
	defer cancel()

	query := fmt.Sprintf("INSERT INTO credentials (login, password) VALUES ('%s', '%s')", inputDTO.Login, inputDTO.Password)

	err := repository.db.WithContext(ctx).Exec(query)
	if err.Error != nil {
		return err.Error
	}

	//fmt.Println(rows.RowsAffected())
	return nil
}

func (repository *loginRepository) GetIdByLogin(login string) (int, error) {

	var userId int

	ctx, cancel := context.WithTimeout(context.Background(), consts.QueryTimeoutShort)
	defer cancel()

	query := fmt.Sprintf("SELECT ID FROM credentials WHERE login = '%s'", login)

	err := repository.db.WithContext(ctx).Raw(query).Scan(&userId)
	if err.Error != nil {
		return 0, err.Error
	}

	return userId, nil
}
