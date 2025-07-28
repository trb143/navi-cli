package helper

import (
	"fmt"
)

var Verbose bool

func HandleError(e error) {
	if e != nil {
		fmt.Println(e)
	}
}

func CheckPrint(p string) {
	if Verbose {
 		fmt.Println(p)
	}
}
