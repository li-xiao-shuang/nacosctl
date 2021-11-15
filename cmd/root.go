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

func initConfig() {
	//viper.SetConfigFile("./config.yaml") // 指定配置文件路径
	//viper.SetConfigName("config")        // 配置文件名称(无扩展名)
	//viper.SetConfigType("yaml")          // 如果配置文件的名称中没有扩展名，则需要配置此项
	//
	//var author string
	//RootCmd.Flags()
	//RootCmd.PersistentFlags().StringVar(&author, "author", "YOUR NAME", "Author name for copyright attribution")
	//viper.BindPFlag("author", RootCmd.PersistentFlags().Lookup("author"))
	//fmt.Println(author)
}
