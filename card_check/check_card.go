package main

import (
	"errors"
	"fmt"
)

func read() string {
	var x string
	fmt.Println("Write number of card!")
	fmt.Scan(&x)
	return x
}
func checkNumberOfCard(a string) (bool, error) {
	if len(a) != 16 {
		return false, errors.New("error: 01")
	}
	for i := 0; i < len(a); i++ {
		if a[i] < '0' || a[i] > '9' {
			return false, errors.New("error: 2")
		}
	}
	sum := 0
	x := 0
	for i := len(a) - 1; i >= 0; i-- {
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
		return false, errors.New("error: 03")
	}
	return true, nil
}

func main() {
	//a := read()
	a := "5168441223630339"
	card, err := checkNumberOfCard(a)
	if card {
		fmt.Println("your card have accepted")
	} else {
		fmt.Println(err)
	}
	return
}

//5168441223630339
