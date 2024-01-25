package logic

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"server/dao/mysql"
	"server/model"
	"server/tools"
)

func History(c *gin.Context) (week_res *model.WeekRes) {
	week_reports, err := mysql.History()

	week_res = new(model.WeekRes)
	week_res.XAxis = []string{"1", "2", "3", "4", "5"}

	week_res.Stock = make([]float64, 5, 5)
	week_res.Option = make([]float64, 5, 5)
	week_res.Total_asset = make([]float64, 5, 5)
	week_res.Yield = make([]float64, 5, 5)

	for i := 0; i < 5; i++ {
		week_res.Stock[i] = 0
		week_res.Option[i] = 0
		week_res.Total_asset[i] = 0
		week_res.Yield[i] = 0
	}

	if len(week_reports) < 5 {
		zap.L().Error("Insufficient number of reports", zap.Error(err))
		return
	}

	if err != nil {
		zap.L().Error("logic.History failed", zap.Error(err))
		return
	}

	week_res.Stock[0] = week_reports[4].Stock
	week_res.Option[0] = week_reports[4].Option
	week_res.Total_asset[0] = week_res.Stock[0] + week_res.Option[0]
	week_res.Yield[0] = 0

	for i := 3; i >= 0; i-- {
		week_res.Stock[5-i-1] = week_reports[i].Stock
		week_res.Option[5-i-1] = week_reports[i].Option
		week_res.Total_asset[5-i-1] = week_res.Stock[5-i-1] + week_res.Option[5-i-1]
	}

	for i := 1; i < 5; i++ {
		week_res.Yield[i] = (week_res.Total_asset[i] - week_res.Total_asset[i-1]) * 100 / week_res.Total_asset[i-1]
		//fmt.Println(week_res.Yield[i], week_res.Total_asset[i], week_res.Total_asset[i-1])
	}

	for i := 0; i < 5; i++ {
		week_res.Stock[i] = tools.Decimal(week_res.Stock[i])
		week_res.Option[i] = tools.Decimal(week_res.Option[i])
		week_res.Total_asset[i] = tools.Decimal(week_res.Total_asset[i])
		week_res.Yield[i] = tools.Decimal(week_res.Yield[i])

	}

	return

}
