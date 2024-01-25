package controllers

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"server/logic"
	"server/model"
	"strconv"
)

func Rank(c *gin.Context) {
	stock_param := c.Query("stock")
	option_param := c.Query("option")
	cash_param := c.Query("cash")

	stock_f, err := strconv.ParseFloat(stock_param, 64)
	option_f, err := strconv.ParseFloat(option_param, 64)
	cash_f, err := strconv.ParseFloat(cash_param, 64)

	res_data := new(model.Rank)

	if err != nil {
		res_data.Stock = 5
		res_data.Option = 5
		res_data.Cash = 5
		ResponseSuccess(c, res_data)
		zap.L().Error("Params Invalid:", zap.Error(err))
		return
	}

	logic.Rank(res_data, stock_f, option_f, cash_f)
	ResponseSuccess(c, res_data)

}
