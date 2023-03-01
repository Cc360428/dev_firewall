/**
 * @Author: Cc
 * @Description: configs
 * @File: base.go
 * @Version: 1.0.0
 * @Date: 2023/3/1 17:29
 * @Software : GoLand
 */

package configs

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/BurntSushi/toml"
)

type BaseConfig struct {
	BaseHosts   string `toml:"BaseHosts"`
	Name        string `toml:"Name"`
	Password    string `toml:"Password"`
	NetName     string `toml:"NetName"`
	SendMessage string `toml:"SendMessage"`
}

var Base BaseConfig

func getDir() string {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		fmt.Println("Get path config error ", err.Error())
		panic(err)
	}
	return dir
}

func InitConfigStart() {
	_, err := toml.DecodeFile(fmt.Sprintf("%s/conf/base.toml", strings.Replace(getDir(), "\\", "/", -1)), &Base)
	if err != nil {
		fmt.Println("init config error ", err.Error())
		panic(err)
	}
}
