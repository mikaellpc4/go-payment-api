package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mikaellpc4/go-payments-api/models"
	"github.com/mikaellpc4/go-payments-api/services"
	"github.com/mikaellpc4/go-payments-api/initializers"
)

func FindPaymentByID(c *gin.Context) {
	id := c.Param("id")

	var payment models.Payment
	result := initializers.DB.First(&payment, id)

	if result.Error != nil {
		c.Status(400)
		return
	}

	c.JSON(200, gin.H{
		"payment": payment,
	})
}

func IndexPayments(c *gin.Context) {
	var payments []models.Payment
	result := initializers.DB.Find(&payments)

	if result.Error != nil {
		c.Status(400)
		return
	}

	c.JSON(200, gin.H{
		"payments": payments,
	})
}

func CreatePayment(c *gin.Context) {
	var body struct {
		Value          int    `json:"value" binding:"required,min=1"`
		Name           string `json:"name" binding:"required"`
		DocumentNumber string `json:"documentNumber" binding:"required"`
		PlatformSlug   string `json:"platformSlug" binding:"required"`
	}

	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var platform models.Platform

	platformResult := initializers.DB.Where("slug = ? ", body.PlatformSlug).Find(&platform)

	if platformResult.Error != nil || platform.ID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Platform doesn't exists",
		})
		return
	}

	pixData := services.CreatePix{
		Platform: platform,
		Amount:   body.Value,
		Name:     body.Name,
		Document: body.DocumentNumber,
	}

	pix, err := services.CreatePIX(pixData)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	newPayment := models.Payment{
		Value: body.Value,

		Name:           body.Name,
		DocumentNumber: body.DocumentNumber,
		QrCode:         pix.Data.QRCodeURL,
		TransactionId:  pix.Data.ExternalReference,

		PlatformId: platform.ID,
	}

	result := initializers.DB.Create(&newPayment)

	if result.Error != nil {
		c.Status(http.StatusBadRequest)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"payment": newPayment,
	})
}

func UpdatePayment(c *gin.Context) {
	id := c.Param("id")

	var body struct {
		Value *int

		Name           *string
		DocumentNumber *string
		QrCode         *string
		TransactionId  *string

		Paid *bool

		PlatformId *uint
	}
	c.Bind(&body)

	var payment models.Payment
	initializers.DB.First(&payment, id)

	if body.Value != nil {
		payment.Value = *body.Value
	}

	if body.Name != nil {
		payment.Name = *body.Name
	}

	if body.DocumentNumber != nil {
		payment.DocumentNumber = *body.DocumentNumber
	}

	if body.QrCode != nil {
		payment.QrCode = *body.QrCode
	}

	if body.TransactionId != nil {
		payment.TransactionId = *body.TransactionId
	}

	if body.Paid != nil {
		payment.Paid = *body.Paid
	}

	if body.PlatformId != nil {
		payment.PlatformId = *body.PlatformId
	}

	result := initializers.DB.Model(&payment).Updates(models.Payment{
		Value: payment.Value,

		Name:           payment.Name,
		DocumentNumber: payment.DocumentNumber,
		QrCode:         payment.QrCode,
		TransactionId:  payment.TransactionId,

		Paid: payment.Paid,

		PlatformId: payment.PlatformId,
	})

	if result.Error != nil {
		c.Status(400)
	}

	c.JSON(200, gin.H{
		"payment": payment,
	})
}

func DeletePayment(c *gin.Context) {
	id := c.Param("id")

	result := initializers.DB.Delete(&models.Payment{}, id)

	if result.Error != nil {
		c.Status(400)
	}

	c.Status(200)
}
