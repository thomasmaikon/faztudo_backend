package servicesPageServices

import (
	"projeto/FazTudo/dto"
	"projeto/FazTudo/repositorys"
	"projeto/FazTudo/services/loginService"
)

type servicesPage struct {
	loginService.LoginService
	repositorys.RepositoryServicePage
}

const paginationSize = 5

func NewServicePage() *servicesPage {
	return &servicesPage{
		RepositoryServicePage: repositorys.NewServicesPageRepository(),
		LoginService:          loginService.NewLoginService(),
	}
}

func (service *servicesPage) GetAllServicesPaginateds(index int) ([]dto.ServicePageOutput, error) {

	return service.RepositoryServicePage.GetAllServicesPaginated(index, paginationSize)
}

func (service *servicesPage) GetAmountPages() (int, error) {
	return service.RepositoryServicePage.GetAmountAtPages(paginationSize)
}

func (service *servicesPage) CreateService(input dto.ServicePageInput, login string) error {
	id, err := service.LoginService.GetIdByLogin(login)
	if err != nil {
		return err
	}

	err = service.RepositoryServicePage.CreateServicePage(input, id)
	if err != nil {
		return err
	}

	return nil
}

func (service *servicesPage) GetAllServicesPageByLogin(login string) []dto.ServicePageOutput {
	output, err := service.RepositoryServicePage.GetServicesPageByLogin(login)
	if err != nil {
		return nil
	}

	return output
}
