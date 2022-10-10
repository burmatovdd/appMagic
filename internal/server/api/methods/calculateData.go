package methods

import (
	"appMagic/internal/server/api/methods/helpers"
	"appMagic/internal/server/models"
	"fmt"
)

func (service *Service) Calculate(data models.Eth) (
	map[string]float64,
	map[string][]float64,
	map[string]map[string]int,
	float64) {
	TimeMethod := helpers.Service{}

	monthsMap := TimeMethod.CreateMapWithoutArray(12)
	dayMap := TimeMethod.CreateMap(31)
	hourMap := TimeMethod.CreateHourMap(24)

	var total float64
	for i := 0; i < len(data.Ethereum.Transactions); i++ {
		monthsMap[data.Ethereum.Transactions[i].Time[3:5]] += data.Ethereum.Transactions[i].GasValue
		dayMap[data.Ethereum.Transactions[i].Time[6:8]][0] += data.Ethereum.Transactions[i].GasPrice
		dayMap[data.Ethereum.Transactions[i].Time[6:8]][1] += 1
		total = data.Ethereum.Transactions[i].GasPrice*data.Ethereum.Transactions[i].GasValue + total
		key := data.Ethereum.Transactions[i].Time[9:14]
		price := fmt.Sprintf("%.0f", data.Ethereum.Transactions[i].GasPrice)
		if _, ok := hourMap[key]; ok {
			if _, k := hourMap[key][price]; k {
				hourMap[key][price]++
				continue
			}
			hourMap[key][price] = 1
		}
	}
	return monthsMap, dayMap, hourMap, total
}
