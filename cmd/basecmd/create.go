package basecmd

import (
	"github.com/spf13/cobra"
	"nacosctl/cmd/basecmd/config"
	"nacosctl/cmd/basecmd/namespace"
)

var CreateCmd = &cobra.Command{
	Use:     "create",
	Short:   "command to create some resources",
	Long:    "Create commands can be executed on resources such as configurations, services, namespaces, users, etc.",
	Example: "nacosctl create [resource]",
}

func init() {
	CreateCmd.AddCommand(namespace.CreateNamespaceCmd)
	CreateCmd.AddCommand(config.CreateConfigCmd)
}
