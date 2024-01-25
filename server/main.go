package main

import (
	"fmt"
	"os"
	"server/dao/mysql"
	"server/logger"
	"server/routes"
	"server/settings"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Printf("Insufficient parameters. Need path to config file")
		return
	}

	if err := settings.Init(os.Args[1]); err != nil {
		fmt.Printf("Load config file failed, err:%v\n", err)
		return
	}

	if err := logger.Init(settings.Conf.LogConfig, settings.Conf.Mode); err != nil {
		fmt.Printf("Init logger failed, err:%v\n", err)
		return
	}

	if err := mysql.Init(settings.Conf.MySQLConfig); err != nil {
		fmt.Printf("init mysql failed, err:%v\n", err)
		return
	}
	defer mysql.Close() // 程序退出关闭数据库连接

	r := routes.SetUp(settings.Conf.Mode)
	err := r.Run(fmt.Sprintf(":%d", settings.Conf.Port))
	if err != nil {
		fmt.Printf("run server failed, err:%v\n", err)
		return
	}

}
