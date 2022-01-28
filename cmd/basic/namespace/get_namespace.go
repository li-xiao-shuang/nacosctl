package namespace

import (
	"github.com/spf13/cobra"
	"nacosctl/printer"
	"nacosctl/process/namespace"
)

var GetNamespaceCmd = &cobra.Command{
	Use:     "namespace [namespaceId]",
	Short:   "Get the specified namespace information",
	Long:    "Get the specified namespace information",
	Example: "nacosctl get namespace [namespaceId]",
	Args:    cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		namespaceId := args[0]
		if namespaceId == "" {
			printer.Red("[error]:namespaceId must not be null")
			printer.Yellow("[see]:nacosctl get namespace -h")
			return
		}
		namespace.ParseGetNamespaceCmd(cmd, namespaceId)
	},
}
