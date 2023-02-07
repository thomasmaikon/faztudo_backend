package repositorys

import (
	"context"
	"database/sql"
	"fmt"
	"math"
	"projeto/FazTudo/consts"
	"projeto/FazTudo/dto"
)

type servicesPage struct {
	db *sql.DB
}

func NewServicesPage(db *sql.DB) *servicesPage {
	return &servicesPage{db}
}

func (s *servicesPage) GetAllServicesPaginated(paginationIndex, paginationSize int) ([]dto.ServicePageOutput, error) {
	context, cancelContext := context.WithTimeout(context.Background(), consts.QueryTimeoutMedium)
	defer cancelContext()

	jump := paginationIndex * paginationSize
	query := fmt.Sprintf("SELECT id, name, description, image, value, positive_evaluations, negative_evaluations FROM service OFFSET %x ROWS LIMIT %x", jump, paginationSize)

	rows, err := s.db.QueryContext(context, query)

	var outputList []dto.ServicePageOutput

	for rows.Next() {
		output := dto.ServicePageOutput{}
		err = rows.Scan(&output.Id, &output.Name, &output.Description, &output.Image, &output.Value, &output.PositiveEvaluations, &output.NegativeEvaluations)
		if err != nil {
			return nil, err
		}
		outputList = append(outputList, output)
	}

	return outputList, err
}

func (s *servicesPage) GetAmountAtPages(paginationSize int) (int, error) {
	pgSize := float64(paginationSize)
	context, cancelContext := context.WithTimeout(context.Background(), consts.QueryTimeoutLong)
	defer cancelContext()

	query := "SELECT count(*) FROM service"

	rows, err := s.db.QueryContext(context, query)
	if err != nil {
		return 0, err
	}

	var amountRows float64

	for rows.Next() {
		err = rows.Scan(&amountRows)
		if err != nil {
			return 0, err
		}
	}
	result := int(math.Round(amountRows / pgSize))
	return result, nil
}

func (s *servicesPage) CreateServicePage(input dto.ServicePageInput, fkLogin int) error {
	context, cancelContext := context.WithTimeout(context.Background(), consts.QueryTimeoutMedium)
	defer cancelContext()

	query := fmt.Sprintf("INSERT INTO service (fk_login, name, image, value, description) VALUES (%x,'%s', '%s', '%f', '%s')", fkLogin, input.Name, input.Image, input.Value, input.Description)

	_, err := s.db.QueryContext(context, query)
	if err != nil {
		return err
	}

	return nil
}

func (s *servicesPage) GetServicesPageByLogin(login string) ([]dto.ServicePageOutput, error) {
	context, cancelContext := context.WithTimeout(context.Background(), consts.QueryTimeoutMedium)
	defer cancelContext()

	query := fmt.Sprintf("SELECT S.id, S.name, S.description, S.image, S.value, S.positive_evaluations, S.negative_evaluations FROM service AS S INNER JOIN credentials AS C on c.ID = S.fk_login WHERE C.LOGIN = '%s'", login)

	rows, err := s.db.QueryContext(context, query)
	if err != nil {
		return nil, err
	}

	var outputList []dto.ServicePageOutput

	for rows.Next() {
		output := dto.ServicePageOutput{}
		err = rows.Scan(&output.Id, &output.Name, &output.Description, &output.Image, &output.Value, &output.PositiveEvaluations, &output.NegativeEvaluations)
		if err != nil {
			return nil, err
		}
		outputList = append(outputList, output)
	}

	return outputList, err
}

func (s *servicesPage) GetServicePageById(id int) (dto.ServicePageOutput, error) {
	context, cancelContext := context.WithTimeout(context.Background(), consts.QueryTimeoutMedium)
	defer cancelContext()

	var outputList dto.ServicePageOutput

	query := fmt.Sprintf("SELECT id, name, description, image, value, positive_evaluations, negative_evaluations FROM service WHERE id = '%x'", id)

	rows, err := s.db.QueryContext(context, query)
	if err != nil {
		return outputList, err
	}

	for rows.Next() {

		err = rows.Scan(&outputList.Id, &outputList.Name, &outputList.Description, &outputList.Image, &outputList.Value, &outputList.PositiveEvaluations, &outputList.NegativeEvaluations)
		if err != nil {
			return outputList, err
		}
	}

	return outputList, err
}
