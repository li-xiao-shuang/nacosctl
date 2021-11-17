package cmd

import (
	"errors"
	"github.com/spf13/cobra"
	"nacos-cli/config"
	"nacos-cli/config/constant"
)

var getConfigCmd = &cobra.Command{
	Use:   "get [namespaceId] [dataId] [group]",
	Short: "Get configuration information",
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 3 {
			return errors.New("3 parameters are required,[namespaceId] [dataId] [group]")
		}
		if args[2] == "" {
			args[2] = constant.DefaultGroup
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		config.GetCommand(args[0], args[1], args[2], cmd)
	},
}

func init() {
	getConfigCmd.Flags().StringP("address", "b", constant.DefaultAddress, "nacos server ip address")
	getConfigCmd.Flags().StringP("port", "p", constant.DefaultPort, "nacos server port")
}
