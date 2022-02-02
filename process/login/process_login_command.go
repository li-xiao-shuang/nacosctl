package login

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"nacosctl/common/http"
	"nacosctl/printer"
	"net/url"
)

// ParseLoginCmd 解析login命令
func ParseLoginCmd(cmd *cobra.Command, username string, password string) {
	loginUrl := http.GetLoginUrl(cmd)
	payload := url.Values{"username": {username}, "password": {password}}
	resp := http.Post(loginUrl, payload)
	if resp == "" {
		printer.Red("wrong user name or password")
		return
	}
	viper.Set("nacosctl.user", username)
	viper.Set("nacosctl.password", password)
	viper.WriteConfigAs("nacosctl.yaml")
	printer.Cyan("login successful")
}
