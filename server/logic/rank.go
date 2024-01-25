package logic

import (
	"server/model"
	"server/tools"
)

func Rank(res_data *model.Rank, stock_f float64, option_f float64, cash_f float64) {
	stock_random := tools.GetRandomFloat(10)
	option_random := tools.GetRandomFloat(10)
	cash_random := tools.GetRandomFloat(10)

	res_data.Stock = tools.Decimal(stock_random + stock_f)
	res_data.Option = tools.Decimal(option_random + option_f)
	res_data.Cash = tools.Decimal(cash_random + cash_f)

	if res_data.Stock > 100 {
		res_data.Stock -= 50
	} else {
		if res_data.Stock < 0 {
			res_data.Stock = 50
		}
	}

	if res_data.Cash > 100 {
		res_data.Cash -= 50
	} else {
		if res_data.Cash < 0 {
			res_data.Cash = 50
		}
	}

	if res_data.Option > 100 {
		res_data.Option -= 50
	} else {
		if res_data.Option < 0 {
			res_data.Option = 50
		}
	}
}
