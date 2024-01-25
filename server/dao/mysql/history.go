package mysql

import (
	"database/sql"
	"go.uber.org/zap"
	"server/model"
)

func History() (week_reports []*model.Week, err error) {
	sqlStr := "select * from asset order by id desc"

	if err = db.Select(&week_reports, sqlStr); err != nil {
		if err == sql.ErrNoRows {
			zap.L().Warn("there is no week reports in db")
			err = nil
		}
	}
	return
}
