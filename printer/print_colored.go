package printer

import (
	"fmt"
	"github.com/logrusorgru/aurora"
)

func Cyan(content interface{}) {
	fmt.Println(aurora.Cyan(content))
}

func Red(content interface{}) {
	fmt.Println(aurora.Red(content))
}

func Yellow(content interface{}) {
	fmt.Println(aurora.Yellow(content))
}
