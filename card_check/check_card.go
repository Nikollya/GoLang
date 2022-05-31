package main

import "fmt"

func main() {
	fmt.Println("Write number of card!")
	var a string
	fmt.Scan(&a)
	n := len(a)
	if n != 16 {
		fmt.Println("Error: 01")
		return
	}
	for i := 0; i < n; i++ {
		if a[i] < '0' || a[i] > '9' {
			fmt.Println("Error: 02")
			return
		}
	}
	sum := 0
	x := 0
	for i := n - 1; i >= 0; i-- {
		if i%2 == 0 {
			if int(a[i]-'0')*2 > 9 {
				x = int(a[i]-'0') * 2
				sum += x % 10
				x = x / 10
				sum += x
			} else {
				sum += int(a[i]-'0') * 2
			}
		} else {
			sum += int(a[i] - '0')
		}
	}
	if sum%10 != 0 {
		fmt.Println("Error: 03")
		return
	}
	fmt.Println("Your card have accepted!")
	return
}
