package main

import (
	"fmt"
	"math/rand"
)

func quickSort(a []int) []int {
	if len(a) < 2 {
		return a
	}

	left, right := 0, len(a)-1

	// Pick a pivot
	pivotIndex := rand.Int() % len(a)

	// Move the pivot to the right
	a[pivotIndex], a[right] = a[right], a[pivotIndex]

	// Pile elements smaller than the pivot on the left
	for i := range a {
		if a[i] < a[right] {
			a[i], a[left] = a[left], a[i]
			left++
		}
	}

	// Place the pivot after the last smaller element
	a[left], a[right] = a[right], a[left]

	// Go down the rabbit hole
	quickSort(a[:left])
	quickSort(a[left+1:])

	return a
}
func main() {
	var t int
	fmt.Scan(&t)
	for w := 0; w < t; w++ {
		var n int
		fmt.Scan(&n)
		a := make([]int, n)
		for i := 0; i < n; i++ {
			fmt.Scan(&a[i])
		}
		quickSort(a)
		f := true
		for i := 0; i < n-1; i++ {
			if !((a[i] == a[i+1]) || (a[i]+1 == a[i+1])) {
				f = false
				break
			}
		}
		if f {
			fmt.Println("YES")
		} else {
			fmt.Println("NO")
		}
	}
}
