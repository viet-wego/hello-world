package main

import "fmt"

func main() {
	m := make(map[string]int)

	m["k1"] = 1
	m["k2"] = 2
	m["k3"] = 3

	fmt.Println("map: ", m)

	delete(m, "k3")
	fmt.Println("delete: ", m)

	value, exist := m["k1"]
	fmt.Println("Does key k1 exist: ", exist, ", and the value is: ", value)

	declaredMap := map[string]int{"key3": 3, "key2": 1}
	fmt.Println("declaredMap: ", declaredMap)
}
