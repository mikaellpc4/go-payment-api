package models

import (
	"gorm.io/gorm"
)

type Payment struct {
	gorm.Model

	Value int

	Name           string
	DocumentNumber string
	QrCode         string
	TransactionId  string

	Paid bool

	PlatformId uint
}
