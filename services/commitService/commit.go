package commitService

import (
	"database/sql"
	"projeto/FazTudo/dto"
	"projeto/FazTudo/infrastructure/database"
	"projeto/FazTudo/repositorys"
	"projeto/FazTudo/services/loginService"
)

type commitService struct {
	db *sql.DB
}

func NewCommitService() *commitService {
	return &commitService{
		db: database.GetDBAccess(),
	}
}

func (c *commitService) GetCommitByServicePage(servicePageid int) ([]dto.CommitOutput, error) {
	repository := repositorys.NewCommitRepository(c.db)

	output, err := repository.GetCommitByServicePageId(servicePageid)
	if err != nil {
		return nil, err
	}

	return output, nil
}

func (c *commitService) CreateCommit(login string, servicePageId int, commit dto.SimpleCommitInput) error {
	service := loginService.NewLoginSerice()

	id, err := service.GetIdByLogin(login)
	if err != nil {
		return err
	}

	repository := repositorys.NewCommitRepository(c.db)
	return repository.AddCommit(dto.CommitInput{
		IdLogin:       id,
		IdServicePage: servicePageId,
		Commit:        commit.Commit,
	})
}
