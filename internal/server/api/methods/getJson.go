package methods

import (
	appMagicConfig "appMagic/configs/server"
	"appMagic/internal/server/models"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"log"
	"net/http"
)

type Service struct {
	service *Json
}

type Json interface {
	GetData(c *gin.Context)
}

func (service *Service) GetData(c *gin.Context) {
	model := models.Eth{}
	method := appMagicConfig.Service{}

	config, err := method.Loader("configs/server")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}

	response, err := http.Get(config.GETJSON)
	if err != nil {
		log.Fatal(err)
	}

	byteValue, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	err = json.Unmarshal(byteValue, &model)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(model)
}
