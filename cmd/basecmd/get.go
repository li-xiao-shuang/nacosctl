package basecmd

import (
	"github.com/spf13/cobra"
	"nacosctl/cmd/basecmd/namespace"
)

var GetCmd = &cobra.Command{
	Use:     "get",
	Short:   "command to get some resources",
	Long:    "Update commands can be executed on resources such as configurations, services, namespaces, users, etc.",
	Example: "nacosctl update [resource]",
}

func init()  {
	GetCmd.AddCommand(namespace.ListNamespaceCmd)
}
