package http

import (
	"fmt"
	"github.com/spf13/viper"
	"nacosctl/common"
	"nacosctl/common/constant"
)

func GetNamespaceUrl() string {
	return getServerAddressByViper() + constant.NamespaceUrl + "?accessToken=" + viper.GetString("nacosctl.accessToken")
}

func GetConfigUrl() string {
	return getServerAddressByViper() + constant.ConfigUrl + "?accessToken=" + viper.GetString("nacosctl.accessToken")
}

func GetConfigVersionUrl() string {
	return getServerAddressByViper() + constant.VersionUrl + "?accessToken=" + viper.GetString("nacosctl.accessToken")
}

func GetUserUrl() string {
	return getServerAddressByViper() + constant.UserUrl + "?accessToken=" + viper.GetString("nacosctl.accessToken")
}

func GetLoginUrl() string {
	return getServerAddressByViper() + constant.LoginUrl
}

func getServerAddressByViper() string {
	address := common.GetServerAddress()
	prefix := fmt.Sprintf(constant.Prefix, address)
	return prefix
}
