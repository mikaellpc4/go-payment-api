package main

import (
	"github.com/gin-gonic/gin"
	"github.com/mikaellpc4/go-payments-api/controllers"
	"github.com/mikaellpc4/go-payments-api/initializers"
)

func init() {
	initializers.LoadEnv()
  initializers.ConnectToDB()
}

func main() {
	r := gin.Default()

	r.GET("/platforms", controllers.IndexPlatforms)
	r.POST("/platform", controllers.CreatePlatform)

  r.GET("/platform/:id", controllers.FindPlatformByID)
  r.GET("/platform/slug/:slug", controllers.FindPlatformBySlug)
  r.PUT("/platform/:id", controllers.UpdatePlatform)
  r.DELETE("/platform/:id", controllers.DeletePlatform)

	r.GET("/payments", controllers.IndexPayments)
	r.POST("/payment", controllers.CreatePayment)

  r.GET("/payment/:id", controllers.FindPaymentByID)
  r.PUT("/payment/:id", controllers.UpdatePayment)
  r.DELETE("/payment/:id", controllers.DeletePayment)

	r.Run()
}
