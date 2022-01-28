package namespace

import (
	"github.com/spf13/cobra"
	"nacosctl/printer"
	"nacosctl/process/namespace"
)

var UpdateNamespaceCmd = &cobra.Command{
	Use:     "namespace [namespaceName] [namespaceDesc] [namespaceId] ",
	Short:   "Update the name and description of the specified namespace",
	Long:    "Update the name and description of the specified namespace",
	Example: "nacosctl update namespace [namespaceName] [namespaceDesc] [namespaceId]",
	Args:    cobra.MinimumNArgs(3),
	Run: func(cmd *cobra.Command, args []string) {
		namespaceName := args[0]
		namespaceDesc := args[1]
		namespaceId := args[2]
		if namespaceName == "" {
			printer.Red("[error]:namespaceName must not be null")
			printer.Yellow("[see]:nacosctl update namespace -h")
			return
		}
		if namespaceDesc == "" {
			printer.Red("[error]:namespaceDesc must not be null")
			printer.Yellow("[see]:nacosctl update namespace -h")
			return
		}
		if namespaceId == "" {
			printer.Red("[error]:namespaceId must not be null")
			printer.Yellow("[see]:nacosctl update namespace -h")
			return
		}
		namespace.ParseUpdateNamespaceCmd(cmd, namespaceName, namespaceDesc, namespaceId)
	},
}
