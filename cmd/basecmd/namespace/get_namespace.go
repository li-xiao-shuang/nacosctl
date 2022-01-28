package namespace

import (
	"fmt"
	"github.com/spf13/cobra"
	"nacosctl/process/namespace"
)

var getNamespaceCmd = &cobra.Command{
	Use:     "namespace [namespaceId]",
	Short:   "delete the specified namespace",
	Long:    "delete the specified namespace",
	Example: "nacosctl delete namespace [namespaceId]",
	Args:    cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		namespaceId := args[0]
		if namespaceId == "" {
			fmt.Println("[error]:namespaceId must not be null")
			fmt.Println("[see]:nacosctl delete namespace -h")
			return
		}
		namespace.ParseDeleteNamespaceCmd(cmd, namespaceId)
	},
}
