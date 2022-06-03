package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var t int
	in := bufio.NewReader(os.Stdin)
	fmt.Fscan(in, &t)
	for e := 0; e < t; e++ {
		var n int
		fmt.Fscan(in, &n)
		a := make([]int, n)
		f1, f2 := false, true
		kn := 0
		pos := 0
		km := 0
		for i := 0; i < n; i++ {
			fmt.Fscan(in, &a[i])
			if a[i]%2 != 0 {
				f1 = true
			}

			if a[i]%2 == 0 {
				p := a[i]
				kl := 0
				for u := 0; ; u++ {
					kl++
					p /= 2
					if p%2 != 0 {
						break
					}
				}
				if f2 || (kl < km) {
					pos = i
					km = kl
				}
				f2 = false
				kn++
			}
			//fmt.Println(km, "////*")
		}
		if f1 && f2 {
			fmt.Println(0)
		} else {
			if f1 {
				fmt.Println(kn)
			} else {
				k := 0
				for {
					a[pos] /= 2
					k++
					//fmt.Println(pos, "//")
					if a[pos]%2 != 0 {
						fmt.Println(k + kn - 1)
						break
					}
				}
			}
		}

	}
}
