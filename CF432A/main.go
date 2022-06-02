package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
)

var in = bufio.NewReader(os.Stdin)

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
	var n, k int
	fmt.Fscan(in, &n, &k)
	a := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &a[i])
	}
	quickSort(a)
	kk := 0
	for i := 0; i < n; i++ {
		a[i] += k
		if a[i] > 5 {
			if kk < 3 {
				fmt.Println(0)
			} else {
				fmt.Println(int(kk / 3))
			}
			return
		}
		kk++
	}
	fmt.Println(int(n / 3))
}
