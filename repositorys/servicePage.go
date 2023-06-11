package repositorys

import (
	"context"
	"fmt"
	"math"
	"projeto/FazTudo/consts"
	"projeto/FazTudo/dto"
	"projeto/FazTudo/infrastructure/database"

	"gorm.io/gorm"
)

type RepositoryServicePage interface {
	GetAllServicesPaginated(paginationIndex, paginationSize int) ([]dto.ServicePageOutput, error)
	GetAmountAtPages(paginationSize int) (int, error)
	CreateServicePage(input dto.ServicePageInput, userId int) error
	GetServicesPage(userId int) ([]dto.ServicePageOutput, error)
	GetServicePageById(id int) (dto.ServicePageOutput, error)
}

type servicesPageRepository struct {
	db *gorm.DB
	RepositoryServicePage
}

func NewServicesPageRepository() RepositoryServicePage {
	return &servicesPageRepository{
		db: database.GetDB(),
	}
}

func (repository *servicesPageRepository) GetAllServicesPaginated(paginationIndex, paginationSize int) ([]dto.ServicePageOutput, error) {

	var outputList []dto.ServicePageOutput

	context, cancelContext := context.WithTimeout(context.Background(), consts.QueryTimeoutMedium)
	defer cancelContext()

	jump := paginationIndex * paginationSize
	query := fmt.Sprintf("SELECT id, name, description, image, value, positive_evaluations, negative_evaluations FROM service OFFSET %x ROWS LIMIT %x", jump, paginationSize)

	err := repository.db.WithContext(context).
		Raw(query).
		Scan(&outputList)
	//QueryContext(context, query)
	if err.Error != nil {
		return nil, err.Error
	}
	return outputList, nil
}

func (repository *servicesPageRepository) GetAmountAtPages(paginationSize int) (int, error) {
	pgSize := float64(paginationSize)

	var amountRows int

	context, cancelContext := context.WithTimeout(context.Background(), consts.QueryTimeoutLong)
	defer cancelContext()

	query := "SELECT COUNT(*) FROM service"

	err := repository.db.WithContext(context).Raw(query).Scan(&amountRows) //.Raw(query).Scan(&amountRows)
	if err.Error != nil {
		return 0, err.Error
	}

	result := int(math.Round(float64(amountRows) / pgSize))
	return result, nil
}

func (repository *servicesPageRepository) CreateServicePage(input dto.ServicePageInput, userId int) error {
	context, cancelContext := context.WithTimeout(context.Background(), consts.QueryTimeoutMedium)
	defer cancelContext()

	query := fmt.Sprintf("INSERT INTO service (fk_user, name, image, value, description) VALUES (%x,'%s', '%s', '%f', '%s')", userId, input.Name, input.Image, input.Value, input.Description)

	err := repository.db.WithContext(context).Exec(query)
	if err.Error != nil {
		return err.Error
	}

	return nil
}

func (repository *servicesPageRepository) GetServicesPage(userId int) ([]dto.ServicePageOutput, error) {

	var outputList []dto.ServicePageOutput

	context, cancelContext := context.WithTimeout(context.Background(), consts.QueryTimeoutMedium)
	defer cancelContext()

	query := fmt.Sprintf("SELECT S.id, S.name, S.description, S.image, S.value, S.positive_evaluations, S.negative_evaluations FROM service AS S WHERE S.fk_user = %x", userId)

	err := repository.db.WithContext(context).Raw(query).Scan(&outputList)
	if err.Error != nil {
		return nil, err.Error
	}

	return outputList, nil
}

func (repository *servicesPageRepository) GetServicePageById(id int) (dto.ServicePageOutput, error) {
	context, cancelContext := context.WithTimeout(context.Background(), consts.QueryTimeoutMedium)
	defer cancelContext()

	var outputList dto.ServicePageOutput

	query := fmt.Sprintf("SELECT id, name, description, image, value, positive_evaluations, negative_evaluations FROM service WHERE id = '%x'", id)

	err := repository.db.WithContext(context).Raw(query).Scan(&outputList)

	return outputList, err.Error
}
