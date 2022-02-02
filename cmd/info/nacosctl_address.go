package info

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"nacosctl/printer"
)

var ServerAddressCmd = &cobra.Command{
	Use:   "address ",
	Short: "return nacosctl server address information",
	Long:  "return nacosctl server address information",
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		printer.Cyan(viper.GetString("nacosctl.server.address"))
	},
}
