package config

import (
	"errors"
	"github.com/spf13/cobra"
	"nacosctl/process/config"
)

var GetConfigListCmd = &cobra.Command{
	Use:     "configs [pageNo] [pageSize]",
	Short:   "Get a list of configuration information in paginated form",
	Long:    "Get a list of configuration information in paginated form",
	Example: "nacosctl get configs [pageNo] [pageSize]",
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 2 {
			return errors.New("2 parameters are required,[pageNo] [pageSize]")
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		config.ParseGetConfigListCmd(args[0], args[1], cmd)
	},
}
