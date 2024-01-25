package model

type Rank struct {
	Stock  float64 `json:"stock"`
	Option float64 `json:"option"`
	Cash   float64 `json:"cash"`
}
