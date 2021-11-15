package cmd

import (
	"errors"
	"github.com/spf13/cobra"
	"nacos-cli/config"
	"strconv"
)

var getConfigCmd = &cobra.Command{
	Use:   "get [dataId] [group]",
	Short: "Get configuration information",
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 2 {
			return errors.New("Two parameters are required,[dataId] and [group]")
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		viper := config.GetViper()
		address := viper.Get("server.address").(string)
		port := strconv.Itoa(viper.Get("server.port").(int))

		// 如果命令指定了flag  就替换配置文件
		if cmd.Flag("address").Value.String() != "127.0.0.1" {
			address = cmd.Flag("address").Value.String()
		}
		if cmd.Flag("port").Value.String() != "8848" {
			port = cmd.Flag("port").Value.String()
		}
		config.GetCommand(args[0], args[1], address, port)
	},
}

func init() {
	getConfigCmd.Flags().StringP("address", "b", "127.0.0.1", "nacos server ip address")
	getConfigCmd.Flags().StringP("port", "p", "8848", "nacos server port")
	configCmd.AddCommand(getConfigCmd)
}
