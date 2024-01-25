package logic

import (
	"server/model"
	"server/tools"
)

func Panel(res_data *model.Panel, stock_f float64, option_f float64, total_f float64) {
	stock_random := tools.GetRandomFloat(10)
	option_random := tools.GetRandomFloat(10)
	total_random := tools.GetRandomFloat(10)

	res_data.Stock = tools.Decimal(stock_random + stock_f)
	res_data.Option = tools.Decimal(option_random + option_f)
	res_data.Total = tools.Decimal(total_random + total_f)

	if res_data.Stock > 80 {
		res_data.Stock -= 40
	} else {
		if res_data.Stock < -80 {
			res_data.Stock += 40
		}
	}

	if res_data.Total > 80 {
		res_data.Total -= 40
	} else {
		if res_data.Total < -80 {
			res_data.Total += 40
		}
	}

	if res_data.Option > 80 {
		res_data.Option -= 40
	} else {
		if res_data.Option < -80 {
			res_data.Option += 40
		}
	}
}
