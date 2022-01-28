package namespace

import (
	"github.com/spf13/cobra"
	"nacosctl/process/namespace"
)

var ListNamespaceCmd = &cobra.Command{
	Use:     "namespaces",
	Short:   "Returns all namespace information",
	Long:    "Returns all namespace informatione",
	Example: "nacosctl get namespaces",
	Args:    cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		namespace.ParseGetNamespaceListCmd(cmd)
	},
}
