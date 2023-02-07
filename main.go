package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"projeto/FazTudo/infrastructure/App"

	"github.com/joho/godotenv"

	_ "github.com/lib/pq"
)

func main() {

	host, port, user, password, dbname := readEnvironments("config.env")

	Db := openDatabaseConnection(host, port, user, password, dbname)

	defer Db.Close()

	App.NewApp().
		InitializeRoutes().
		InitializeDatabase(Db).
		Run(":8080")

}

func readEnvironments(file string) (string, string, string, string, string) {

	err := godotenv.Load("config.env")
	if err != nil {
		log.Fatal("Env File doesnot Find")
	}

	host := os.Getenv("host")
	port := os.Getenv("port")
	user := os.Getenv("user")
	password := os.Getenv("password")
	dbname := os.Getenv("dbname")

	return host, port, user, password, dbname
}

func openDatabaseConnection(host, port, user, password, dbname string) *sql.DB {
	psqlconn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlconn)

	if err != nil {
		panic(err)
	}

	return db
}
