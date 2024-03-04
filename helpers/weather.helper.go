package helpers

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/snoopy910/golang-gorm-postgres/initializers"
	"github.com/snoopy910/golang-gorm-postgres/models"
)

func GetOpenWeather() (models.JSONB, error) {
	fmt.Println("? Called API for open weather")
	config, err := initializers.LoadConfig(".")
	if err != nil {
		log.Fatal("? Could not lod environment variables", err)
	}
	openWeatherApiUrl := "https://api.openweathermap.org/data/2.5/weather?q="+config.City+"&appid="+config.OpenWeatherApiKey


	res, err := http.Get(openWeatherApiUrl)
	if err != nil {
		return  models.JSONB{}, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return models.JSONB{}, err
	}

	var openWeatherApi models.JSONB

	err = json.Unmarshal(body, &openWeatherApi)
	if err != nil {
		return models.JSONB{}, err
	}

	// fmt.Println("? type of features is ", string(features))
	return openWeatherApi, nil
}