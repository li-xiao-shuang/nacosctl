package config

import (
	"github.com/spf13/cobra"
	"nacosctl/common/constant"
	"nacosctl/printer"
	"nacosctl/process/config"
)

var (
	configNamespaceId string
	dataId            string
	group             string
	content           string
	configType        string
)

var CreateConfigCmd = &cobra.Command{
	Use:     "config [dataId] [content]",
	Short:   "Create a configuration",
	Long:    "Create a configuration",
	Example: "nacosctl create config [dataId] [content]",
	Args:    cobra.MinimumNArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		if args[0] == "" {
			printer.Red("[error]:dataId must not be null")
			printer.Yellow("[see]:nacosctl create config -h")
			return
		}
		if args[1] == "" {
			printer.Red("[error]:content must not be null")
			printer.Yellow("[see]:nacosctl create config -h")
			return
		}
		dataId = args[0]
		content = args[1]
		config.ParseCreateConfigCmd(cmd, configNamespaceId, dataId, group, content, configType)
	},
}

func init() {
	CreateConfigCmd.Flags().StringVarP(&configNamespaceId, "namespaceId", "i", "", "namespace id")
	CreateConfigCmd.Flags().StringVarP(&group, "group", "g", constant.DefaultGroup, "config group")
	CreateConfigCmd.Flags().StringVarP(&configType, "type", "t", constant.DefaultType, "config type")

}
