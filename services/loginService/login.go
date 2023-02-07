package loginService

import (
	"database/sql"
	"projeto/FazTudo/dto"
	"projeto/FazTudo/infrastructure/database"
	"projeto/FazTudo/repositorys"
)

type loginService struct {
	db *sql.DB
}

func NewLoginSerice() *loginService {
	return &loginService{
		db: database.GetDBAccess(),
	}
}

func (service *loginService) ValidateCredential(input dto.LoginDTO) bool {
	repository := repositorys.NewLoginRepository(service.db)

	_, err := repository.FindLogin(input)

	if err != nil {
		return false
	}

	return true
}

func (sercice *loginService) CreateCredential(input dto.LoginDTO) (string, error) {

	repository := repositorys.NewLoginRepository(sercice.db)

	err := repository.CreateLogin(input)
	if err != nil {
		return "", err
	}

	return generateToken(input)
}

func (service *loginService) CreateJWT(input dto.LoginDTO) (string, error) {
	return generateToken(input)
}

func (service *loginService) ValidateJWT(token string) (string, bool) {
	return validateToken(token)
}

func (service *loginService) GetIdByLogin(login string) (int, error) {
	return repositorys.NewLoginRepository(service.db).GetIdByLogin(login)
}
