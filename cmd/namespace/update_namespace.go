package namespace

import (
	"errors"
	"github.com/spf13/cobra"
	"nacos-cli/config/constant"
	"nacos-cli/namespace"
)

var updateNamespaceCmd = &cobra.Command{
	Use:   "update [namespaceId] [namespaceName] [namespaceDesc]",
	Short: "Update the name and description of the namespace",
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
		namespace.UpdateCommand(cmd, args[0], args[1], args[2])
	},
}

func init() {
	updateNamespaceCmd.Flags().StringP("address", "b", constant.DefaultAddress, "nacos server ip address")
	updateNamespaceCmd.Flags().StringP("port", "p", constant.DefaultPort, "nacos server port")
}
