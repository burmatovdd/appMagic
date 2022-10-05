package gasPerDay

import (
	"appMagic/internal/server/api/methods/helpers"
	"strconv"
)

type Service struct {
	service *Average
}

type Average interface {
	CountAverage(map[string][]float64) map[string]float64
}

func (service *Service) CountAverage(dayMap map[string][]float64) map[string]float64 {
	method := helpers.Service{}
	averageMap := method.CreateMapWithoutArray(31)

	for i := 1; i <= len(dayMap); i++ {
		if i < 10 {
			averageMap["0"+strconv.Itoa(i)] = dayMap["0"+strconv.Itoa(i)][0] / dayMap["0"+strconv.Itoa(i)][1]
			continue
		}
		averageMap[strconv.Itoa(i)] = dayMap[strconv.Itoa(i)][0] / dayMap[strconv.Itoa(i)][1]
	}
	return averageMap
}
