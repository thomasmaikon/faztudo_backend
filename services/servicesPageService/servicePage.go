package servicesPageServices

import (
	"database/sql"
	"projeto/FazTudo/dto"
	"projeto/FazTudo/infrastructure/database"
	"projeto/FazTudo/repositorys"
	"projeto/FazTudo/services/loginService"
)

type servicesPage struct {
	db *sql.DB
}

const paginationSize = 5

func NewServicePage() *servicesPage {
	return &servicesPage{
		db: database.GetDBAccess(),
	}
}

func (service *servicesPage) GetAllServicesPaginateds(index int) ([]dto.ServicePageOutput, error) {
	repository := repositorys.NewServicesPage(service.db)
	return repository.GetAllServicesPaginated(index, paginationSize)
}

func (service *servicesPage) GetAmountPages() (int, error) {
	repository := repositorys.NewServicesPage(service.db)

	return repository.GetAmountAtPages(paginationSize)
}

func (service *servicesPage) CreateService(input dto.ServicePageInput, login string) error {
	serviceLogin := loginService.NewLoginSerice()
	repository := repositorys.NewServicesPage(service.db)

	id, err := serviceLogin.GetIdByLogin(login)
	if err != nil {
		return err
	}

	err = repository.CreateServicePage(input, id)
	if err != nil {
		return err
	}

	return nil
}

func (service *servicesPage) GetAllServicesPageByLogin(login string) []dto.ServicePageOutput {
	repository := repositorys.NewServicesPage(service.db)
	output, err := repository.GetServicesPageByLogin(login)
	if err != nil {
		return nil
	}

	return output
}
