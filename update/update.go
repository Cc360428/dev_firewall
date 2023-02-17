/**
 * @Author: Cc
 * @Description: update
 * @File: update
 * @Version: 1.0.0
 * @Date: 2023/2/17 14:16
 * @Software : GoLand
 */

package update

type UserStruct struct {
	Name      string
	Token     string
	Role      string
	Namespace string
}

var (
	User     UserStruct
	Name     string
	Password string
	NetName  string
)
