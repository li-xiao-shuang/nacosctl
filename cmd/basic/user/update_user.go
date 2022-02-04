package user

import (
	"github.com/spf13/cobra"
	"nacosctl/printer"
	"nacosctl/process/user"
)

var UpdateUserCmd = &cobra.Command{
	Use:     "user [username] [newpassword]",
	Short:   "update  user password",
	Long:    "update  user password",
	Example: "nacosctl update user [oldpassword] [newpassword]",
	Args:    cobra.MinimumNArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		if args[0] == "" {
			printer.Red("[error]:username must not be null")
			printer.Yellow("[see]:nacosctl update user -h")
			return
		}
		if args[1] == "" {
			printer.Red("[error]:newpassword must not be null")
			printer.Yellow("[see]:nacosctl update user -h")
			return
		}
		user.ParseUpdateUserCmd(args[0], args[1])
	},
}
