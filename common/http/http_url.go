package http

import (
	"fmt"
	"github.com/spf13/cobra"
	"nacosctl/common"
	"nacosctl/common/constant"
)

func GetNamespaceUrl(cmd *cobra.Command) string {
	return getServerAddressByViper(cmd) + constant.NamespaceUrl
}

func GetConfig(cmd *cobra.Command) string {
	return getServerAddressByViper(cmd) + constant.ConfigUrl
}

func GetConfigVersion(cmd *cobra.Command) string {
	return getServerAddressByViper(cmd) + constant.VersionUrl
}

func GetLoginUrl(cmd *cobra.Command) string {
	return getServerAddressByViper(cmd) + constant.LoginUrl
}

func getServerAddressByViper(cmd *cobra.Command) string {
	address := common.GetServerAddress(cmd)
	prefix := fmt.Sprintf(constant.Prefix, address)
	return prefix
}
