package cmd

import (
	"errors"
	"github.com/spf13/cobra"
	"nacos-cli/config"
)

var getConfigCmd = &cobra.Command{
	Use:   "get [dataId] [group]",
	Short: "Get configuration information",
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 2 {
			return errors.New("Two parameters are required,[dataId] and [group]")
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		config.GetCommand(args[0], args[1])
	},
}

func init() {
	configCmd.AddCommand(getConfigCmd)
}
