package namespace

import (
	"github.com/spf13/cobra"
	"nacosctl/printer"
	"nacosctl/process/namespace"
)

var DeleteNamespaceCmd = &cobra.Command{
	Use:     "namespace [namespaceId]",
	Short:   "delete the specified namespace",
	Long:    "delete the specified namespace",
	Example: "nacosctl delete namespace [namespaceId]",
	Args:    cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		namespaceId := args[0]
		if namespaceId == "" {
			printer.Red("[error]:namespaceId must not be null")
			printer.Yellow("[see]:nacosctl delete namespace -h")
			return
		}
		namespace.ParseDeleteNamespaceCmd(cmd, namespaceId)
	},
}
