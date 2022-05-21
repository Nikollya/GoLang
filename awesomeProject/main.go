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
	fmt.Println(0, " ", os.Args[0])
	fmt.Println(1, " ", os.Args[1:])
	//fmt.Println(s)
}
