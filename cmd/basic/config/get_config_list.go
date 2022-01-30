package config

import (
	"github.com/spf13/cobra"
	"nacosctl/common/constant"
	"nacosctl/printer"
	"nacosctl/process/config"
)

var GetConfigListCmd = &cobra.Command{
	Use:     "configs [pageNo] [pageSize]",
	Short:   "Get a list of configuration information in paginated form",
	Long:    "Get a list of configuration information in paginated form",
	Example: "nacosctl get configs [pageNo] [pageSize]",
	Args:    cobra.MinimumNArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		if args[0] == "" {
			printer.Red("[error]:pageNo must not be null")
			printer.Yellow("[see]:nacosctl get configs -h")
			return
		}
		if args[1] == "" {
			printer.Red("[error]:pageSize must not be null")
			printer.Yellow("[see]:nacosctl get configs -h")
			return
		}
		config.ParseGetConfigListCmd(cmd, args[0], args[1], configNamespaceId, group)
	},
}

func init() {
	GetConfigListCmd.Flags().StringVarP(&configNamespaceId, "namespaceId", "i", "", "namespace id")
	GetConfigListCmd.Flags().StringVarP(&group, "group", "g", constant.DefaultGroup, "config group")
}
