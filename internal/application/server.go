package application

import (
	"company/controllers"
	"company/middlewares"
	"log"
	"os"

	"github.com/gin-gonic/gin"
)

// StartServer is a function for launching server
func (app *Application) StartServer() error {

	//initiate server instance
	router := gin.Default()

	//add routes for api version v1
	api := router.Group("/v1")
	{
		service := controllers.NewController(app.DB)

		api.GET("/healthcheck", service.HealthCheck)
		api.POST("/token", service.GenerateToken)
		api.POST("/user", service.RegisterUser)
		api.POST("/company", service.RegisterCompany).Use(middlewares.Auth())
		api.GET("/company/:id", service.GetCompany).Use(middlewares.Auth())
		api.PATCH("/company/:id", service.UpdateCompany).Use(middlewares.Auth())
		api.DELETE("/company/:id", service.DeleteCompany).Use(middlewares.Auth())
	}

	//run server instance
	if err := router.Run(os.Getenv("API_ENDPOINT")); err != nil {
		log.Fatal(err)
	}

	return nil
}
