package main

import (
	"errors"
	"fmt"
)

func read() (int, string) {
	var x string
	fmt.Println("Write number of card!")
	fmt.Scan(&x)
	n := len(x)
	return n, x
}
func checkNumberOfCard(n int, a string) error {
	if n != 16 {
		return errors.New("error: 01")
	}
	for i := 0; i < n; i++ {
		if a[i] < '0' || a[i] > '9' {
			return errors.New("error: 2")
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
		return errors.New("error: 03")
	}
	return errors.New("your card have accepted")
}

func main() {
	n, a := read()
	err := checkNumberOfCard(n, a)
	fmt.Println(err)
	return
}

//5168441223630339
