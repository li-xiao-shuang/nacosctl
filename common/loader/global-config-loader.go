package loader

import (
	"github.com/spf13/viper"
	"nacosctl/config/constant"
)

// viperConfig 全局配置变量
var viperConfig *viper.Viper

func ConfigInit() error {
	viper.SetConfigFile("./config.yaml") // 指定配置文件路径
	viper.SetConfigName("config")        // 配置文件名称(无扩展名)
	viper.SetConfigType("yaml")          // 如果配置文件的名称中没有扩展名，则需要配置此项
	viper.AddConfigPath(".")             // 还可以在工作目录中查找配置
	err := viper.ReadInConfig()          // 查找并读取配置文件
	if err != nil {                      // 处理读取配置文件的错误
		// 没有读取到文件就设置默认值
		viperConfig = viper.New()
		viperConfig.SetDefault("server.address", constant.DefaultAddress)
		viperConfig.SetDefault("server.port", 8848)
	} else {
		viperConfig = viper.GetViper()
	}
	return nil
}

func GetViper() *viper.Viper {
	return viperConfig
}
