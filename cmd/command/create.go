package command

import "github.com/spf13/cobra"

var createCmd = &cobra.Command{
	Use:     "create",
	Short:   "command to create some resources",
	Long:    "Create commands can be executed on resources such as configurations, services, namespaces, users, etc.",
	Example: "nacosctl create [resource]",
}
