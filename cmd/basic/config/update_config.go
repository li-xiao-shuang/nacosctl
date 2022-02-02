package config

import (
	"github.com/spf13/cobra"
	"nacosctl/common/constant"
	"nacosctl/printer"
	"nacosctl/process/config"
)

var UpdateConfigCmd = &cobra.Command{
	Use:     "config [dataId] [content]",
	Short:   "Update configuration information",
	Long:    "Update configuration information",
	Example: "nacosctl update config [dataId] [content]",
	Args:    cobra.MinimumNArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		if args[0] == "" {
			printer.Red("[error]:dataId must not be null")
			printer.Yellow("[see]:nacosctl update config -h")
			return
		}
		if args[1] == "" {
			printer.Red("[error]:content must not be null")
			printer.Yellow("[see]:nacosctl update config -h")
			return
		}
		dataId = args[0]
		content = args[1]
		config.ParseCreateConfigCmd(cmd, configNamespaceId, dataId, group, content, configType)
	},
}

func init() {
	UpdateConfigCmd.Flags().StringVarP(&configNamespaceId, "namespaceId", "n", "", "namespace id")
	UpdateConfigCmd.Flags().StringVarP(&group, "group", "g", constant.DefaultGroup, "config group")
	UpdateConfigCmd.Flags().StringVarP(&configType, "type", "t", constant.DefaultType, "config type")

}
