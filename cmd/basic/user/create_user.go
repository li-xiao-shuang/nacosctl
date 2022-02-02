package user

import (
	"github.com/spf13/cobra"
	"nacosctl/printer"
	"nacosctl/process/user"
)

var CreateUserCmd = &cobra.Command{
	Use:     "user [username] [password]",
	Short:   "create a user",
	Long:    "create a user",
	Example: "nacosctl create user [username] [password]",
	Args:    cobra.MinimumNArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		if args[0] == "" {
			printer.Red("[error]:username must not be null")
			printer.Yellow("[see]:nacosctl create user -h")
			return
		}
		if args[1] == "" {
			printer.Red("[error]:password must not be null")
			printer.Yellow("[see]:nacosctl create user -h")
			return
		}
		user.ParseCreateUserCmd(cmd, args[0], args[1])
	},
}
