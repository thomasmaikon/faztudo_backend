package repositorys

import (
	"context"
	"projeto/FazTudo/consts"
	"projeto/FazTudo/dto"
	"projeto/FazTudo/entitys"
	"projeto/FazTudo/infrastructure/database"

	"gorm.io/gorm"
)

type RepositoryLike interface {
	AddLikeOrUnlike(input dto.LikeInput) error
}

type likeRepository struct {
	db *gorm.DB
}

func NewLikeRepository() RepositoryLike {
	return &likeRepository{database.GetDB()}
}

func (l *likeRepository) AddLikeOrUnlike(input dto.LikeInput) error {
	ctx, cancel := context.WithTimeout(context.Background(), consts.QueryTimeoutMedium)
	defer cancel()

	//query := fmt.Sprintf("INSERT INTO likes (fk_service_page, fk_login, liker) VALUES (%x, %x, %x)", input.ServicePageId, input.LoginId, input.Like)

	err := l.db.WithContext(ctx).Create(&entitys.Like{
		LoginId:       input.LoginId,
		ServicePageId: input.ServicePageId,
		Like:          input.Like,
	})

	//.ExecContext(ctx, query)
	if err.Error != nil {
		return err.Error
	}
	return nil
}
