package namespace

import (
	"errors"
	"github.com/spf13/cobra"
	"nacos-cli/config/constant"
	"nacos-cli/namespace"
)

var delNamespace = &cobra.Command{
	Use:   "del [namespaceId]",
	Short: "Delete a namespace",
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return errors.New("1 parameters are required,[namespaceId]")
		}
		if args[0] == "" {
			return errors.New("[namespaceId] can not be blank")
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		namespace.DelCommand(cmd, args[0])
	},
}

func init() {
	delNamespace.Flags().StringP("address", "b", constant.DefaultAddress, "nacos server ip address")
	delNamespace.Flags().StringP("port", "p", constant.DefaultPort, "nacos server port")
}
