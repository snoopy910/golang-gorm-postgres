package helpers

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/snoopy910/golang-gorm-postgres/models"
)

func GetIndegoJson() (models.Indego, error) {
	fmt.Println("? Called API for indego status")
	indegoApiUrl := "https://bts-status.bicycletransit.workers.dev/phl"

	res, err := http.Get(indegoApiUrl)
	if err != nil {
		return  models.Indego{}, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return models.Indego{}, err
	}

	var indegoApi models.Indego

	err = json.Unmarshal(body, &indegoApi)
	if err != nil {
		return models.Indego{}, err
	}

	// fmt.Println("? type of features is ", string(features))
	return indegoApi, nil
}