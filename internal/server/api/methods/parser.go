package methods

import (
	"appMagic/internal/server/models"
	"encoding/json"
	"log"
)

func (service *Service) ParseData(byteValue []byte) models.Eth {
	model := models.Eth{}
	err := json.Unmarshal(byteValue, &model)
	if err != nil {
		log.Fatal(err)
	}
	return model
}
