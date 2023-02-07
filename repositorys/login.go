package repositorys

import (
	"context"
	"database/sql"
	"fmt"
	"projeto/FazTudo/consts"
	"projeto/FazTudo/dto"
)

type loginRepository struct {
	db *sql.DB
}

func NewLoginRepository(db *sql.DB) *loginRepository {
	return &loginRepository{db}
}

func (l *loginRepository) FindLogin(inputDTO dto.LoginDTO) (dto.LoginDTO, error) {

	ctx, cancel := context.WithTimeout(context.Background(), consts.QueryTimeoutShort)
	defer cancel()

	query := fmt.Sprintf("SELECT login, password FROM credentials WHERE credentials.login like '%s' and credentials.password like '%s'", inputDTO.Login, inputDTO.Password)

	rows, err := l.db.QueryContext(ctx, query)
	if err != nil {
		return dto.LoginDTO{}, err
	}

	var login string
	var password string

	rows.Next()
	err = rows.Scan(&login, &password)
	if err != nil {
		return dto.LoginDTO{}, err
	}

	return dto.LoginDTO{login, password}, nil
}

func (l *loginRepository) CreateLogin(inputDTO dto.LoginDTO) error {

	ctx, cancel := context.WithTimeout(context.Background(), consts.QueryTimeoutShort)
	defer cancel()

	query := fmt.Sprintf("INSERT INTO credentials (login, password) VALUES ('%s', '%s')", inputDTO.Login, inputDTO.Password)

	rows, err := l.db.ExecContext(ctx, query)
	if err != nil {
		return err
	}

	fmt.Println(rows.RowsAffected())
	return nil
}

func (l *loginRepository) GetIdByLogin(login string) (int, error) {
	ctx, cancel := context.WithTimeout(context.Background(), consts.QueryTimeoutShort)
	defer cancel()

	query := fmt.Sprintf("SELECT ID FROM credentials WHERE login = '%s'", login)

	rows, err := l.db.QueryContext(ctx, query)
	if err != nil {
		return 0, err
	}

	var id int
	rows.Next()
	err = rows.Scan(&id)
	if err != nil {
		return 0, err
	}

	return id, nil
}
