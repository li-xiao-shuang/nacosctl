package cmd

import (
	"errors"
	"github.com/spf13/cobra"
	"nacos-cli/config"
	"nacos-cli/config/constant"
)

var deleteConfigCmd = &cobra.Command{
	Use:   "del [namespaceId] [dataId] [group]",
	Short: "Delete a configuration",
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
		config.DelCommand(args[0], args[1], args[2], cmd)
	},
}

func init() {
	deleteConfigCmd.Flags().StringP("address", "b", constant.DefaultAddress, "nacos server ip address")
	deleteConfigCmd.Flags().StringP("port", "p", constant.DefaultPort, "nacos server port")
}
