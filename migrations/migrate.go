package main

import (
	"github.com/mikaellpc4/goPaymentWebhook/initializers"
	"github.com/mikaellpc4/goPaymentWebhook/models"
)

func init() {
  initializers.LoadEnv()
  initializers.ConnectToDB()
}

func main() {
  initializers.DB.AutoMigrate(&models.Platform{}, &models.Payment{})
}
