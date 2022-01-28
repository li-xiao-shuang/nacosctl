package namespace

import (
	"fmt"
	"github.com/spf13/cobra"
	"nacos-cli/namespace"
)

var (
	namespaceId   = ""
	namespaceName = ""
	namespaceDesc = ""
	example       = "nacosctl create namespace [namespaceName] [namespaceDesc] [namespaceId] "
)

var CreateNamespaceCmd = &cobra.Command{
	Use:     "namespace [namespaceName] [namespaceDesc] [namespaceId]",
	Short:   "Create a namespace, If namespaceid is not specified, it will be automatically generated",
	Long:    "Create a namespace, If namespaceid is not specified, it will be automatically generated",
	Example: example,
	Args:    cobra.MinimumNArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		namespaceName = args[0]
		namespaceDesc = args[1]
		if namespaceName == "" {
			fmt.Println("[error]:namespaceName must not be null")
			fmt.Println("[see]:nacosctl create namespace -h")
			return
		}
		if namespaceDesc == "" {
			fmt.Println("[error]:namespaceDesc must not be null")
			fmt.Println("[see]:nacosctl create namespace -h")
			return
		}
		namespace.AddCommand(cmd, namespaceId, namespaceName, namespaceDesc)
	},
}

func init() {

}
