package models

type HeartRateMetricExport struct {
	Source string  `json:"source,omitempty"`
	Date   string  `json:"date,omitempty"`
	Max    float64 `json:"max,omitempty"`
	Min    float64 `json:"min,omitempty"`
	Avg    float64 `json:"avg,omitempty"`
}

type HealthMetric[T any] struct {
	Units string `json:"units,omitempty"`
	Name  string `json:"name,omitempty"`
	Data  []T    `json:"data,omitempty"`
}

type HealthData[T any] struct {
	Metrics []HealthMetric[T] `json:"metrics"`
}

type HealthExport[T any] struct {
	Data HealthData[T] `json:"data"`
}
