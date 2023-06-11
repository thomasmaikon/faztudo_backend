package IntegrationTest

import (
	"projeto/FazTudo/dto"
	"projeto/FazTudo/repositorys"
	"testing"
)

func TestCriateSimpleService(t *testing.T) {

	// initialize login for tests
	repositoryLogin := repositorys.NewLoginRepository()
	input := dto.LoginDTO{
		Login:    "TestServicePage1@test.com",
		Password: "simplePassword",
		User: dto.UserDTO{
			Name: "TestTask",
			Cpf:  "TestTask",
		},
	}

	userLogin, _ := repositoryLogin.CreateLogin(input)

	/* id, err := repositoryLogin.GetIdByLogin(input.Login)
	if err != nil {
		t.Fatal(err)
	} */
	repository := repositorys.NewServicesPageRepository()

	servicePage := dto.ServicePageInput{
		Name:        "test",
		Image:       "/home/images/peolpleID",
		Value:       89.80,
		Description: "Simple description example",
	}

	err := repository.CreateServicePage(servicePage, userLogin.Id)
	if err != nil {
		t.Fatal(err.Error())
	}
}

func TestAmountAtPages(t *testing.T) {
	// initialize login for tests
	repositoryLogin := repositorys.NewLoginRepository()
	repository := repositorys.NewServicesPageRepository()

	input := dto.LoginDTO{
		Login:    "TestServicePage2@test.com",
		Password: "simplePassword",
		User: dto.UserDTO{
			Name: "TestTask2",
			Cpf:  "TestTask2",
		},
	}
	userLogin, _ := repositoryLogin.CreateLogin(input)
	/* id, err := repositoryLogin.GetIdByLogin(input.Login)
	if err != nil {
		t.Fatal(err)
	} */

	for i := 1; i <= 3; i++ {
		servicePage := dto.ServicePageInput{
			Name:        "test",
			Image:       "/home/images/peolpleID",
			Value:       20.20,
			Description: "Simple description example",
		}

		err := repository.CreateServicePage(servicePage, userLogin.Id)
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
		t.Fatalf("Total pages 'result' value [%x] differente are [%x] expected", expectedAmountPages, result)
	}
}

func TestPaginateService(t *testing.T) {
	// initialize login for tests
	repositoryLogin := repositorys.NewLoginRepository()
	repository := repositorys.NewServicesPageRepository()

	input := dto.LoginDTO{
		Login:    "TestServicePage3@test.com",
		Password: "simplePassword",
		User: dto.UserDTO{
			Name: "TestTask3",
			Cpf:  "TestTask3",
		},
	}
	userLogin, _ := repositoryLogin.CreateLogin(input)

	/* id, err := repositoryLogin.GetIdByLogin(input.Login)
	if err != nil {
		t.Fatal(err)
	} */

	for i := 1; i <= 3; i++ {

		servicePage := dto.ServicePageInput{

			Name:        "test",
			Image:       "/home/images/peolpleID",
			Value:       30.30,
			Description: "Simple description example",
		}

		err := repository.CreateServicePage(servicePage, userLogin.Id)
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

/* func TestGetAmountCommitsInServicePage(t *testing.T) {
	loginService := loginService.NewLoginService()
	servicePageService := servicesPageServices.NewServicePage()
	commitService := commitService.NewCommitService()

	inputs := []dto.LoginDTO{
		{
			Login:    "TestServicePage3@test.com",
			Password: "simplePassword",
		},
		{
			Login:    "TestServicePage4@test.com",
			Password: "simplePassword",
		},
		{
			Login:    "TestServicePage5@test.com",
			Password: "simplePassword",
		},
	}
	for _, input := range inputs {
		loginService.CreateCredential(input)
	}

	servicePageService.CreateService(dto.ServicePageInput{
		Name:        "Simple Example",
		Image:       "paht at image",
		Value:       0.0,
		Description: "any",
	}, inputs[0].Login)

	services := servicePageService.GetAllServicesPageByLogin(inputs[0].Login)

	idUser1, _ := loginService.GetIdByLogin(inputs[1].Login)
	idUser2, _ := loginService.GetIdByLogin(inputs[2].Login)

	commits := []dto.CommitInput{
		{
			IdLogin:       idUser1,
			IdServicePage: services[0].Id,
			Commit:        "any 1",
		},
		{
			IdLogin:       idUser1,
			IdServicePage: services[0].Id,
			Commit:        "any 2",
		},
		{
			IdLogin:       idUser2,
			IdServicePage: services[0].Id,
			Commit:        "any 3",
		},
	}

	for _, commit := range commits {
		commitService.AddCommit(commit)
	}

	expetedAmountAtCommitsInPage := len(commits)
	commitsResulted, err := commitService.GetCommitByServicePage(services[0].Id)

	if err != nil {
		t.Fatal(err)
	}

	if expetedAmountAtCommitsInPage != len(commitsResulted) {
		t.Fatal("Amount at commits in servicePage differnte are expected")
	}

} */
