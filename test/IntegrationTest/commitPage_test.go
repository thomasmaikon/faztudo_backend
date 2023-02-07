package IntegrationTest

import (
	"projeto/FazTudo/dto"
	"projeto/FazTudo/repositorys"
	"testing"
)

func TestAddingCommit(t *testing.T) {
	loginRepository := repositorys.NewLoginRepository(db)
	serviceRepository := repositorys.NewServicesPage(db)
	commitRepository := repositorys.NewCommitRepository(db)

	login := dto.LoginDTO{
		Login:    "exampleCommitLogin@hotmail.com",
		Password: "123qwe",
	}

	err := loginRepository.CreateLogin(login)
	if err != nil {
		t.Fatal(err.Error())
	}

	loginId, err := loginRepository.GetIdByLogin(login.Login)
	if err != nil {
		t.Fatal(err.Error())
	}

	servicePage := dto.ServicePageInput{
		Name:        "Input Commit",
		Image:       "sdlfkmskldf",
		Value:       0.0,
		Description: "Simple Example for input commit",
	}

	err = serviceRepository.CreateServicePage(servicePage, loginId)
	if err != nil {
		t.Fatal(err.Error())
	}

	listServicePage, err := serviceRepository.GetServicesPageByLogin(login.Login)
	if err != nil {
		t.Fatal(err.Error())
	}

	commit := dto.CommitInput{
		IdLogin:       loginId,
		IdServicePage: listServicePage[0].Id,
		Commit:        "example commit that i use in my test",
	}

	result := commitRepository.AddCommit(commit)

	if result != nil {
		t.Fatal(result.Error())
	}
}

func TestAddingCommitThatLoginAndServicePageDoesnotExist(t *testing.T) {

	commitRepository := repositorys.NewCommitRepository(db)

	commit := dto.CommitInput{
		IdLogin:       0,
		IdServicePage: 0,
		Commit:        "example commit that i use in my test",
	}

	result := commitRepository.AddCommit(commit)

	if result == nil {
		t.Fatal("Expected Error, because id at login and servicePage doesnot exist")
	}

}
