package main

import (
	"flag"
	"log"
	"os"
	"update/crontab"
	"update/update"
)

var hosts string
var name string
var password string
var netName string

func init() {
	flag.StringVar(&hosts, "host", "10.1.1.1", "服务地址")
	flag.StringVar(&name, "name", "api", "用户名")
	flag.StringVar(&password, "password", "Sxf@123456", "密码")
	flag.StringVar(&netName, "netName", "wan_ip", "网络对象名字")
	flag.Parse()
}

func main() {
	log.Println(hosts, name, password, netName)

	update.Hosts = hosts
	update.Name = name
	update.Password = password
	update.NetName = netName

	err := update.Login()
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	crontab.Start()
	ch := make(chan bool)
	<-ch
}
