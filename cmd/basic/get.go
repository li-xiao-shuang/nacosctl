package basic

import (
	"github.com/spf13/cobra"
	"nacosctl/cmd/basic/config"
	"nacosctl/cmd/basic/namespace"
	"nacosctl/cmd/basic/user"
)

var GetCmd = &cobra.Command{
	Use:     "get",
	Short:   "command to get some resources",
	Long:    "Update commands can be executed on resources such as configurations, services, namespaces, users, etc.",
	Example: "nacosctl get [resource]",
}

func init() {
	GetCmd.AddCommand(namespace.GetNamespaceListCmd)
	GetCmd.AddCommand(namespace.GetNamespaceCmd)
	GetCmd.AddCommand(config.GetConfigCmd)
	GetCmd.AddCommand(config.GetConfigListCmd)
	GetCmd.AddCommand(user.GetUserCmd)
}
