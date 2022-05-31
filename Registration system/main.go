package main

import (
	"fmt"
)

func main() {
	var n int
	var s string
	fmt.Scan(&n)
	m := make(map[string]int)
	for i := 0; i < n; i++ {
		fmt.Scan(&s)
		if m[s] == 0 {
			fmt.Println("OK")
			m[s] = 1
		} else {
			fmt.Print(s, m[s])
			fmt.Println()
			m[s]++
		}
	}
}
