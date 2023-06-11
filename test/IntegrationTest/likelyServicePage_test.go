package IntegrationTest

import (
	"projeto/FazTudo/consts"
	"projeto/FazTudo/dto"
	"projeto/FazTudo/repositorys"
	"testing"
)

func TestLikeServicePage(t *testing.T) {
	loginRepository := repositorys.NewLoginRepository()
	servicePageRepository := repositorys.NewServicesPageRepository()
	likeRepository := repositorys.NewLikeRepository()

	user := dto.LoginDTO{
		Login:    "SimpleExampleLikeLogin@hotmail.com",
		Password: "123qwe",
	}

	servicePage := dto.ServicePageInput{
		Name:        "SimpleServiceForlike",
		Image:       "any",
		Value:       10.01,
		Description: "any",
	}

	userLogin, err := loginRepository.CreateLogin(user)
	if err != nil {
		t.Fatal(err.Error())
	}

	/* loginId, err := loginRepository.GetIdByLogin(user.Login)
	if err != nil {
		t.Fatal(err.Error())
	} */

	err = servicePageRepository.CreateServicePage(servicePage, userLogin.Id)
	if err != nil {
		t.Fatal(err.Error())
	}

	listServicePage, err := servicePageRepository.GetAllServicesPaginated(0, 3)
	if err != nil {
		t.Fatal(err.Error())
	}

	// -----------------------
	like := dto.LikeInput{
		LoginId:       userLogin.Id,
		ServicePageId: listServicePage[0].Id,
		Like:          consts.Likely,
	}
	// -----------------------

	err = likeRepository.AddLikeOrUnlike(like)
	if err != nil {
		t.Fatal(err.Error())
	}

}

func TestAddLikeThatAlredyExist(t *testing.T) {
	loginRepository := repositorys.NewLoginRepository()
	servicePageRepository := repositorys.NewServicesPageRepository()
	likeRepository := repositorys.NewLikeRepository()

	user := dto.LoginDTO{
		Login:    "SimpleExampleLikeLogin2@hotmail.com",
		Password: "123qwe",
	}

	servicePage := dto.ServicePageInput{
		Name:        "SimpleServiceForlike",
		Image:       "any",
		Value:       10.01,
		Description: "any",
	}

	userLogin, err := loginRepository.CreateLogin(user)
	if err != nil {
		t.Fatal(err.Error())
	}

	/* loginId, err := loginRepository.GetIdByLogin(user.Login)
	if err != nil {
		t.Fatal(err.Error())
	}
 */
	err = servicePageRepository.CreateServicePage(servicePage, userLogin.Id)
	if err != nil {
		t.Fatal(err.Error())
	}

	listServicePage, err := servicePageRepository.GetAllServicesPaginated(0, 3)
	if err != nil {
		t.Fatal(err.Error())
	}

	// -----------------------
	like := dto.LikeInput{
		LoginId:       userLogin.Id,
		ServicePageId: listServicePage[0].Id,
		Like:          consts.Likely,
	}
	// -----------------------

	err = likeRepository.AddLikeOrUnlike(like)
	if err != nil {
		t.Fatal(err.Error())
	}

	err = likeRepository.AddLikeOrUnlike(like)
	if err != nil {
		t.Fatal("Expected error because primary key are duplicated ")
	}
}
