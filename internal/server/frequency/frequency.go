package frequency

import (
	"appMagic/internal/server/api/methods/helpers"
	"appMagic/internal/server/models"
)

type Service struct {
	service *Frequency
}
type Frequency interface {
	Count(m map[string][]float64) map[string][]float64
}

func (service *Service) Count(m map[string][]models.FrequencyInfo) map[string][]models.FrequencyInfo {
	TimeMethod := helpers.Service{}
	hourMap := TimeMethod.CreateHourMap(23)

	for key, _ := range m {
		for i := 0; i < len(m); i++ {
			k := checkInArray(m[key], m[key][i].Value)
			hourMap[key] = append(hourMap[key], models.FrequencyInfo{
				Value: m[key][i].Value,
				Count: k,
			})
		}

	}
	return hourMap
}

func checkInArray(array []models.FrequencyInfo, n int) int {
	k := 0
	for i := 0; i < len(array); i++ {
		if array[i].Value == n {
			k++
		}
	}
	return k
}
