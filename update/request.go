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
)

var Hosts string

const (
	front         = "https://"
	base          = "/api/v1/namespaces/public/" // after base
	login         = "login"                      // login
	ipGroupsWanIp = "ipgroups/"                  // 网络对象获取 （get/put）

)

/**

获取公网IP地址：

业务操作：
1.用户登录
 POST https://10.1.1.1/api/v1/namespaces/public/login
 header字段==》Connection：keep-alive Content-Type：application/json
 body字段====》
  {
 "name": "api",
 "password": "Sxf@123456"
}
返回值：token
2.网络对象获取
 GET https://10.1.1.1/api/v1/namespaces/public/ipgroups/wan_ip
 header字段==》cookie:token=上一步的token
返回值："message": "成功"
3.全量修改该对象
 PUT  https://10.1.1.1/api/v1/namespaces/public/ipgroups/wan_ip
 header字段==》cookie:token=上一步的token
 body字段====》
 {
    "businessType": "IP",
 "name": "wan_ip",
 "ipRanges": [
  {
   "start": "61.189.121.42"
  }
 ]
}
返回值： "message": "成功"
4.用户注销 (可不做)
*/

func Request(method string, url string, body []byte) (*http.Response, error) {

	log.Println("Request:", method, url, len(body))
	req, err := http.NewRequest(method, front+Hosts+base+url, bytes.NewBuffer(body))
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
