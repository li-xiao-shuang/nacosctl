package cmd

import (
	"github.com/spf13/cobra"
	"nacosctl/cmd/basic"
	"nacosctl/common/constant"
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
	rootCmd.AddCommand(basic.CreateCmd)
	rootCmd.AddCommand(basic.DeleteCmd)
	rootCmd.AddCommand(basic.UpdateCmd)
	rootCmd.AddCommand(basic.GetCmd)

}
