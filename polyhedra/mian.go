package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	mp := make(map[string]int)
	mp["Tetrahedron"] = 0  //4
	mp["Cube"] = 0         //6
	mp["Octahedron"] = 0   //8
	mp["Dodecahedron"] = 0 //12
	mp["Icosahedron"] = 0  //20
	for i := 0; i < n; i++ {
		var s string
		fmt.Scan(&s)
		mp[s]++
	}
	var counter int
	counter = mp["Tetrahedron"]*4 + mp["Cube"]*6 + mp["Octahedron"]*8 + mp["Dodecahedron"]*12 + mp["Icosahedron"]*20
	fmt.Println(counter)
}
