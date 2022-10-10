package helpers

import (
	"strconv"
)

type Service struct {
	service *Map
}

type Map interface {
	CreateMap(n int) []string
	CreateMonthMap() map[string]float64
	CreateHourMap() map[string]float64
	Sort(m map[string]float64) []string
	SortWithArray(m map[string][]float64) []string
}

func (service *Service) CreateMap(n int) map[string][]float64 {
	m := make(map[string][]float64)
	for i := 1; i <= n; i++ {
		if i < 10 {
			m["0"+strconv.Itoa(i)] = []float64{0, 0}
			continue
		}
		m[strconv.Itoa(i)] = []float64{0, 0}
	}
	return m
}

func (service *Service) CreateHourMap(n int) map[string]map[string]int {
	m := make(map[string]map[string]int)
	for i := 0; i < n; i++ {
		if i < 10 {
			m["0"+strconv.Itoa(i)+":00"] = map[string]int{}
			continue
		}
		m[strconv.Itoa(i)+":00"] = map[string]int{}
	}
	return m
}
