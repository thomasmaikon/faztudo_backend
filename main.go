package main

import (
	"projeto/FazTudo/infrastructure/App"
	"projeto/FazTudo/infrastructure/database"

	_ "github.com/lib/pq"
)

func main() {

	db := database.GetDB()
	defer db.DB()

	App.NewApp().
		InitializeRoutes().
		RunMigrations(db).
		Run(":8080")

}
