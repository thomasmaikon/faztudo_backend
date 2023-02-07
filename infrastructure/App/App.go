package App

import (
	"database/sql"
	"net/http"
	"projeto/FazTudo/controller"
	"projeto/FazTudo/infrastructure/database"
	"projeto/FazTudo/infrastructure/migrations"
	"projeto/FazTudo/services/loginService"

	"github.com/gin-gonic/gin"
)

type appEngine struct {
	Router *gin.Engine
}

func NewApp() *appEngine {
	router := gin.Default()
	return &appEngine{router}
}

func (app *appEngine) InitializeRoutes() *appEngine {

	app.Router.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{})

	})

	app.Router.POST("/login/credential", controller.ValidateCrendential)
	app.Router.POST("/login/create", controller.CreateLogin)

	app.Router.GET("/servicepage/all/:index", controller.GetAllServicePage)

	app.Router.POST("/servicepage/all/:id/commit", loginService.IsAuthorized, controller.CreateCommitInservicePage)
	app.Router.POST("/servicepage/create", loginService.IsAuthorized, controller.CreateServicePage)
	app.Router.GET("/servicepage/myservices", loginService.IsAuthorized, controller.GetMyServicesPage)

	return app
}

func (app *appEngine) InitializeDatabase(db *sql.DB) *appEngine {

	migrations.RunMigrations(db)
	database.SetDB(db)
	return app
}

func (app *appEngine) Run(port string) *appEngine {
	app.Router.Run(port)

	return app
}
