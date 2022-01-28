package command

import "github.com/spf13/cobra"

var deleteCmd = &cobra.Command{
	Use:     "delete",
	Short:   "command to delete some resources",
	Long:    "Delete commands can be executed on resources such as configurations, services, namespaces, users, etc.",
	Example: "nacosctl delete [resource]",
}
