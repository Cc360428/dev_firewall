package crontab

import (
	"github.com/robfig/cron/v3"
	"log"
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

	c.AddFunc("@every 1m", func() {
		log.Println("1分钟触发一次")

	})

	c.Start()
	select {}
}
