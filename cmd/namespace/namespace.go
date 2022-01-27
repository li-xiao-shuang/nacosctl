package namespace

import (
	"github.com/spf13/cobra"
)

var NamespaceCmd = &cobra.Command{
	Use:       "namespace",
	Short:     "Namespace related operations",
	ValidArgs: make([]string, 1),
}

func init() {
	NamespaceCmd.AddCommand(createNamespaceCmd)
	NamespaceCmd.AddCommand(updateNamespaceCmd)
	NamespaceCmd.AddCommand(deleteNamespace)
	NamespaceCmd.AddCommand(listNamespaceCmd)
}
