package crontab

import (
	"encoding/json"
	"github.com/robfig/cron/v3"
	"io"
	"log"
	"net/http"
	"update/consts"
	"update/mod"
	"update/net"
)

var User UserStruct

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
		Login()
	})

	c.Start()
	select {}
}

type UserStruct struct {
	Name      string
	Token     string
	Role      string
	Namespace string
}

func Login() {

	var marshal, err = json.Marshal(&mod.LoginRequest{
		Name:     "api",
		Password: "Sxf@123456",
	})

	if err != nil {
		return
	}
	response := net.Request(http.MethodPost, consts.Login, marshal)
	all, err := io.ReadAll(response.Body)
	if err != nil {
		return
	}

	loginResponse := new(mod.LoginResponse)
	err = json.Unmarshal(all, loginResponse)

	if err != nil {
		return
	}

	if loginResponse.Code != 0 {
		return
	}
	User.Name = loginResponse.Data.Name
	User.Token = loginResponse.Data.LoginResult.Token
	User.Namespace = loginResponse.Data.Namespace
	User.Role = loginResponse.Data.Role
	log.Println("请求最后回应：", loginResponse)
}
