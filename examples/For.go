package main

import "fmt"

func main() {
	i := 0
	for i < 10 {
		fmt.Println(i)
		i++
	}

	for j := 9; j >= 0 ; j-- {
		fmt.Println(j)
	}

	for {
		fmt.Println("Loop")
		break;
	}

	for k := 0; k < 10; k++ {
		if k % 2 == 0 {
			fmt.Println(k)
		}
	}
}
