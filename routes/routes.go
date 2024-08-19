package routes

import (
	"sagara-msib-test/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine, clothController *controllers.ClothController) {
	clothRoutes := router.Group("/clothes")
	{
		clothRoutes.POST("/", clothController.CreateCloth)
		clothRoutes.GET("/low-stock", clothController.FindLowStockClothes)
		clothRoutes.GET("/out-of-stock", clothController.FindOutOfStockClothes)
		clothRoutes.GET("/:id", clothController.GetCloth)
		clothRoutes.GET("/", clothController.GetClothes)
		clothRoutes.PATCH("/:id/increment", clothController.IncrementClothStock)
		clothRoutes.PATCH("/:id/decrement", clothController.DecrementClothStock)
		clothRoutes.PUT("/:id", clothController.UpdateCloth)
		clothRoutes.DELETE("/:id", clothController.DeleteCloth)
	}
}
