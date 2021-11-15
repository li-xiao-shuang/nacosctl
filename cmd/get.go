package cmd

import (
	"errors"
	"fmt"
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
		fmt.Println(cmd.Flag("address").Value.String())
		fmt.Println(cmd.Flag("port").Value.String())
		config.GetCommand(args[0], args[1], cmd.Flag("address").Value.String(), cmd.Flag("port").Value.String())
	},
}

func init() {
	getConfigCmd.Flags().StringP("address", "b", "127.0.0.1", "nacos server ip address")
	getConfigCmd.Flags().StringP("port", "p", "8848", "nacos server port")
	configCmd.AddCommand(getConfigCmd)
}
