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
