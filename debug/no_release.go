// +build !release

package debug

import (
	"fmt"
	"log"
)

func Println(v ...interface{}) {
	log.Println(v...)
}

func Print(v ...interface{}) {
	fmt.Print(v...)
}
