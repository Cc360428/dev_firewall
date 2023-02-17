/**
 * @Author: Cc
 * @Description: login
 * @File: login
 * @Version: 1.0.0
 * @Date: 2023/2/17 14:16
 * @Software : GoLand
 */

package update

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"update/mod"
)

func Login() error {

	var marshal, err = json.Marshal(&mod.LoginRequest{
		Name:     Name,
		Password: Password,
	})
	if err != nil {
		return err
	}

	response, err := Request(http.MethodPost, login, marshal)
	all, err := io.ReadAll(response.Body)
	if err != nil {
		return err
	}

	loginResponse := new(mod.LoginResponse)
	err = json.Unmarshal(all, loginResponse)
	if err != nil {
		return err
	}

	if loginResponse.Code != 0 {
		return err
	}
	User.Name = loginResponse.Data.Name
	User.Token = loginResponse.Data.LoginResult.Token
	User.Namespace = loginResponse.Data.Namespace
	User.Role = loginResponse.Data.Role
	log.Println("login 请求最后回应：", loginResponse)
	return nil
}
