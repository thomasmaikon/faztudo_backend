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

type RepositoryCommit interface {
	AddCommit(input dto.CommitInput) error
	GetCommitByServicePageId(servicePageId int) ([]dto.CommitOutput, error)
}

type commitRepository struct {
	db *gorm.DB
}

func NewCommitRepository() RepositoryCommit {
	return &commitRepository{database.GetDB()}
}

func (repository *commitRepository) AddCommit(input dto.CommitInput) error {
	ctx, cancel := context.WithTimeout(context.Background(), consts.QueryTimeoutMedium)
	defer cancel()

	//query := fmt.Sprintf("INSERT INTO commit (fk_login, fk_service_page, commit) VALUES ('%x', '%x', '%s')", input.IdLogin, input.IdServicePage, input.Commit)

	result := repository.db.WithContext(ctx).Create(&entitys.Commit{
		UserId: uint64(input.IdLogin),
		PageId: uint64(input.IdServicePage),
		Commit: input.Commit,
	})

	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (repository *commitRepository) GetCommitByServicePageId(servicePageId int) ([]dto.CommitOutput, error) {
	var outputList []dto.CommitOutput

	ctx, cancel := context.WithTimeout(context.Background(), consts.QueryTimeoutMedium)
	defer cancel()

	query := fmt.Sprintf(`SELECT l.login, c.commit FROM commit as c INNER JOIN service as s ON s.id = c.fk_service 
	INNER JOIN login as l ON s.fk_login = l.id WHERE c.fk_service = %x`, servicePageId)

	result := repository.db.WithContext(ctx).
		Raw(query).
		Scan(&outputList)

	if result.Error != nil {
		return nil, result.Error
	}

	return outputList, nil
}
