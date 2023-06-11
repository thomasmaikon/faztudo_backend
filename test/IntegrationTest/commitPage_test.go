package IntegrationTest

import (
	"projeto/FazTudo/dto"
	"projeto/FazTudo/repositorys"
	"testing"
)

func TestAddingCommit(t *testing.T) {
	loginRepository := repositorys.NewLoginRepository()
	serviceRepository := repositorys.NewServicesPageRepository()
	commitRepository := repositorys.NewCommitRepository()

	login := dto.LoginDTO{
		Login:    "exampleCommitLogin@hotmail.com",
		Password: "123qwe",
	}

	userLogin, err := loginRepository.CreateLogin(login)
	if err != nil {
		t.Fatal(err.Error())
	}

	/* loginId, err := loginRepository.GetIdByLogin(login.Login)
	if err != nil {
		t.Fatal(err.Error())
	} */

	servicePage := dto.ServicePageInput{
		Name:        "Input Commit",
		Image:       "sdlfkmskldf",
		Value:       0.0,
		Description: "Simple Example for input commit",
	}

	err = serviceRepository.CreateServicePage(servicePage, userLogin.Id)
	if err != nil {
		t.Fatal(err.Error())
	}

	listServicePage, err := serviceRepository.GetServicesPage(userLogin.Id)
	if err != nil {
		t.Fatal(err.Error())
	}

	commit := dto.CommitInput{
		UserId:        userLogin.Id,
		ServicePageId: listServicePage[0].Id,
		Commit:        "example commit that i use in my test",
	}

	result := commitRepository.AddCommit(commit)

	if result != nil {
		t.Fatal(result.Error())
	}
}

func TestAddingCommitThatLoginAndServicePageDoesnotExist(t *testing.T) {

	commitRepository := repositorys.NewCommitRepository()

	commit := dto.CommitInput{
		UserId:        1,
		ServicePageId: 1,
		Commit:        "example commit that i use in my test",
	}

	result := commitRepository.AddCommit(commit)

	if result != nil {
		t.Fatal("Expected Error, because id at login and servicePage doesnot exist")
	}

}
