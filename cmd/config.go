package cmd

import (
	"github.com/spf13/cobra"
)

var configCmd = &cobra.Command{
	Use:   "config",
	Short: "Related operations on configuration information",
}

func init() {
	RootCmd.AddCommand(configCmd)
}
