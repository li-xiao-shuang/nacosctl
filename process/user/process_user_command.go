package user

import (
	"github.com/spf13/cobra"
	"nacosctl/common/http"
	"nacosctl/printer"
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
