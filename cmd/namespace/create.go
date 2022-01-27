package namespace

import (
	"github.com/spf13/cobra"
)

var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create namespace",
	Long:  "The account that creates the namespace must have sufficient permissions",
}

func init() {
	createCmd.AddCommand(NamespaceCmd)
}

func AddNamespaceCmd(rootCmd *cobra.Command) {
	rootCmd.AddCommand(createCmd)
}
