package models

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"time"
)

// JSONB Interface for JSONB Field
type JSONB map[string]interface{}

// Value Marshal
func (a JSONB) Value() (driver.Value, error) {
    return json.Marshal(a)
}

// Scan Unmarshal
func (a *JSONB) Scan(value interface{}) error {
    b, ok := value.([]byte)
    if !ok {
        return errors.New("type assertion to []byte failed")
    }
    return json.Unmarshal(b,&a)
}

type ResponseByTime struct {
	At time.Time `json:"at"`
	Stations 	[]Feature `json:"stations"`
	Weather 	JSONB	`json:"weather"`
}

type ResponseByKiosk struct {
	At time.Time `json:"at"`
	Station 	Feature `json:"station"`
	Weather 	JSONB	`json:"weather"`
}

type Feature struct {
	Geometry string `json:"geometry"`
	Properties Property `json:"properties"`
	Type string `json:"type"`
}

type Property struct {
	ID 		uint64 `json:"id"`
	Name	string `json:"name"`
	Coordinates driver.Value `json:"coordinates"`
	TotalDocks uint8 `json:"totalDocks"`
	DocksAvailable uint8 `json:"docksAvailable"`
	BikesAvailable uint8 `json:"bikesAvailable"`
	ClassicBikesAvailable uint8 `json:"classicBikesAvailable"`
	SmartBikesAvailable uint8 `json:"smartBikesAvailable"`
	ElectricBikesAvailable uint8 `json:"electricBikesAvailable"`
	RewardBikesAvailable uint8 `json:"rewardBikesAvailable"`
	RewardDocksAvailable uint8 `json:"rewardDocksAvailable"`
	KioskStatus string `json:"kioskStatus"`
	KioskPublicStatus string `json:"kioskPublicStatus"`
	KioskConnectionStatus string `json:"kioskConnectionStatus"`
	KioskType uint8 `json:"kioskType"`
	AddressStreet string `json:"addressStreet"`
	AddressCity string `json:"addressCity"`
	AddressState string `json:"addressState"`
	AddressZipCode string `json:"addressZipCode"`
	Bikes []JSONB `json:"bikes"`
	CloseTime driver.Value `json:"closeTime"`
	EventEnd driver.Value `json:"eventEnd"`
	EventStart driver.Value `json:"eventStart"`
	IsEventBased bool `json:"isEventBased"`
	IsVirtual bool `json:"isVirtual"`
	KioskId uint64 `json:"kioskId"`
	Notes driver.Value `json:"notes"`
	OpenTime driver.Value `json:"openTime"`
	PublicText string `json:"publicText"`
	TimeZone driver.Value `json:"timeZone"`
	TrikesAvailable uint8 `json:"trikesAvailable"`
	Latitude float32 `json:"latitude"`
	Longitude float32 `json:"longitude"`
}