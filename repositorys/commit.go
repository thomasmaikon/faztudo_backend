package repositorys

import (
	"context"
	"database/sql"
	"fmt"
	"projeto/FazTudo/consts"
	"projeto/FazTudo/dto"
)

type commitRepository struct {
	db *sql.DB
}

func NewCommitRepository(db *sql.DB) *commitRepository {
	return &commitRepository{db}
}

func (c *commitRepository) AddCommit(input dto.CommitInput) error {
	ctx, cancel := context.WithTimeout(context.Background(), consts.QueryTimeoutMedium)
	defer cancel()

	query := fmt.Sprintf("INSERT INTO commit (fk_login, fk_service_page, commit) VALUES ('%x', '%x', '%s')", input.IdLogin, input.IdServicePage, input.Commit)

	_, err := c.db.QueryContext(ctx, query)
	if err != nil {
		return err
	}

	return nil
}

func (c *commitRepository) GetCommitByServicePageId(servicePageId int) ([]dto.CommitOutput, error) {
	ctx, cancel := context.WithTimeout(context.Background(), consts.QueryTimeoutMedium)
	defer cancel()

	query := fmt.Sprintf(`SELECT l.login, c.commit FROM commit as c INNER JOIN service as s ON s.id = c.fk_service 
	INNER JOIN login as l ON s.fk_login = l.id WHERE c.fk_service = %x`, servicePageId)

	rows, err := c.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}

	var outputList []dto.CommitOutput

	for rows.Next() {
		output := dto.CommitOutput{}
		rows.Scan(output.Login, output.Commit)
		outputList = append(outputList, output)
	}

	return outputList, nil
}
