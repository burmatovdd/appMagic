package models

type Data struct {
	GasPerMonth           map[string]float64 `json:"gas_per_month"`
	AverageSumPerDay      map[string]float64 `json:"average_sum_per_day"`
	FrequencyDistribution string             `json:"frequency_distribution"`
	Total                 float64            `json:"total"`
}
