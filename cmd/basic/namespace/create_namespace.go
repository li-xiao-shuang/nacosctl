package namespace

import (
	"github.com/spf13/cobra"
	"nacosctl/printer"
	"nacosctl/process/namespace"
)

var defaultNamespaceId string

var CreateNamespaceCmd = &cobra.Command{
	Use:     "namespace [namespaceName] [namespaceDesc]",
	Short:   "Create a namespace, If namespaceId is not specified, it will be automatically generated",
	Long:    "Create a namespace, If namespaceId is not specified, it will be automatically generated",
	Example: "nacosctl create namespace [namespaceName] [namespaceDesc] [flags]",
	Args:    cobra.MinimumNArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		namespaceName := args[0]
		namespaceDesc := args[1]
		if namespaceName == "" {
			printer.Red("[error]:namespaceName must not be null")
			printer.Yellow("[see]:nacosctl create namespace -h")
			return
		}
		if namespaceDesc == "" {
			printer.Red("[error]:namespaceDesc must not be null")
			printer.Yellow("[see]:nacosctl create namespace -h")
			return
		}
		namespace.ParseCreateNamespaceCmd(cmd, namespaceName, namespaceDesc, defaultNamespaceId)
	},
}

func init() {
	CreateNamespaceCmd.Flags().StringVarP(&defaultNamespaceId, "namespaceId", "i", "", "namespace id")
}
