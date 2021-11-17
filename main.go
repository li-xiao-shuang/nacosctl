package main

import (
	"fmt"
	"nacos-cli/cmd"
	"nacos-cli/common"
	"os"
)

func main() {
	common.Init()
	if err := cmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
