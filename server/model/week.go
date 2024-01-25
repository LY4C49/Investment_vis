package model

import "time"

type Week struct {
	Id          int64     `db:"id"`
	Stock       float64   `db:"stock"`
	Option      float64   `db:"option"`
	Create_time time.Time `db:"create_time"`
}

type WeekRes struct {
	Stock       []float64 `json:"stock"`
	Option      []float64 `json:"option"`
	Total_asset []float64 `json:"total_asset"`
	Yield       []float64 `json:"yield"`
	XAxis       []string  `json:"xaxis"`
}
