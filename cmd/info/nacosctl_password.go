package info

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"nacosctl/printer"
)

var PasswordCmd = &cobra.Command{
	Use:   "password ",
	Short: "return nacosctl password information",
	Long:  "return nacosctl password information",
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		printer.Cyan(viper.GetString("nacosctl.password"))
	},
}
