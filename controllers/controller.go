package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
	"github.com/snoopy910/golang-gorm-postgres/helpers"
	"github.com/snoopy910/golang-gorm-postgres/models"
)

type StationController struct {
	DB *gorm.DB
}

func NewStationController(DB *gorm.DB) StationController {
	return StationController{DB}
}

// This endpoint will be trigger every hour to fetch the data and insert it into database
func (sc *StationController) DataFetchAndStore(ctx *gin.Context) {
	indegoApi, err := helpers.GetIndegoJson()
	if err != nil {
		fmt.Println("? Could not call indego Api", err)
	}

	openWeatherApi, err := helpers.GetOpenWeather()
	if err!= nil {
		fmt.Println("? Could not call open weather Api", err)
	}

	uuid := uuid.New()

	features, err := json.Marshal(indegoApi.Features)
	if err != nil {
		return
	}

	newData := models.DataToInsert{
		ID: uuid,
		LastUpdated: indegoApi.LastUpdated,
		Features: string(features),
		Type: indegoApi.Type,
		Weather: openWeatherApi,
	}
	
	results := sc.DB.Create(&newData)

	// results := sc.DB.Exec("INSERT INTO cores (id, last_updated, features, type, weather) VALUES ($1, $2, $3, $4, $5)", uuid, indegoApi.LastUpdated, string(features), indegoApi.Type, openWeatherApi)
	if results.Error != nil {
		fmt.Println(results.Error)
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"status": "Fetch and Store Success"})
}

func (sc *StationController) GetByTime (ctx *gin.Context) {
	timeRequired, _ := time.Parse(time.RFC3339, ctx.Query("at"))
	var data models.DataToCall
	results := sc.DB.Where("last_updated > ?", timeRequired).First(&data)
	if results.Error != nil {
		fmt.Println(results.Error)
	}

	var stations []models.Feature
	json.Unmarshal(data.Features, &stations)
	response := models.ResponseByTime{
		At: data.LastUpdated,
		Stations: stations,
		Weather: data.Weather,
	}
	ctx.JSON(http.StatusOK, response)
}

func (sc *StationController) GetByKiosk (ctx *gin.Context) {
	timeRequired, _ := time.Parse(time.RFC3339, ctx.Query("at"))
	kioskId, _ := strconv.ParseUint(ctx.Param("kioskId"), 10, 64)
	var data models.DataToCall
	result := sc.DB.Where("last_updated > ?", timeRequired).First(&data)
	if result.Error != nil {
		fmt.Println(result.Error)
	}

	var stations []models.Feature
	var stationByKiosk models.Feature
	json.Unmarshal(data.Features, &stations)
	for _, station := range stations {
		if station.Properties.KioskId == kioskId {
			stationByKiosk = station
		}
	}
	response := models.ResponseByKiosk{
		At: data.LastUpdated,
		Station: stationByKiosk,
		Weather: data.Weather,
	}

	ctx.JSON(http.StatusOK, response)
}