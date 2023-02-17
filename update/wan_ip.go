/**
 * @Author: Cc
 * @Description: wan_ip
 * @File: wan_ip
 * @Version: 1.0.0
 * @Date: 2023/2/17 14:46
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

func GetNetName() error {
	response, err := Request(http.MethodGet, ipGroupsWanIp+NetName, nil)
	if err != nil {
		return err
	}

	all, err := io.ReadAll(response.Body)
	if err != nil {
		return err
	}

	inResponse := new(mod.IGroup)
	err = json.Unmarshal(all, inResponse)
	if err != nil {
		return err
	}

	if inResponse.Code != 0 {
		return err
	}

	for k, v := range inResponse.Data.IPRanges {
		log.Println(k, v.Start)
		log.Println(v.Start, "这里需要对比")
	}

	return nil
}
