package main

import (
	"fmt"
	"sagara-msib-test/config"
	"sagara-msib-test/controllers"
	"sagara-msib-test/database"
	"sagara-msib-test/repositories"
	"sagara-msib-test/routes"
	"sagara-msib-test/services"

	"github.com/gin-gonic/gin"
)

func main() {
	config.InitConfig()
	database.InitDatabase(config.AppConfig.Database.Driver, config.AppConfig.Database.URL)
	clothRepo := repositories.NewClothRepository(database.DB)
	clothService := services.NewClothService(clothRepo)
	clothHandler := controllers.NewClothController(clothService)

	router := gin.Default()
	routes.SetupRoutes(router, clothHandler)

	router.Run(fmt.Sprintf(":%d", config.AppConfig.App.Port))
}
