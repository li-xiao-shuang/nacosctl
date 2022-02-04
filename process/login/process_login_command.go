package login

import (
	"encoding/json"
	"github.com/spf13/viper"
	"nacosctl/common/http"
	"nacosctl/printer"
	"nacosctl/process/model"
	"net/url"
)

// ParseLoginCmd 解析login命令
func ParseLoginCmd(username string, password string) {
	loginUrl := http.GetLoginUrl()
	payload := url.Values{"username": {username}, "password": {password}}
	resp := http.Post(loginUrl, payload)
	if resp == "" {
		printer.Red("wrong user name or password")
		return
	}
	login := &model.LoginInfo{}
	err := json.Unmarshal([]byte(resp), login)
	if err != nil {
		printer.Yellow(err)
	}
	viper.Set("nacosctl.username", username)
	viper.Set("nacosctl.password", password)
	viper.Set("nacosctl.accessToken", login.AccessToken)
	viper.WriteConfigAs("nacosctl.yaml")
	printer.Cyan("login successful")
}
