package main

import (
	"fmt"
	"math"
)

const s = "Bye bye bye"

func main() {
	fmt.Println(s)

	const n = 100

	const x = 3e10 / n
	fmt.Println(x)

	fmt.Println(int64(x))

	fmt.Println(math.Sin(x))
}
