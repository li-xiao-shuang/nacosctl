package cmd

import (
	"errors"
	"github.com/spf13/cobra"
	"nacos-cli/config"
	"nacos-cli/config/constant"
	"strconv"
)

var versionConfigCmd = &cobra.Command{
	Use:   "version [id] [namespaceId] [dataId] [group]",
	Short: "Query the configuration information of the previous version",
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 4 {
			return errors.New("4 parameters are required,[id] [namespaceId] [dataId] [group]")
		}
		if args[3] == "" {
			args[3] = constant.DefaultGroup
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		viper := config.GetViper()
		address := viper.Get("server.address").(string)
		port := strconv.Itoa(viper.Get("server.port").(int))

		// 如果命令指定了flag  就替换配置文件
		if cmd.Flag("address").Value.String() != constant.DefaultAddress {
			address = cmd.Flag("address").Value.String()
		}
		if cmd.Flag("port").Value.String() != constant.DefaultPort {
			port = cmd.Flag("port").Value.String()
		}
		config.VersionCommand(args[0], args[1], args[2], args[3], address, port)
	},
}

func init() {
	versionConfigCmd.Flags().StringP("address", "b", constant.DefaultAddress, "nacos server ip address")
	versionConfigCmd.Flags().StringP("port", "p", constant.DefaultPort, "nacos server port")
	configCmd.AddCommand(versionConfigCmd)
}
