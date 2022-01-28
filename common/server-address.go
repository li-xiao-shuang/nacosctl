package common

import (
	"github.com/spf13/cobra"
	"nacosctl/common/loader"
	"nacosctl/config/constant"
	"strconv"
)

func GetServerAddress(cmd *cobra.Command) (address string, port string) {
	viper := loader.GetViper()
	viperAddress := viper.Get("server.address").(string)
	viperPort := strconv.Itoa(viper.Get("server.port").(int))

	// 如果命令指定了flag  就替换配置文件
	if cmd.Flag("address").Value.String() != constant.DefaultAddress {
		viperAddress = cmd.Flag("address").Value.String()
	}
	if cmd.Flag("port").Value.String() != constant.DefaultPort {
		viperPort = cmd.Flag("port").Value.String()
	}
	return viperAddress, viperPort
}
