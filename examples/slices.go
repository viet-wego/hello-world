package main

import "fmt"

func main() {
	s := make([]string, 3)
	fmt.Println("empty: ", s)

	s[0] = "aha"
	s[1] = "oho"
	s[2] = "ihi"

	fmt.Println("set: ", s)
	fmt.Println("get: ", s[2])

	fmt.Println("len ", len(s))

	s = append(s, "ehe")
	s = append(s, "uhu")
	fmt.Println("append: ", s)

	c := make([]string, len(s))
	copy(c, s)
	fmt.Println("copy: ", c)

	l := s[2:4]
	fmt.Println("subslice 2:4 ", l)
	l = s[:4]
	fmt.Println("subslice :4 ", l)
	l = s[2:]
	fmt.Println("subslice 2: ", l)

	d := []string{"abc", "def"}
	fmt.Println("declare: ", d)

	d = append(d,"f***")
	fmt.Println(d)

	twoD := make([][]int, 3)
	for i := 0; i < 3; i++{
		innerLength := i + 1
		twoD[i] = make([]int, innerLength)
		for j := 0; j < innerLength; j++{
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2D: ", twoD)
}
