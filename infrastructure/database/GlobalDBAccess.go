package database

import "database/sql"

var globaldbacess *sql.DB

func SetDB(db *sql.DB) {
	globaldbacess = db
}

func GetDBAccess() *sql.DB {
	return globaldbacess
}
