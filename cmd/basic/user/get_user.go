package user

import (
	"github.com/spf13/cobra"
	"nacosctl/process/user"
)

var GetUserCmd = &cobra.Command{
	Use:     "users",
	Short:   "get all users",
	Long:    "get all users",
	Example: "nacosctl get users",
	Args:    cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		user.ParseGetUsersCmd(cmd)
	},
}
