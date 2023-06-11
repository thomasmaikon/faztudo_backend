package SystemTest

import (
	"log"
	"net/http"
	"os"
	"strings"

	"projeto/FazTudo/dto"
	"projeto/FazTudo/infrastructure/App"
	"projeto/FazTudo/infrastructure/database"
	"projeto/FazTudo/test"
	"testing"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"github.com/steinfletcher/apitest"
)

var app *gin.Engine

func TestMain(m *testing.M) {
	dockerInfo := test.StartPostgresDockerFormTest()
	strs := strings.Split(dockerInfo.HostAndPort, ":")
	port := strs[1]
	ormDB := database.GetDBWithParams(port, dockerInfo.User, dockerInfo.Password, dockerInfo.Dbname)
	// initialize app
	app = App.NewApp().InitializeRoutes().RunMigrations(ormDB).Router

	//Run tests
	code := m.Run()

	if err := dockerInfo.Pool.Purge(dockerInfo.Resource); err != nil {
		log.Fatalf("Could not purge resource: %s", err)
	}

	os.Exit(code)
}

func TestSimpleRequestWithGetMethod(t *testing.T) {
	apitest.New().
		Handler(app).
		Get("/").
		Expect(t).
		Status(http.StatusOK).
		End()
}

func TestCreateLogin(t *testing.T) {
	inputDTO := dto.LoginDTO{
		Login:    "example@hotmail.com",
		Password: "123qwe",
		User: dto.UserDTO{
			Name: "userTest",
			Cpf:  "userCPFTest",
		},
	}

	apitest.New().
		Handler(app).
		Post("/login/create").
		JSON(inputDTO).
		Expect(t).
		Status(http.StatusCreated).
		Body("").
		End()
}

func TestCreateLoginAreExist(t *testing.T) {
	inputDTO := dto.LoginDTO{
		Login:    "example2@hotmail.com",
		Password: "123qwe",
		User: dto.UserDTO{
			Name: "userTest2",
			Cpf:  "userCPFTest2",
		},
	}

	apitest.New().
		Handler(app).
		Post("/login/create").
		JSON(inputDTO).
		Expect(t).
		Status(http.StatusCreated).
		Body("").
		End()

	apitest.New().
		Handler(app).
		Post("/login/create").
		JSON(inputDTO).
		Expect(t).
		Status(http.StatusInternalServerError).
		Body("").
		End()
}

func TestSimpleValidateLogin(t *testing.T) {

	inputDTO := dto.LoginDTO{
		Login:    "example3@hotmail.com",
		Password: "123qwe",
		User: dto.UserDTO{
			Name: "userTest3",
			Cpf:  "userCPFTest3",
		},
	}

	apitest.New().
		Handler(app).
		Post("/login/create").
		JSON(inputDTO).
		Expect(t).
		Status(http.StatusCreated).
		Body("").
		End()

	/*expected := "sdfsdfdsf"
	json, _ := json.Marshal(expected)
	expectedToekn := bytes.NewBuffer(json).String()*/

	apitest.New().
		Handler(app).
		Post("/login/credential").
		JSON(inputDTO).
		Expect(t).
		Status(http.StatusAccepted).
		HeaderPresent("Authorization").
		//Header("Authorization", "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJleGFtcGxlNUBob3RtYWlsLmNvbSIsImV4cCI6MTY3MDk4MTg3N30.9p1Q0p2uM7WTRToTizR2GcF_9JxVJdELxZvJDSWnPJw").
		End()
}

func TestSimpleValidateLoginAreNotExist(t *testing.T) {

	inputDTO := dto.LoginDTO{
		Login:    "example4@hotmail.com",
		Password: "123qwe",
		User: dto.UserDTO{
			Name: "userTest4",
			Cpf:  "userCPFTest4",
		},
	}

	apitest.New().
		Handler(app).
		Post("/login/credential").
		JSON(inputDTO).
		Expect(t).
		Status(http.StatusForbidden).
		Body("").
		End()
}

