package info

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"nacosctl/printer"
)

var AccessTokenCmd = &cobra.Command{
	Use:   "accesstoken ",
	Short: "return nacosctl accesstoken information",
	Long:  "return nacosctl accesstoken information",
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		printer.Cyan(viper.GetString("nacosctl.accesstoken"))
	},
}
