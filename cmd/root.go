package cmd

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"nacosctl/cmd/basic"
	"nacosctl/cmd/info"
	"nacosctl/cmd/login"
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
	rootCmd.PersistentFlags().StringP("address", "a", constant.DefaultAddress, "nacos server ip+port")
	viper.BindPFlag("nacosctl.server.address", rootCmd.Flag("address"))

	// base command
	rootCmd.AddCommand(basic.CreateCmd)
	rootCmd.AddCommand(basic.DeleteCmd)
	rootCmd.AddCommand(basic.UpdateCmd)
	rootCmd.AddCommand(basic.GetCmd)

	//login
	rootCmd.AddCommand(login.NLoginCmd)

	//local
	rootCmd.AddCommand(info.InfoCmd)

}
