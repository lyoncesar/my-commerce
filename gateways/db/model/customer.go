package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Customer struct {
	gorm.Model
	UUID     uuid.UUID `gorm:"uniqueIndex:idx_customer_uuid,unique"`
	Name     string
	Email    string
	Document string `gorm:"uniqueIndex:idx_customer_document,unique"`
}
