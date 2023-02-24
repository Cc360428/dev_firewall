package main

import (
	"flag"
	"log"
	"os"
	"update/crontab"
	"update/update"

	"github.com/Cc360428/HelpPackage/monitor_message"
)

var hosts string
var name string
var password string
var netName string
var ddToken string

func init() {
	flag.StringVar(&hosts, "host", "10.1.1.1", "服务地址")
	flag.StringVar(&name, "name", "api", "用户名")
	flag.StringVar(&password, "password", "Sxf@123456", "密码")
	flag.StringVar(&netName, "netName", "wan_ip", "网络对象名字")
	flag.StringVar(&ddToken, "dd", "9aec70a47db2a6bbea12682f113e15d86b66fae93bd9f5d70391557369a55798", "dingding toekn")

	flag.Parse()
}

func main() {
	monitor_message.NewDingDing(ddToken)
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
	monitor_message.Send("info", "PublicIp: dev_firewall start", true)
	ch := make(chan bool)
	<-ch
}
