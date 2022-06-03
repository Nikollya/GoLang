package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	var t int
	fmt.Fscan(in, &t)
	for i := 0; i < t; i++ {
		var x int
		fmt.Fscan(in, &x)
		y := 0
		for {
			if x&y > 0 && x^y > 0 {
				fmt.Println(y)
				break
			}
			y++
		}
	}
}
