package loader

import (
	"github.com/spf13/viper"
	"nacosctl/common/constant"
)

// viperConfig 全局配置变量
var viperConfig *viper.Viper

func ConfigInit() error {
	viper.SetConfigName("nacosctl") // 配置文件名称(无扩展名)
	viper.SetConfigType("yaml")     // 如果配置文件的名称中没有扩展名，则需要配置此项
	viper.AddConfigPath(".")        // 还可以在工作目录中查找配置

	viper.SetDefault("nacosctl.server.address", constant.DefaultAddress)
	viper.SetDefault("nacosctl.user", "")
	viper.SetDefault("nacosctl.password", "")

	err := viper.ReadInConfig() // 查找并读取配置文件
	if err != nil {             // 处理读取配置文件的错误
		viper.SafeWriteConfigAs("nacosctl.yaml")
	}
	return nil
}
