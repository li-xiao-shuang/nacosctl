package main

import (
	"fmt"
	"nacos-cli/cmd"
	"nacos-cli/config"
	"os"
)

func main() {
	config.Init()
	if err := cmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
