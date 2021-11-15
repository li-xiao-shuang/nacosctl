package cmd

import (
	"errors"
	"github.com/spf13/cobra"
	"nacos-cli/config"
	"nacos-cli/config/constant"
	"strconv"
)

var listConfigCmd = &cobra.Command{
	Use:   "list [pageNo] [pageSize]",
	Short: "Configuration information list",
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 2 {
			return errors.New("2 parameters are required,[pageNo] [pageSize]")
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
		config.ListCommand(args[0], args[1], address, port)
	},
}

func init() {
	listConfigCmd.Flags().StringP("address", "b", constant.DefaultAddress, "nacos server ip address")
	listConfigCmd.Flags().StringP("port", "p", constant.DefaultPort, "nacos server port")
	configCmd.AddCommand(listConfigCmd)
}
