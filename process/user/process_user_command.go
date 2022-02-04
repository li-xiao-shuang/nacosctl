package user

import (
	"encoding/json"
	"fmt"
	"github.com/logrusorgru/aurora"
	"github.com/spf13/cobra"
	"github.com/tidwall/gjson"
	"nacosctl/common/http"
	"nacosctl/printer"
	"nacosctl/process/model"
	"net/url"
)

// ParseCreateUserCmd 解析创建用户命令
func ParseCreateUserCmd(cmd *cobra.Command, username string, password string) {
	userUrl := http.GetUserUrl(cmd)
	payload := url.Values{"username": {username}, "password": {password}}
	resp := http.Post(userUrl, payload)
	if resp != "" {
		printer.Cyan("done")
		return
	}
	printer.Red("fail")
}

// ParseDeleteUserCmd 解析删除用户命令
func ParseDeleteUserCmd(cmd *cobra.Command, username string) {
	userUrl := http.GetUserUrl(cmd)
	resp := http.Delete(userUrl + "&username=" + username)
	if resp != "" {
		printer.Cyan("done")
		return
	}
	printer.Red("fail")
}

// ParseUpdateUserCmd 解析更新用户密码命令
func ParseUpdateUserCmd(cmd *cobra.Command, user string, newpassword string) {
	// todo 后续完成

}

// ParseGetUsersCmd 解析获取用户列表命令
func ParseGetUsersCmd(cmd *cobra.Command) {
	userUrl := http.GetUserUrl(cmd)
	resp := http.Get(userUrl + "&pageNo=1&pageSize=500")
	table := printer.NewTable(100)
	table.AddRow(aurora.Magenta("用户名"), aurora.Magenta("密码"))
	if resp == "" {
		fmt.Println(table)
		return
	}
	items := gjson.Get(resp, "pageItems").String()
	users := &[]model.UserInfo{}
	json.Unmarshal([]byte(items), users)

	for _, user := range *users {
		table.AddRow(user.Username, user.Password)
	}
	fmt.Println(table)
}
