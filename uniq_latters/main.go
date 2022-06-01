package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	s, _ := in.ReadString('\n')
	n := len(s) - 2
	mp := make(map[string]int)
	count := 0
	for i := 1; i < n; i++ {
		if s[i] >= 'a' && s[i] <= 'z' {
			if mp[string(s[i])] == 0 {
				mp[string(s[i])] = 1
				count++
			}
		}
	}
	fmt.Println(count)
}
