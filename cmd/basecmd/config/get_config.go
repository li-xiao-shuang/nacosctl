package config

import (
	"errors"
	"github.com/spf13/cobra"
	"nacosctl/config/constant"
	"nacosctl/process/config"
)

var GetConfigCmd = &cobra.Command{
	Use:     "config [namespaceId] [dataId] [group]",
	Short:   "Get configuration information",
	Long:    "Get configuration information",
	Example: "nacosctl get config [namespaceId] [dataId] [group]",
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
		config.ParseGetConfigCmd(args[0], args[1], args[2], cmd)
	},
}
