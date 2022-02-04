package common

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func GetServerAddress(cmd *cobra.Command) (address string) {
	viperAddress := viper.Get("nacosctl.server.address").(string)
	return viperAddress
}
