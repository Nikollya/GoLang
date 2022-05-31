package main

import "fmt"

func read() (int, string) {
	var x string
	fmt.Println("Write number of card!")
	fmt.Scan(&x)
	n := len(x)
	return n, x
}
func check(n int, x string) bool {
	if !f1(n) {
		return false
	}
	if !f2(n, x) {
		return false
	}
	if !f3(n, x) {
		return false
	}
	fmt.Println("Your card have accepted!")
	return true
}
func f1(n int) bool {
	if n != 16 {
		fmt.Println("Error: 01")
		return false
	}
	return true
}
func f2(n int, a string) bool {
	for i := 0; i < n; i++ {
		if a[i] < '0' || a[i] > '9' {
			fmt.Println("Error: 02")
			return false
		}
	}
	return true
}
func f3(n int, a string) bool {
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
		return false
	}
	return true
}
func main() {
	n, a := read()
	check(n, a)
	return
}

//5168441223630339
