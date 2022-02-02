package login

import (
	"github.com/spf13/cobra"
	"nacosctl/printer"
	"nacosctl/process/login"
)

var NLoginCmd = &cobra.Command{
	Use:     "login [username] [password]",
	Short:   "Set username and password for nacos authentication",
	Long:    "Set username and password for nacos authentication",
	Example: "nacosctl login [resource]",
	Args:    cobra.MinimumNArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		if args[0] == "" {
			printer.Red("[error]:username must not be null")
			printer.Yellow("[see]:nacosctl login  -h")
			return
		}
		if args[1] == "" {
			printer.Red("[error]:password must not be null")
			printer.Yellow("[see]:nacosctl login  -h")
			return
		}
		login.ParseLoginCmd(cmd, args[0], args[1])
	},
}
