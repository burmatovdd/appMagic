package helpers

import "strconv"

type Service struct {
	service *Map
}

type Map interface {
	CreateMap(n int) []string
	CreateMonthMap() map[string]float64
	Sort(m map[string]float64) []string
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
