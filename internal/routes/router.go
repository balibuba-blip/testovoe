package routes

import (
	"test_task/internal/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()

	products := router.Group("/products")
	{
		products.GET("", controllers.GetProducts)
		products.GET("/:id", controllers.GetProduct)
		products.POST("", controllers.CreateProduct)
		products.PUT("/:id", controllers.UpdateProduct)
		products.DELETE("/:id", controllers.DeleteProduct)
	}

	measures := router.Group("/measures")
	{
		measures.POST("", controllers.CreateMeasure)
		//measures.GET("", controllers.GetAllMeasures)
		//measures.GET("/:id", controllers.GetMeasure)
		//measures.PUT("/:id", controllers.UpdateMeasure)
		//measures.DELETE("/:id", controllers.DeleteMeasure)
	}

	return router
}
