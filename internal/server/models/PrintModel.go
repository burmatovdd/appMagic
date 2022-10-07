package models

type Data struct {
	GasPerMonth           map[string]float64 `json:"gas_per_month"`
	AverageSumPerDay      map[string]float64 `json:"average_sum_per_day"`
	FrequencyDistribution []Frequency        `json:"frequency_distribution"`
	Total                 float64            `json:"total"`
}

type Frequency struct {
	Time          string          `json:"time"`
	FrequencyInfo []FrequencyInfo `json:"frequencyInfo"`
}

type FrequencyInfo struct {
	Value int `json:"value"`
	Count int `json:"count"`
}
