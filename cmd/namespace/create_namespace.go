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
	example       = "nacosctl namespace create [--namespaceId=] --namespaceName=test --namespaceDesc=test"
)

var createNamespaceCmd = &cobra.Command{
	Use:     "create",
	Short:   "Create a namespace, If namespaceid is not specified, it will be automatically generated",
	Long:    "Create a namespace, If namespaceid is not specified, it will be automatically generated",
	Example: example,
	Run: func(cmd *cobra.Command, args []string) {
		if namespaceName == "" || namespaceDesc == "" {
			fmt.Printf("reference example ：%s\n", example)
			return
		}
		namespace.AddCommand(cmd, namespaceId, namespaceName, namespaceDesc)
	},
}

func init() {
	createNamespaceCmd.Flags().StringVarP(&namespaceId, "namespaceId", "i", "", "namespace id")
	createNamespaceCmd.Flags().StringVarP(&namespaceName, "namespaceName", "n", "", "namespace name")
	createNamespaceCmd.Flags().StringVarP(&namespaceDesc, "namespaceDesc", "d", "", "namespace desc")
}
