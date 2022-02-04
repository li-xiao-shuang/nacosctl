package user

import (
	"github.com/spf13/cobra"
	"nacosctl/printer"
	"nacosctl/process/user"
)

var DeleteUserCmd = &cobra.Command{
	Use:     "user [username] ",
	Short:   "delete a user",
	Long:    "delete a user",
	Example: "nacosctl delete user [username]",
	Args:    cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		if args[0] == "" {
			printer.Red("[error]:username must not be null")
			printer.Yellow("[see]:nacosctl delete user -h")
			return
		}
		user.ParseDeleteUserCmd(cmd, args[0])
	},
}
