package main

import (
	"fmt"
	"os"
)

func main() {
	/*s, sep := "", ""
	for _, arg := range os.Args[1:] {
		s += sep + arg
		sep = " "
		break
	}*/

	/*
		s := ""
		var s string
		var s = ""
		var s string = ""
	*/
	fmt.Println(os.Args[1:])
	//fmt.Println(s)
}
