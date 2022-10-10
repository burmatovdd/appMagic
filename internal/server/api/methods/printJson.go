package methods

import (
	"appMagic/internal/server/models"
	"bytes"
	"encoding/json"
	"log"
)

func (service *Service) PrintJson(data models.Data) string {
	newData, err := json.Marshal(data)
	if err != nil {
		log.Fatal(err)
	}

	return jsonPrettyPrint(string(newData))

}

func jsonPrettyPrint(in string) string {
	var out bytes.Buffer
	err := json.Indent(&out, []byte(in), "", "\t")
	if err != nil {
		return in
	}
	return out.String()
}
