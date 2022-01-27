package namespace

import (
	"github.com/spf13/cobra"
	"nacos-cli/config/constant"
	"nacos-cli/namespace"
)

var listNamespaceCmd = &cobra.Command{
	Use:   "list",
	Short: "Namespace list",
	Run: func(cmd *cobra.Command, args []string) {
		namespace.ListCommand(cmd)
	},
}

func init() {
	listNamespaceCmd.Flags().StringP("address", "b", constant.DefaultAddress, "nacos server ip address")
	listNamespaceCmd.Flags().StringP("port", "p", constant.DefaultPort, "nacos server port")
}
