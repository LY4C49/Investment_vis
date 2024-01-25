package controllers

import "github.com/gin-gonic/gin"

type Asset_data struct {
	Xaxis []int     `json:"Xaxis"`
	Yaxis []float64 `json:"Yaxis"`
}

func Asset(c *gin.Context) {
	week := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	asset := []float64{100, 110, 115, 120, 121, 118, 119, 130, 150, 180}

	resp_data := new(Asset_data)
	resp_data.Xaxis = week
	resp_data.Yaxis = asset

	ResponseSuccess(c, resp_data)
}
