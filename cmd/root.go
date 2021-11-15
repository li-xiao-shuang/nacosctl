package cmd

import (
	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use:   "nacos-cli",
	Short: "Nacos command line tool",
	Long:  `A fast and flexible command line tool .Created by the nacos community。`,
}

func Execute() error {
	return RootCmd.Execute()
}
