package cmd

import (
	"github.com/spf13/cobra"
	"nacos-cli/cmd/config"
	"nacos-cli/cmd/namespace"
)

var rootCmd = &cobra.Command{
	Use:   "nacos-cli",
	Short: "Nacos command line tool",
	Long:  `A fast and flexible command line tool .Created by the nacos community。`,
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	rootCmd.AddCommand(config.ConfigCmd)
	rootCmd.AddCommand(namespace.NamespaceCmd)
}
