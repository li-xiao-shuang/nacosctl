package common

import (
	"github.com/spf13/viper"
)

func GetServerAddress() (address string) {
	viperAddress := viper.Get("nacosctl.server.address").(string)
	return viperAddress
}
