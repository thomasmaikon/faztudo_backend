package IntegrationTest

import (
	"projeto/FazTudo/dto"
	"projeto/FazTudo/repositorys"
	"testing"
)

func TestCriateSimpleService(t *testing.T) {

	// initialize login for tests
	repositoryLogin := repositorys.NewLoginRepository(db)
	input := dto.LoginDTO{
		Login:    "TestServicePage1@test.com",
		Password: "simplePassword",
	}

	repositoryLogin.CreateLogin(input)
	id, err := repositoryLogin.GetIdByLogin(input.Login)
	if err != nil {
		t.Fatal(err)
	}
	repository := repositorys.NewServicesPage(db)

	servicePage := dto.ServicePageInput{
		Name:        "test",
		Image:       "/home/images/peolpleID",
		Value:       89.80,
		Description: "Simple description example",
	}

	err = repository.CreateServicePage(servicePage, id)
	if err != nil {
		t.Fatal(err.Error())
	}
}

func TestAmountAtPages(t *testing.T) {
	// initialize login for tests
	repositoryLogin := repositorys.NewLoginRepository(db)
	input := dto.LoginDTO{
		Login:    "TestServicePage2@test.com",
		Password: "simplePassword",
	}
	repositoryLogin.CreateLogin(input)
	id, err := repositoryLogin.GetIdByLogin(input.Login)
	if err != nil {
		t.Fatal(err)
	}

	repository := repositorys.NewServicesPage(db)

	for i := 1; i <= 3; i++ {
		servicePage := dto.ServicePageInput{
			Name:        "test",
			Image:       "/home/images/peolpleID",
			Value:       20.20,
			Description: "Simple description example",
		}

		err := repository.CreateServicePage(servicePage, id)
		if err != nil {
			t.Fatal(err.Error())
		}
	}
	var paginationSizeTest int = 3
	var expectedAmountPages int = 1
	result, err := repository.GetAmountAtPages(paginationSizeTest)
	if err != nil {
		t.Fatal(err)
	}

	if expectedAmountPages != result {
		t.Fatalf("Total pages 'result' value differente are expected, [%x] != [%x]", expectedAmountPages, result)
	}
}

func TestPaginateService(t *testing.T) {
	// initialize login for tests
	repositoryLogin := repositorys.NewLoginRepository(db)
	input := dto.LoginDTO{
		Login:    "TestServicePage3@test.com",
		Password: "simplePassword",
	}
	repositoryLogin.CreateLogin(input)
	id, err := repositoryLogin.GetIdByLogin(input.Login)
	if err != nil {
		t.Fatal(err)
	}
	repository := repositorys.NewServicesPage(db)

	for i := 1; i <= 3; i++ {

		servicePage := dto.ServicePageInput{

			Name:        "test",
			Image:       "/home/images/peolpleID",
			Value:       30.30,
			Description: "Simple description example",
		}

		err := repository.CreateServicePage(servicePage, id)
		if err != nil {
			t.Fatal(err.Error())
		}
	}

	paginationSize := 1
	for i := 0; i < 3; i++ {
		result, err := repository.GetAllServicesPaginated(i, paginationSize)
		if err != nil {
			t.Fatal(err)
		}

		// 3 e referente a quantidade que foi solicitada na paginacao, o 2 e devido a quantidade existente <11> entao vai haver um momento que so retorna 2, que e a quantidade restande dos valores do banco
		if len(result) == paginationSize {
			continue
		}
		if len(result) != 1 {
			t.Fatal("Amount at servicePage differnte are expected")
		}
	}
}
