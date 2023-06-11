package commitService

import (
	"projeto/FazTudo/dto"
	"projeto/FazTudo/repositorys"
	"projeto/FazTudo/services/loginService"
)

type commitService struct {
	repositorys.RepositoryCommit
	loginService.LoginService
}

func NewCommitService() *commitService {
	return &commitService{
		RepositoryCommit: repositorys.NewCommitRepository(),
		LoginService:     loginService.NewLoginService(),
	}
}

func (service *commitService) GetCommitByServicePage(servicePageid int) ([]dto.CommitOutput, error) {
	output, err := service.RepositoryCommit.GetCommitByServicePageId(servicePageid)
	if err != nil {
		return nil, err
	}

	return output, nil
}

func (service *commitService) CreateCommit(userId int, servicePageId int, commit dto.SimpleCommitInput) error {
	//	service := loginService.NewLoginSerice()

	/* id, err := service.LoginService.GetIdByLogin(login)
	if err != nil {
		return err
	} */

	//repository := repositorys.NewCommitRepository(c.db)
	return service.RepositoryCommit.AddCommit(dto.CommitInput{
		UserId:        userId,
		ServicePageId: servicePageId,
		Commit:        commit.Commit,
	})
}
