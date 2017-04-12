package main

import (
	"fmt"
	"time"
)

func main() {
	i := 2
	fmt.Println("Write ", i, " as")
	switch i {
	case 1:
		fmt.Println("One")
	case 2:
		fmt.Println("Two")
	case 3:
		fmt.Println("Three")
	default:
		fmt.Println("Unknow")
	}

	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("It's the weekend.")
	default:
		fmt.Println("It's a week day.")
	}

	t := time.Now().Hour()
	fmt.Println(t)
	switch {
	case t < 12:
		fmt.Println("It's before noon.")
	default:
		fmt.Println("It's after noon.")
	}

	whatAmI := func(i interface{}) {
		switch t := i.(type) {
		case bool:
			fmt.Println("bool")
		case int:
			fmt.Println("int")
		default:
			fmt.Printf("unknow type %T\n", t)
		}
	}
	whatAmI(false)
	whatAmI(35)
	whatAmI("halo")
}
