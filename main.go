package main

import "update/crontab"

/*
*

	// 目标设备地址：10.1.1.1
	// 访问方式 https://10.1.1.1 用户密码：admin/Sxf@123456
	// API接入用户：api/Sxf@123456
*/

func main() {
	//crontab.Start()
	crontab.Login()
	ch := make(chan bool)
	<-ch
}
