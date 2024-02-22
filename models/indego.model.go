package models

import (
	"github.com/google/uuid"
)

type Indego struct {
	ID uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primary_key"`
	LastUpdated string `gorm:"type:varchar(50)" json:"last_updated"`
	Features	JSONB  `gorm:"type:jsonb" json:"features"`
	Type 		string `gorm:"type:string" json:"type"`
 }
