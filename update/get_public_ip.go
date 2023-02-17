/**
 * @Author: Cc
 * @Description: 描述
 * @File: get_public_ip
 * @Version: 1.0.0
 * @Date: 2023/2/17 15:24
 * @Software : GoLand
 */

package update

import (
	"encoding/json"
	"io"
	"net/http"
)

type IPInfo struct {
	Origin string `json:"origin"`
}

func GetPublicIp() string {
	// 向 httpbin.org 发送 GET 请求，以获取客户端的公网 IP 地址
	resp, err := http.Get("http://httpbin.org/ip")
	if err != nil {
		// 处理错误
	}

	defer resp.Body.Close()

	// 读取响应内容
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		// 处理错误
	}

	// 解析响应内容中的 IP 地址
	var ipInfo IPInfo
	err = json.Unmarshal(body, &ipInfo)
	if err != nil {
		// 处理错误
		return ""
	}

	return ipInfo.Origin
}
