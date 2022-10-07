package methods

import (
	appMagicConfig "appMagic/configs/server"
	"appMagic/internal/server/api/methods/helpers"
	"appMagic/internal/server/frequency"
	"appMagic/internal/server/gasPerDay"
	"appMagic/internal/server/models"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"log"
	"net/http"
	"sort"
)

// todo: частотное распределение, например сколько раз за весь период встретилась одинаковая цена

type Service struct {
	service *Json
}

type Json interface {
	GetData(c *gin.Context)
}

func (service *Service) GetData(c *gin.Context) {
	model := models.Eth{}
	method := appMagicConfig.Service{}
	TimeMethod := helpers.Service{}
	Frequency := frequency.Service{}
	dayMethod := gasPerDay.Service{}

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

	monthsMap := TimeMethod.CreateMapWithoutArray(12)
	dayMap := TimeMethod.CreateMap(31)
	hourMap := TimeMethod.CreateHourMap(24)

	var total float64
	for i := 0; i < len(model.Ethereum.Transactions); i++ {
		monthsMap[model.Ethereum.Transactions[i].Time[3:5]] += model.Ethereum.Transactions[i].GasValue
		dayMap[model.Ethereum.Transactions[i].Time[6:8]][0] += model.Ethereum.Transactions[i].GasPrice
		dayMap[model.Ethereum.Transactions[i].Time[6:8]][1] += 1
		total = model.Ethereum.Transactions[i].GasPrice*model.Ethereum.Transactions[i].GasValue + total
		for key, _ := range hourMap {
			if key == model.Ethereum.Transactions[i].Time[9:14] {
				hourMap[key] = append(hourMap[key], models.FrequencyInfo{
					Value: int(model.Ethereum.Transactions[i].GasPrice),
					Count: 0,
				})
			}
		}
	}

	newHourMap := Frequency.Count(hourMap)
	dayAverage := dayMethod.CountAverage(dayMap)
	//
	freqArray := []models.Frequency{}

	keys := make([]string, 0, len(newHourMap))
	for k := range newHourMap {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	for _, key := range keys {
		freqArray = append(freqArray, models.Frequency{
			Time:          key,
			FrequencyInfo: newHourMap[key],
		})
	}

	data := models.Data{
		monthsMap,
		dayAverage,
		freqArray,
		total,
	}

	newData, err := json.Marshal(data)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("data: ", string(newData))

}
