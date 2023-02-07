package repositorys

import (
	"context"
	"database/sql"
	"fmt"
	"projeto/FazTudo/consts"
	"projeto/FazTudo/dto"
)

type likeRepository struct {
	db *sql.DB
}

func NewLikeRepository(db *sql.DB) *likeRepository {
	return &likeRepository{db}
}

func (l *likeRepository) AddLikeOrUnlike(input dto.LikeInput) error {
	ctx, cancel := context.WithTimeout(context.Background(), consts.QueryTimeoutMedium)
	defer cancel()

	query := fmt.Sprintf("INSERT INTO likes (fk_service_page, fk_login, liker) VALUES (%x, %x, %x)", input.ServicePageId, input.LoginId, input.Like)

	_, err := l.db.ExecContext(ctx, query)
	if err != nil {
		return err
	}
	return nil
}
