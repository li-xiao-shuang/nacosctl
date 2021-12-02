package instance

import (
	"github.com/spf13/cobra"
)

var InstanceCmd = &cobra.Command{
	Use:   "instance",
	Short: "Service instance operation command",
}

func init() {
	//NamespaceCmd.AddCommand(listNamespaceCmd)
	//NamespaceCmd.AddCommand(addNamespaceCmd)
	//NamespaceCmd.AddCommand(delNamespace)
	//NamespaceCmd.AddCommand(updateNamespaceCmd)
	////NamespaceCmd.AddCommand(versionConfigCmd)
}
