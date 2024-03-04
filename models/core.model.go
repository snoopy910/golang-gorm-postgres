package models

import (
	"time"

	"github.com/google/uuid"
)

// DB struct for auto-migration
type Core struct{
	ID uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primary_key"`
	LastUpdated time.Time `gorm:"type:timestamp with time zone"`
	Features []JSONB `gorm:"type:jsonb"`
	Type string `gorm:"type:varchar(50)"`
	Weather JSONB `gorm:"type:json"`
}

// Struct to unMarshal the result from indego API
type Indego struct{
	ID uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primary_key"`
	LastUpdated time.Time `gorm:"type:timestamp with time zone" json:"last_updated"`
	Features []JSONB `gorm:"type:jsonb" json:"features"`
	Type string `gorm:"type:varchar(50)" json:"type"`
}

// Struct to insert data into PostgreSQL
type DataToInsert struct{
	ID uuid.UUID
	LastUpdated time.Time
	Features string
	Type string
	Weather JSONB
}

// Strcut to call data from PostgreSQL
type DataToCall struct{
	ID uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primary_key"`
	LastUpdated time.Time `gorm:"type:timestamp with time zone" json:"last_updated"`
	Features []byte `gorm:"type:jsonb" json:"features"`
	Type string `gorm:"type:varchar(50)" json:"type"`
	Weather JSONB `gorm:"type:json" json:"weather"`
}


func (DataToInsert) TableName() string {
	return "cores"
}