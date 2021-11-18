package config

import (
	"errors"
	"github.com/spf13/cobra"
	"nacos-cli/config"
	"nacos-cli/config/constant"
)

var versionConfigCmd = &cobra.Command{
	Use:   "version [id] [namespaceId] [dataId] [group]",
	Short: "Query the configuration information of the previous version",
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 4 {
			return errors.New("4 parameters are required,[id] [namespaceId] [dataId] [group]")
		}
		if args[3] == "" {
			args[3] = constant.DefaultGroup
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		config.VersionCommand(args[0], args[1], args[2], args[3], cmd)
	},
}

func init() {
	versionConfigCmd.Flags().StringP("address", "b", constant.DefaultAddress, "nacos server ip address")
	versionConfigCmd.Flags().StringP("port", "p", constant.DefaultPort, "nacos server port")
}
