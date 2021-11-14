package cmd

import (
	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use:     "nacos-cli",
	Aliases: []string{"nacos", "na"},
	Short:   "Nacos command line tool",
	Long:    `A fast and flexible command line tool .Created by the nacos community。`,
}

func Execute() error {
	return RootCmd.Execute()
}

//func initConfig() {
//	// 勿忘读取config文件，无论是从cfgFile还是从home文件
//	if cfgFile != "" {
//		viper.SetConfigName(cfgFile)
//	} else {
//		// 找到home文件
//		home, err := homedir.Dir()
//		if err != nil {
//			fmt.Println(err)
//			os.Exit(1)
//		}
//
//		// 在home文件夹中搜索以“.cobra”为名称的config
//		viper.AddConfigPath(home)
//		viper.SetConfigName(".cobra")
//	}
//	// 读取符合的环境变量
//	viper.AutomaticEnv()
//
//	if err := viper.ReadInConfig(); err != nil {
//		fmt.Println("Can not read config:", viper.ConfigFileUsed())
//	}
//}
