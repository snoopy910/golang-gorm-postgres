package handlers

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/snoopy910/golang-gorm-postgres/models"
)

func getIndegoJson() ([]models.Indego, error) {
	indegoApiUrl := "https://bts-status.bicycletransit.workers.dev/phl"

	res, err := http.Get(indegoApiUrl)
	if err != nil {
		return  nil, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	var indegoData []models.Indego
	err = json.Unmarshal(body, &indegoData)
	if err != nil {
		return nil, err
	}

	return indegoData, nil
}

