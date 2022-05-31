package main

import "fmt"

func min(a int, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

func main() {
	var n, m int
	fmt.Scan(&n, &m)
	var a [31]int
	var b [31]int
	var k, x = 0, 0
	for i := 0; i < n; i++ {
		fmt.Scan(&a[i], &b[i])
		k += a[i]
		x += b[i]
	}
	if m < k || m > x {
		fmt.Println("NO")
		return
	}
	fmt.Println("YES")
	for i := 0; i < n; i++ {
		fmt.Print(min(b[i], m-k+a[i]), " ")
		m -= min(b[i], m-k+a[i])
		k -= a[i]
	}
}
