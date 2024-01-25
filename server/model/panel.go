package model

type Panel struct {
	Stock  float64 `json:"stock"`
	Option float64 `json:"option"`
	Total  float64 `json:"total"`
}
