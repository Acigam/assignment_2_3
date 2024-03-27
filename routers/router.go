package routers

import (
	"assignment_2_3/controllers"

	"github.com/gin-gonic/gin"
)

func StartApp() *gin.Engine {
	router := gin.Default()

	router.POST("/orders", controllers.CreateOrder)
	router.GET("/orders", controllers.GetOrders)
	router.PUT("/orders/:orderId", controllers.UpdateOrder)
	router.DELETE("/orders/:orderId", controllers.DeleteOrder)

	return router
}
