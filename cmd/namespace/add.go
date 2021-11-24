package namespace

import (
	"errors"
	"github.com/spf13/cobra"
	"nacos-cli/config/constant"
	"nacos-cli/namespace"
)

var addNamespaceCmd = &cobra.Command{
	Use:   "add [namespaceId] [namespaceName] [namespaceDesc]",
	Short: "Add a namespace",
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 3 {
			return errors.New("3 parameters are required,[namespaceId] [namespaceName] [namespaceDesc]")
		}
		if args[0] == "" {
			return errors.New("[namespaceId] can not be blank")
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		namespace.AddCommand(cmd, args[0], args[1], args[2])
	},
}

func init() {
	addNamespaceCmd.Flags().StringP("address", "b", constant.DefaultAddress, "nacos server ip address")
	addNamespaceCmd.Flags().StringP("port", "p", constant.DefaultPort, "nacos server port")
}
