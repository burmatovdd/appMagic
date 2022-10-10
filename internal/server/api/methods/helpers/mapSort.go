package helpers

import "sort"

type Number interface {
	map[string]float64 | map[string][]int
}

func (service *Service) Sort(m map[string]float64) []string {
	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	return keys
}

func (service *Service) SortWithArray(m map[string][]int) []string {
	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	return keys
}
