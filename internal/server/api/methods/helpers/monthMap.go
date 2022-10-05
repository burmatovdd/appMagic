package helpers

import "strconv"

func (service *Service) CreateMapWithoutArray(n int) map[string]float64 {
	m := make(map[string]float64)
	for i := 1; i <= n; i++ {
		if i < 10 {
			m["0"+strconv.Itoa(i)] = 0
			continue
		}
		m[strconv.Itoa(i)] = 0
	}
	return m
}
