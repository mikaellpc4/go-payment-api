package models

import (
	"gorm.io/gorm"
)

type Platform struct {
	gorm.Model

	Name       string
	Domain     string
	WebhookUrl string
	Slug       string `gorm:"unique"`

	Payments []Payment `gorm:"foreignKey:PlatformId"`
}
