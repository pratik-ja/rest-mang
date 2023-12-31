package routes

import (
	"github.com/gin-gonic/gin"
	controller "rest-mang/controllers"
)

func OrderRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.GET("/orders", controller.GetOrders())
	incomingRoutes.GET("/order/:order_id", controller.GetOrder())
	incomingRoutes.POST("/order", controller.CreateOrder())
	incomingRoutes.PATCH("/orders/:order_id", controller.UpdateOrder())
}
