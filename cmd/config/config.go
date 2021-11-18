package config

import (
	"github.com/spf13/cobra"
)

var ConfigCmd = &cobra.Command{
	Use:   "config",
	Short: "Related operations on configuration information",
}

func init() {
	ConfigCmd.AddCommand(getConfigCmd)
	ConfigCmd.AddCommand(addConfigCmd)
	ConfigCmd.AddCommand(deleteConfigCmd)
	ConfigCmd.AddCommand(listConfigCmd)
	ConfigCmd.AddCommand(versionConfigCmd)
}
