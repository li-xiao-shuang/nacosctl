package command

import "github.com/spf13/cobra"

var getCmd = &cobra.Command{
	Use:     "get",
	Short:   "command to get some resources",
	Long:    "Update commands can be executed on resources such as configurations, services, namespaces, users, etc.",
	Example: "nacosctl update [resource]",
}
