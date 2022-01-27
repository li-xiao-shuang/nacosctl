package command

import (
	"github.com/spf13/cobra"
	"nacos-cli/cmd/namespace"
)

var CreateCmd = &cobra.Command{
	Use:   "create [type]",
	Short: "create command",
}

func init() {
	CreateCmd.AddCommand(namespace.NamespaceCmd)
}
