/**
 * @Author: Cc
 * @Description: login
 * @File: login
 * @Version: 1.0.0
 * @Date: 2023/2/17 11:35
 * @Software : GoLand
 */

package mod

type LoginRequest struct {
	Name     string `json:"name"`
	Password string `json:"password"`
}

type LoginResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    Data   `json:"data"`
}

type LoginResult struct {
	Token string `json:"token"`
}

type Data struct {
	Name         string      `json:"name"`
	LoginResult  LoginResult `json:"loginResult"`
	PasswdStatus bool        `json:"passwdStatus"`
	Role         string      `json:"role"`
	Namespace    string      `json:"namespace"`
}
