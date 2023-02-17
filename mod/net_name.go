/**
 * @Author: Cc
 * @Description: 描述
 * @File: wan_ip
 * @Version: 1.0.0
 * @Date: 2023/2/17 14:57
 * @Software : GoLand
 */

package mod

type IGroup struct {
	Code    int        `json:"code"`
	Message string     `json:"message"`
	Data    IGroupData `json:"data"`
}
type IPRanges struct {
	Start string `json:"start"`
}
type IGroupData struct {
	UUID         string     `json:"uuid"`
	Name         string     `json:"name"`
	ID           int        `json:"id"`
	BusinessType string     `json:"businessType"`
	Description  string     `json:"description"`
	AddressType  string     `json:"addressType"`
	IPRanges     []IPRanges `json:"ipRanges"`
	Important    string     `json:"important"`
	Hasref       bool       `json:"hasref"`
}

type PutIGroup struct {
	BusinessType string     `json:"businessType"`
	Name         string     `json:"name"`
	IPRanges     []IPRanges `json:"ipRanges"`
}
