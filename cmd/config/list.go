package config

import (
	"errors"
	"github.com/spf13/cobra"
	"nacos-cli/config"
	"nacos-cli/config/constant"
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
		config.ListCommand(args[0], args[1], cmd)
	},
}

func init() {
	listConfigCmd.Flags().StringP("address", "b", constant.DefaultAddress, "nacos server ip address")
	listConfigCmd.Flags().StringP("port", "p", constant.DefaultPort, "nacos server port")
}
