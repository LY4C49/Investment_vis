package controllers

import (
	"github.com/gin-gonic/gin"
	"server/model"
)

func Current(c *gin.Context) {
	res_data := new(model.Rank)

	res_data.Stock = 60
	res_data.Option = 80
	res_data.Cash = 10
	ResponseSuccess(c, res_data)
	return
}
