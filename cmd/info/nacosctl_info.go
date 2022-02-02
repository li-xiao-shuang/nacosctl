package info

import (
	"github.com/spf13/cobra"
)

var InfoCmd = &cobra.Command{
	Use:   "info ",
	Short: "Get nacosctl configuration file information",
	Long:  "Get nacosctl configuration file information",
}

func init() {
	InfoCmd.AddCommand(UserNameCmd)
	InfoCmd.AddCommand(PasswordCmd)
	InfoCmd.AddCommand(ServerAddressCmd)
}
