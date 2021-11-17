package cmd

import (
	"github.com/spf13/cobra"
)

var configCmd = &cobra.Command{
	Use:   "config",
	Short: "Related operations on configuration information",
}

func init() {
	rootCmd.AddCommand(configCmd)
	configCmd.AddCommand(getConfigCmd)
	configCmd.AddCommand(addConfigCmd)
	configCmd.AddCommand(deleteConfigCmd)
	configCmd.AddCommand(listConfigCmd)
	configCmd.AddCommand(versionConfigCmd)
}
