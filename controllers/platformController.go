package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/mikaellpc4/go-payments-api/models"
	"github.com/mikaellpc4/go-payments-api/initializers"
)

func FindPlatformByID(c *gin.Context) {
	id := c.Param("id")

	var platform models.Platform
	result := initializers.DB.First(&platform, id)

	if result.Error != nil {
		c.Status(400)
		return
	}

	c.JSON(200, gin.H{
		"platform": platform,
	})
}

func FindPlatformBySlug(c *gin.Context) {
	slug := c.Param("slug")

	var platform models.Platform
	result := initializers.DB.Where("slug = ?", slug).First(&platform)

	if result.Error != nil {
		c.Status(400)
		return
	}

	c.JSON(200, gin.H{
		"platform": platform,
	})
}

func IndexPlatforms(c *gin.Context) {
	var platforms []models.Platform
	result := initializers.DB.Find(&platforms)

	if result.Error != nil {
		c.Status(400)
		return
	}

	c.JSON(200, gin.H{
		"platforms": platforms,
	})
}

func CreatePlatform(c *gin.Context) {
	var body struct {
		Name       string
		Domain     string
		WebhookUrl string
		Slug       string
	}

	c.Bind(&body)

	newPlatform := models.Platform{
		Name:       body.Name,
		Domain:     body.Domain,
		WebhookUrl: body.WebhookUrl,
		Slug:       body.Slug,
	}
	result := initializers.DB.Create(&newPlatform)

	if result.Error != nil {
		c.Status(400)
		return
	}

	c.JSON(200, gin.H{
		"platform": newPlatform,
	})
}

func UpdatePlatform(c *gin.Context) {
	id := c.Param("id")

	var body struct {
		Name       *string
		Domain     *string
		WebhookUrl *string
	}
	c.Bind(&body)

	var platform models.Platform
	initializers.DB.First(&platform, id)

	if body.Name != nil {
		platform.Name = *body.Name
	}

	if body.Domain != nil {
		platform.Name = *body.Domain
	}

	if body.WebhookUrl != nil {
		platform.Name = *body.WebhookUrl
	}

	result := initializers.DB.Model(&platform).Updates(models.Platform{
		Name:       platform.Name,
		Domain:     platform.Domain,
		WebhookUrl: platform.WebhookUrl,
	})

	if result.Error != nil {
		c.Status(400)
	}

	c.JSON(200, gin.H{
		"platform": platform,
	})
}

func DeletePlatform(c *gin.Context) {
	id := c.Param("id")

	result := initializers.DB.Delete(&models.Platform{}, id)

	if result.Error != nil {
		c.Status(400)
	}

	c.Status(200)
}
