package database

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var dbConnection *gorm.DB

func GetDB() *gorm.DB {

	if dbConnection == nil {
		host, port, user, password, dbname := readEnvironments("config.env")

		dbConnection = openDatabaseConnection(host, port, user, password, dbname)
	}

	return dbConnection
}

func GetDBWithParams(port string, user, password, dbname string) *gorm.DB {

	if dbConnection == nil {

		dbConnection = openDatabaseConnection("localhost",
			port,
			user,
			password,
			dbname)
	}

	return dbConnection
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

func openDatabaseConnection(host, port, user, password, dbname string) *gorm.DB {
	psqlconn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := gorm.Open(postgres.Open(psqlconn), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	sql, err := db.DB()
	if err != nil {
		panic(err)
	}

	sql.SetMaxOpenConns(5)
	sql.SetConnMaxIdleTime(5)
	sql.SetConnMaxLifetime(time.Hour)

	return db
}
