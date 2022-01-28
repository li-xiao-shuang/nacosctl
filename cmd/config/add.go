package config

import (
	"errors"
	"github.com/spf13/cobra"
	"nacosctl/config"
	"nacosctl/config/constant"
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
		config.AddCommand(args[0], args[1], args[2], args[3], args[4], cmd)
	},
}

func init() {
	addConfigCmd.Flags().StringP("address", "b", constant.DefaultAddress, "nacos server ip address")
	addConfigCmd.Flags().StringP("port", "p", constant.DefaultPort, "nacos server port")
}
