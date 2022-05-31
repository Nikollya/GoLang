package main

import (
	"bufio"
	"fmt"
	"os"
)

var in = bufio.NewReader(os.Stdin)

func main() {
	var n int
	var st string
	fmt.Scan(&n)
	m := make(map[string]int)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &st)
		if m[st] == 0 {
			fmt.Println("OK")
			m[st] = 1
		} else {
			fmt.Printf("%s%d\n", st, m[st])
			m[st]++
		}
	}
}
