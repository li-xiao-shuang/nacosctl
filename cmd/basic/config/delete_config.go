package config

import (
	"github.com/spf13/cobra"
	"nacosctl/common/constant"
	"nacosctl/printer"
	"nacosctl/process/config"
)

var DeleteConfigCmd = &cobra.Command{
	Use:     "config [dataId]",
	Short:   "Delete a configuration",
	Long:    "Delete a configuration",
	Example: "nacosctl delete config [dataId]",
	Args:    cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		if args[0] == "" {
			printer.Red("[error]:dataId must not be null")
			printer.Yellow("[see]:nacosctl delete config -h")
			return
		}
		config.ParseDeleteConfigCmd(cmd, args[0], configNamespaceId, group)
	},
}

func init() {
	DeleteConfigCmd.Flags().StringVarP(&configNamespaceId, "namespaceId", "i", "", "namespace id")
	DeleteConfigCmd.Flags().StringVarP(&group, "group", "g", constant.DefaultGroup, "config group")
}
