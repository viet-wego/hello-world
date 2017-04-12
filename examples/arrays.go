package main

import "fmt"

func main() {
	var a [5]int
	fmt.Println("empty: ", a)

	a[3] = 20
	fmt.Println("set: ", a)
	fmt.Println("get: ", a[3])

	fmt.Println("len: ", len(a))

	b := [3]int{3, 5, 7}
	fmt.Println("declare: ", b)

	var c [2][3]int
	for i := 0; i < 2; i++{
		for j := 0; j < 3; j++{
			c[i][j] = i + j
		}
	}
	fmt.Println("2D: ", c)
}
