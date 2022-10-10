package methods

import (
	appMagicConfig "appMagic/configs/server"
	"appMagic/internal/server/models"
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
	ParseData(byteValue []byte) models.Eth
	Calculate(data models.Eth)
	FillModel(monthsMap map[string]float64, dayMap map[string][]float64, hourMap map[string]map[string]int, total float64)
	PrintJson(data models.Data) string
}

func (service *Service) GetData(c *gin.Context) {
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

	data := service.ParseData(byteValue)
	monthData, dayData, hourData, total := service.Calculate(data)
	model := service.FillModel(monthData, dayData, hourData, total)
	fmt.Println(service.PrintJson(model))

}
