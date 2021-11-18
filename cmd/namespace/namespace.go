package namespace

import (
	"github.com/spf13/cobra"
)

var NamespaceCmd = &cobra.Command{
	Use:   "namespace",
	Short: "Namespace related operations",
}

func init() {
	NamespaceCmd.AddCommand(listNamespaceCmd)
	//NamespaceCmd.AddCommand(addConfigCmd)
	//NamespaceCmd.AddCommand(deleteConfigCmd)
	//NamespaceCmd.AddCommand(listConfigCmd)
	//NamespaceCmd.AddCommand(versionConfigCmd)
}
