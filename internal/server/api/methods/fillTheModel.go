package methods

import (
	"appMagic/internal/server/gasPerDay"
	"appMagic/internal/server/models"
	"sort"
)

func (service *Service) FillModel(
	monthsMap map[string]float64,
	dayMap map[string][]float64,
	hourMap map[string]map[string]int,
	total float64) models.Data {
	dayMethod := gasPerDay.Service{}

	dayAverage := dayMethod.CountAverage(dayMap)
	freqArray := []models.Frequency{}

	keys := make([]string, 0, len(hourMap))
	for k := range hourMap {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	for _, key := range keys {
		freqArray = append(freqArray, models.Frequency{
			Time:          key,
			FrequencyInfo: hourMap[key],
		})
	}

	return models.Data{
		monthsMap,
		dayAverage,
		freqArray,
		total,
	}

}
