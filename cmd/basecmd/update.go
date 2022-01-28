package basecmd

import (
	"github.com/spf13/cobra"
	"nacos-cli/cmd/basecmd/namespace"
)

var UpdateCmd = &cobra.Command{
	Use:     "update",
	Short:   "command to update some resources",
	Long:    "Update commands can be executed on resources such as configurations, services, namespaces, users, etc.",
	Example: "nacosctl update [resource]",
}

func init() {
	UpdateCmd.AddCommand(namespace.UpdateNamespaceCmd)
}
