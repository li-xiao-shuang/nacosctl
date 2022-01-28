package config

import (
	"errors"
	"github.com/spf13/cobra"
	"nacosctl/process/config"
	"nacosctl/process/constant"
)

var CreateConfigCmd = &cobra.Command{
	Use:     "config [namespaceId] [dataId] [group] [content] [type]",
	Short:   "Create a configuration",
	Long:    "Create a configuration",
	Example: "nacosctl create config [namespaceId] [dataId] [group] [content] [type]",
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
		config.ParseCreateConfigCmd(args[0], args[1], args[2], args[3], args[4], cmd)
	},
}
