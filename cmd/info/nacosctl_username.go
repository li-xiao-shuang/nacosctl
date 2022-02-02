package info

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"nacosctl/printer"
)

var UserNameCmd = &cobra.Command{
	Use:   "username ",
	Short: "return nacosctl username information",
	Long:  "return nacosctl username information",
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		printer.Cyan(viper.GetString("nacosctl.username"))
	},
}
