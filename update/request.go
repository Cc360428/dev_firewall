/**
 * @Author: Cc
 * @Description: get
 * @File: get
 * @Version: 1.0.0
 * @Date: 2023/2/17 10:20
 * @Software : GoLand
 */

package update

import (
	"bytes"
	"crypto/tls"
	"log"
	"net/http"
	"time"
	"update/configs"
)

const (
	front         = "https://"
	base          = "/api/v1/namespaces/public/" // after base
	login         = "login"                      // login
	ipGroupsWanIp = "ipgroups/"                  // 网络对象获取 （get/put）

)

func Request(method string, url string, body []byte) (*http.Response, error) {

	log.Println("Request:", method, url, len(body))
	req, err := http.NewRequest(method, front+configs.Base.BaseHosts+base+url, bytes.NewBuffer(body))
	if err != nil {
		log.Println("NewRequest error:", err)
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")

	if User.Token != "" {
		req.Header.Set("Cookie", "token="+User.Token)
	}

	client := &http.Client{}
	client.Transport = &http.Transport{TLSClientConfig: &tls.Config{InsecureSkipVerify: true}}
	client.Timeout = time.Second * 15

	resp, err := client.Do(req)
	if err != nil {
		log.Println("request Do error:", err)
		return nil, err
	}

	return resp, nil
}
