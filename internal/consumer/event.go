package consumer

import "time"

type Event struct {
	BloodTestID string    `json:"bloodTestId"`
	UserID      string    `json:"userId"`
	Timestamp   time.Time `json:"timestamp"`
	Results     []Result  `json:"results"`
}

type Result struct {
	AnalyteCode string  `json:"analyteCode"`
	Value       float64 `json:"value"`
	NormalLow   float64 `json:"normalLow"`
	NormalHigh  float64 `json:"normalHigh"`
	Unit        string  `json:"unit"`
}
