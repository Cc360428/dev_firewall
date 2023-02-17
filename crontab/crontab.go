package crontab

import (
	"github.com/robfig/cron/v3"
	"log"
	"update/update"
)

func Start() {
	go start()
}

func start() {
	//每天0点重置一些统计信息或同步数据到后台
	c := cron.New()
	spec := "0 0 * * *"
	c.AddFunc(spec, func() {
		// 每天凌晨
		log.Println("凌晨触发")
	})

	c.AddFunc("@every 30s", func() {
		log.Println("1分钟触发一次")
		err := update.GetNetName()
		if err != nil {
			log.Println("请求发生错误了", err.Error())
		}
	})

	c.Start()
	select {}
}
