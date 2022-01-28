package cmd

import (
	"github.com/spf13/cobra"
	"nacos-cli/cmd/basecmd"
	"nacos-cli/config/constant"
)

var rootCmd = &cobra.Command{
	Use:   "nacosctl",
	Short: "Nacos command line tool",
	Long:  "A fast and flexible command line tool, Created by the nacos community.",
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	// global flags
	rootCmd.PersistentFlags().StringP("address", "a", constant.DefaultAddress, "nacos server ip address")
	rootCmd.PersistentFlags().StringP("port", "p", constant.DefaultPort, "nacos server port")

	// base command
	rootCmd.AddCommand(basecmd.CreateCmd)
	rootCmd.AddCommand(basecmd.DeleteCmd)
	rootCmd.AddCommand(basecmd.UpdateCmd)
	rootCmd.AddCommand(basecmd.GetCmd)

	//rootCmd.AddCommand(command.CreateCmd)
	//rootCmd.AddCommand(config.ConfigCmd)
	//rootCmd.AddCommand(instance.InstanceCmd)
}
