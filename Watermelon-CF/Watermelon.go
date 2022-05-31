package main

import "fmt"

func main() {
	var a int
	fmt.Scan(&a)
	if a%2 != 0 || a <= 3 {
		fmt.Println("NO")
	} else {
		fmt.Println("YES")
	}
}
