package main

import "fmt"

func main() {
	x := 199
	fmt.Println("x = ", x)
	if x %2 == 0 {
		fmt.Println("x is even")
	} else {
		fmt.Println("x is odd")
	}

	if n := 3; n < 0 {
		fmt.Println("n is negative")
	} else if n < 10 {
		fmt.Println("n has 1 digit")
	} else {
		fmt.Println("n has multiple digits")
	}
}
