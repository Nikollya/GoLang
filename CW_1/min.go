package main

import "fmt"

func solution(str, ending string) bool {
	n, m := len(str), len(ending)
	if m > n {
		return false
	}
	for i := 1; i <= m; i++ {
		if str[n-i] != ending[m-i] {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(solution("aba", "ba"))
}
