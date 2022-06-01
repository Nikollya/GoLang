package main

import "fmt"

func main() {
	var s, n int
	fmt.Scan(&s, &n)
	var x [10000]int
	var y [10000]int
	var u [10000]int
	for i := 0; i < n; i++ {
		fmt.Scan(&x[i], &y[i])
		u[i] = 0
	}
	f := true
	localChange := false
	for f {
		localChange = false
		for i := 0; i < n; i++ {
			if u[i] == 0 {
				if x[i] < s {
					u[i] = 1
					s += y[i]
					localChange = true
				}
			}
		}
		if localChange == false {
			fmt.Println("NO")
			return
		}
		f = false
		for i := 0; i < n; i++ {
			if u[i] == 0 {
				f = true
				break
			}
		}
	}
	fmt.Println("YES")
}
