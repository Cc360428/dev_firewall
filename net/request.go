/**
 * @Author: Cc
 * @Description: get
 * @File: get
 * @Version: 1.0.0
 * @Date: 2023/2/17 10:20
 * @Software : GoLand
 */

package net

import (
	"bytes"
	"crypto/tls"
	"log"
	"net/http"
	"time"
	"update/consts"
)

func Request(method string, url string, body []byte) *http.Response {

	req, err := http.NewRequest(method, consts.BaseURL+url, bytes.NewBuffer(body))
	if err != nil {
		log.Println("NewRequest error:", err)
		return nil
	}

	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	client.Transport = &http.Transport{TLSClientConfig: &tls.Config{InsecureSkipVerify: true}}
	client.Timeout = time.Second * 30

	resp, err := client.Do(req)
	if err != nil {
		log.Println("request Do error:", err)
		return nil
	}

	return resp
}
