package main

import (
	"fmt"
	"nacosctl/cmd"
	"nacosctl/common/loader"
	"os"
)

func main() {
	loader.ConfigInit()
	if err := cmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}