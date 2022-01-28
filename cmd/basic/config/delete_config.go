package config

import (
	"errors"
	"github.com/spf13/cobra"
	"nacosctl/process/config"
	"nacosctl/process/constant"
)

var DeleteConfigCmd = &cobra.Command{
	Use:     "config [namespaceId] [dataId] [group]",
	Short:   "Delete a configuration",
	Long:    "Delete a configuration",
	Example: "nacosctl delete config [namespaceId] [dataId] [group]",
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
		config.ParseDeleteConfigCmd(args[0], args[1], args[2], cmd)
	},
}
