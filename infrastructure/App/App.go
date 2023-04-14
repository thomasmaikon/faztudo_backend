package App

import (
	"log"
	"net/http"
	"projeto/FazTudo/controller"
	"projeto/FazTudo/entitys"
	"projeto/FazTudo/services/loginService"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type appEngine struct {
	Router *gin.Engine
}

func NewApp() *appEngine {
	router := gin.Default()
	cors.DefaultConfig()
	corsConfig := cors.New(
		cors.Config{
			AllowOrigins: []string{"http://localhost:3000"},
			AllowMethods: []string{"PUT", "PATCH"},
			AllowHeaders: []string{"Origin", "Content-Length", "Content-Type"},
			MaxAge:       12 * time.Hour,
		},
	)
	router.Use(corsConfig)
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

func (app *appEngine) RunMigrations(db *gorm.DB) *appEngine {

	err := db.AutoMigrate(&entitys.LoginUser{})
	if err != nil {
		log.Fatalf(err.Error())
	}

	err = db.AutoMigrate(&entitys.ServicePage{})
	if err != nil {
		log.Fatalf(err.Error())
	}

	err = db.AutoMigrate(&entitys.Commit{})
	if err != nil {
		log.Fatalf(err.Error())
	}

	err = db.AutoMigrate(&entitys.Like{})
	if err != nil {
		log.Fatalf(err.Error())
	}

	/*array := []string{
		`CREATE TABLE IF NOT EXISTS public.credentials (
			ID SERIAL PRIMARY KEY,
			login varchar(450) UNIQUE,
			password varchar(450)
		);`,
		`CREATE TABLE IF NOT EXISTS public.service (
			ID SERIAL PRIMARY KEY,
			fk_login integer,
			CONSTRAINT fk_login FOREIGN KEY (fk_login) REFERENCES credentials (ID),
			name text,
			image text,
			description text,
			value double precision,
			positive_evaluations integer,
			negative_evaluations integer
		);`,
		`CREATE TABLE IF NOT EXISTS public.commit (
			ID SERIAL PRIMARY KEY,
			fk_login integer,
			CONSTRAINT fk_login FOREIGN KEY (fk_login) REFERENCES credentials (ID),
			fk_service_page integer,
			CONSTRAINT fk_service_page FOREIGN KEY (fk_service_page) REFERENCES service (ID),
			commit text
		)`,
		`CREATE TABLE IF NOT EXISTS public.likes (
			fk_login integer NOT NULL,
			fk_service_page integer NOT NULL,
			liker integer,
			PRIMARY KEY(fk_login, fk_service_page)
		)`,
	}

	for _, migrate := range array {
		_, err := db.Exec(migrate)

		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(migrate)
	}*/

	return app
}

func (app *appEngine) Run(port string) *appEngine {
	app.Router.Run(port)

	return app
}
