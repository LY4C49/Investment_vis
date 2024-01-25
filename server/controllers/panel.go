package controllers

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"server/logic"
	"server/model"
	"server/tools"
	"strconv"
)

func Panel(c *gin.Context) {

	stock_param := c.Query("stock")
	option_param := c.Query("option")
	total_param := c.Query("total")

	stock_f, err := strconv.ParseFloat(stock_param, 64)
	option_f, err := strconv.ParseFloat(option_param, 64)
	total_f, err := strconv.ParseFloat(total_param, 64)

	res_data := new(model.Panel)

	if err != nil {
		res_data.Stock = 5
		res_data.Option = 5
		res_data.Total = 5
		ResponseSuccess(c, res_data)
		zap.L().Error("Params Invalid:", zap.Error(err))
		return
	}

	logic.Panel(res_data, stock_f, option_f, total_f)
	ResponseSuccess(c, res_data)
}

func StockIndex(c *gin.Context) {

	sh_param := c.Query("sh")
	nq_param := c.Query("nq")

	sh_int, err := strconv.ParseInt(sh_param, 10, 64)
	nq_int, err := strconv.ParseInt(nq_param, 10, 64)

	res_data := new(model.Stock_index)

	if err != nil {
		res_data.ShangHai = 2600
		res_data.Nasdaq = 15000
		ResponseSuccess(c, res_data)
		zap.L().Error("Params Invalid:", zap.Error(err))
		return
	}

	sh_index := tools.GetRandomInt(40)
	nq_index := tools.GetRandomInt(40)

	res_data.ShangHai = int64(sh_index) + sh_int
	res_data.Nasdaq = int64(nq_index) + nq_int

	ResponseSuccess(c, res_data)
}
