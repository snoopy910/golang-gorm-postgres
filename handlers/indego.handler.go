package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/google/uuid"
	"github.com/snoopy910/golang-gorm-postgres/models"
)

type Indego struct {
	ID 			uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primary_key"`
	LastUpdated string `gorm:"type:varchar(50)" json:"last_updated"`
	Features	[]models.JSONB  `gorm:"type:jsonb" json:"features"`
	Type 		string `gorm:"type:string" json:"type"`
 }

func GetIndegoJson() (Indego, error) {
	fmt.Println("? Called API for indego status")
	indegoApiUrl := "https://bts-status.bicycletransit.workers.dev/phl"

	res, err := http.Get(indegoApiUrl)
	if err != nil {
		return  Indego{}, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return Indego{}, err
	}
	// fmt.Println(string(body))

	var indegoApi Indego
	err = json.Unmarshal(body, &indegoApi)
	if err != nil {
		return Indego{}, err
	}

	// fmt.Println(indegoApi)
	return indegoApi, nil
}

