package command

import "github.com/spf13/cobra"

var updateCmd = &cobra.Command{
	Use:     "update",
	Short:   "command to update some resources",
	Long:    "Update commands can be executed on resources such as configurations, services, namespaces, users, etc.",
	Example: "nacosctl update [resource]",
}
