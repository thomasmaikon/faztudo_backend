package IntegrationTest

import (
	"database/sql"
	"log"
	"os"
	"projeto/FazTudo/infrastructure/App"
	"projeto/FazTudo/infrastructure/database"
	"projeto/FazTudo/test"
	"strings"
	"testing"

	_ "github.com/lib/pq"
)

var db *sql.DB

func TestMain(m *testing.M) {

	dockerInfo := test.StartPostgresDockerFormTest()
	strs := strings.Split(dockerInfo.HostAndPort, ":")
	port := strs[1]

	ormDB := database.GetDBWithParams(port, dockerInfo.User, dockerInfo.Password, dockerInfo.Dbname)
	// initialize app
	App.NewApp().RunMigrations(ormDB)

	//Run tests
	code := m.Run()

	if err := dockerInfo.Pool.Purge(dockerInfo.Resource); err != nil {
		log.Fatalf("Could not purge resource: %s", err)
	}

	os.Exit(code)
}
