package main

import (
	"log"
	"os"
	"update/configs"
	"update/crontab"
	"update/update"

	"github.com/Cc360428/HelpPackage/monitor_message"
)

func init() {
	configs.InitConfigStart()
}

func main() {
	monitor_message.NewDingDing(configs.Base.SendMessage)
	log.Println("base:", configs.Base)

	err := update.Login()
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	crontab.Start()
	monitor_message.Send("info", "PublicIp: dev_firewall start", true)
	ch := make(chan bool)
	<-ch
}
