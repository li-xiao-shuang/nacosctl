package main

import (
	"fmt"
	"nacos-cli/cmd"
	"nacos-cli/common/loader"
	"os"
)

func main() {
	loader.Init()
	if err := cmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