func TestSimpleCreateServicePageWithLoginUser(t *testing.T) {

	inputMock := dto.LoginDTO{
		Login:    "example5@hotmail.com",
		Password: "123qwe",
		User: dto.UserDTO{
			Name: "userTest5",
			Cpf:  "userCPFTest5",
		},
	}

	inputServicePageMock := dto.ServicePageInput{
		Name:        "SimpleService",
		Image:       "SimplePathFromImage",
		Value:       70.70,
		Description: "MockServiceCreated",
	}

	result := apitest.New().
		Handler(app).
		Post("/login/create").
		JSON(inputMock).
		Expect(t).
		Status(http.StatusCreated).
		HeaderPresent("Authorization").
		Body("").
		End()

	/* result := apitest.New().
	Handler(app).
	Post("/login/credential").
	JSON(inputMock).
	Expect(t).
	Status(http.StatusAccepted).
	HeaderPresent("Authorization").
	Body("").
	End() */

	jwt := result.Response.Header.Get("Authorization")

	apitest.New().
		Handler(app).
		Post("/servicepage/create").
		JSON(inputServicePageMock).
		Header("Authorization", jwt).
		Expect(t).
		Status(http.StatusCreated).
		Body("").
		End()
}

func TestSimpleListALLPaginatedServicePage(t *testing.T) {
	pageIndex := "0"
	apitest.New().
		Handler(app).
		Get("/servicepage/all/" + pageIndex).
		Expect(t).
		Status(http.StatusOK).
		End()

}

func TestSimpleListPersonPaginatedServicePage(t *testing.T) {
	inputDTO := dto.LoginDTO{
		Login:    "example6@hotmail.com",
		Password: "123qwe",
		User: dto.UserDTO{
			Name: "userTest6",
			Cpf:  "userCPFTest6",
		},
	}

	inputServicePageMock := dto.ServicePageInput{
		Name:        "SimpleService",
		Image:       "SimplePathFromImage",
		Value:       80.80,
		Description: "MockServiceCreated",
	}

	result := apitest.New().
		Handler(app).
		Post("/login/create").
		JSON(inputDTO).
		Expect(t).
		Status(http.StatusCreated).
		Body("").
		End()

	jwt := result.Response.Header.Get("Authorization")

	result = apitest.New().
		Handler(app).
		Post("/servicepage/create").
		JSON(inputServicePageMock).
		Header("Authorization", jwt).
		Expect(t).
		Status(http.StatusCreated).
		Body("").
		End()

	apitest.New().
		Handler(app).
		Get("/servicepage/myservices").
		Header("Authorization", jwt).
		Expect(t).
		Status(http.StatusOK).
		Body(`{"services":[{"Id":1,"Name":"SimpleService","Description":"MockServiceCreated","Image":"SimplePathFromImage","Value":80.8,"PositiveEvaluations":null,"NegativeEvaluations":null}]}`).
		End()

}

func TestSimpleCreateCommitInServicePage(t *testing.T) {
	inputDTO := dto.LoginDTO{
		Login:    "example7@hotmail.com",
		Password: "123qwe",
		User: dto.UserDTO{
			Name: "userTest7",
			Cpf:  "userCPFTest7",
		},
	}

	inputServicePageMock := dto.ServicePageInput{
		Name:        "SimpleServiceForReceiveCommit",
		Image:       "SimplePathFromImage",
		Value:       80.80,
		Description: "MockServiceCreated",
	}

	commit := dto.SimpleCommitInput{
		Commit: "A Simple Commit for test",
	}

	result := apitest.New().
		Handler(app).
		Post("/login/create").
		JSON(inputDTO).
		Expect(t).
		Status(http.StatusCreated).
		Body("").
		End()

	jwt := result.Response.Header.Get("Authorization")

	apitest.New().
		Handler(app).
		Post("/servicepage/create").
		JSON(inputServicePageMock).
		Header("Authorization", jwt).
		Expect(t).
		Status(http.StatusCreated).
		Body("").
		End()

	apitest.New().
		Handler(app).
		Post("/servicepage/all/1/commit").
		JSON(commit).
		Header("Authorization", jwt).
		Expect(t).
		Status(http.StatusCreated).
		End()

}
