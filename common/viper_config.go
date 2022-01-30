package common

import (
	"github.com/spf13/cobra"
	"nacosctl/common/loader"
	"nacosctl/process/constant"
)

func GetServerAddress(cmd *cobra.Command) (address string) {
	viper := loader.GetViper()
	viperAddress := viper.Get("nacosctl.server.address").(string)

	// 如果命令指定了flag  就替换配置文件
	if cmd.Flag("address").Value.String() != constant.DefaultAddress {
		viperAddress = cmd.Flag("address").Value.String()
	}
	return viperAddress
}
