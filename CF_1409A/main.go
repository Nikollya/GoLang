package main

import (
	"bufio"
	"fmt"
	"os"
)

var in = bufio.NewReader(os.Stdin)

func main() {
	var t int
	fmt.Scan(&t)
	for p := 0; p < t; p++ {
		var a, b int
		fmt.Fscan(in, &a, &b)
		if a == b {
			fmt.Println(0)
		} else {
			if a < b {
				a, b = b, a
			}
			x := a - b
			k := int(x / 10)
			if k*10 != x {
				fmt.Println(k + 1)
			} else {
				fmt.Println(k)
			}
		}
	}
}
