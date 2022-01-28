package namespace

import (
	"fmt"
	"github.com/spf13/cobra"
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
			fmt.Println("[error]:namespaceId must not be null")
			fmt.Println("[see]:nacosctl get namespace -h")
			return
		}
		namespace.ParseGetNamespaceCmd(cmd, namespaceId)
	},
}
