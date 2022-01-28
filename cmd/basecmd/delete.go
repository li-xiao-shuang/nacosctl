package basecmd

import (
	"github.com/spf13/cobra"
	"nacosctl/cmd/basecmd/namespace"
)

var DeleteCmd = &cobra.Command{
	Use:     "delete",
	Short:   "command to delete some resources",
	Long:    "Delete commands can be executed on resources such as configurations, services, namespaces, users, etc.",
	Example: "nacosctl delete [resource]",
}

func init() {
	DeleteCmd.AddCommand(namespace.DeleteNamespaceCmd)
}
