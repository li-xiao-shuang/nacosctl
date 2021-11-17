package cmd

import (
	"errors"
	"github.com/spf13/cobra"
	"nacos-cli/common/loader"
	"nacos-cli/config"
	"nacos-cli/config/constant"
	"strconv"
)

var addConfigCmd = &cobra.Command{
	Use:   "add [namespaceId] [dataId] [group] [content] [type]",
	Short: "Add a configuration",
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 5 {
			return errors.New("5 parameters are required,[namespaceId] [dataId] [group] [content] [type]")
		}
		if args[2] == "" {
			args[2] = constant.DefaultGroup
		}
		if args[4] == "" {
			args[4] = constant.DefaultType
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		viper := loader.GetViper()
		address := viper.Get("server.address").(string)
		port := strconv.Itoa(viper.Get("server.port").(int))

		// 如果命令指定了flag  就替换配置文件
		if cmd.Flag("address").Value.String() != constant.DefaultAddress {
			address = cmd.Flag("address").Value.String()
		}
		if cmd.Flag("port").Value.String() != constant.DefaultPort {
			port = cmd.Flag("port").Value.String()
		}
		config.AddCommand(args[0], args[1], args[2], args[3], args[4], address, port)
	},
}

func init() {
	addConfigCmd.Flags().StringP("address", "b", constant.DefaultAddress, "nacos server ip address")
	addConfigCmd.Flags().StringP("port", "p", constant.DefaultPort, "nacos server port")
	configCmd.AddCommand(addConfigCmd)
}
