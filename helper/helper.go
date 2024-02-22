package helper

import (
	"io"
	"net/http"

	// "github.com/snoopy910/golang-gorm-postgres/initializers"
	"github.com/snoopy910/golang-gorm-postgres/models"
)

func getIndegoJSON () error {
	indegoApiUrl := "https://bts-status.bicycletransit.workers.dev/phl"

	res, err := http.Get(indegoApiUrl)
	if err != nil {
		return  err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return err
	}

	var apiLatest models.JSONB
	return apiLatest.Scan(body)
}

func getWeatherJSON ([]models.Indego) error {
	openWeatherApiUrl := "https://api.openweathermap.org/data/2.5/weather?q=philadelphia&appid=d2166073a527e474bdb3141103689b9b"
	
	res, err := http.Get(openWeatherApiUrl)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return err
	}

	var apiLatest models.JSONB
	return apiLatest.Scan(body)
	
}

	