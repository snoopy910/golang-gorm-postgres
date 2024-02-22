package models

import (
	"gorm.io/gorm"
)

type Indego struct {
	gorm.Model
	LastUpdated string `gorm:"type:varchar(50)" json:"last_updated"`
	Features	JSONB  `gorm:"type:jsonb" json:"features"`
 }
