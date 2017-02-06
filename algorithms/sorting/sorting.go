package sorting

import (
  "fmt"
  "time"
  )

func BubbleSort(array []int) []int  {
  fmt.Println("\nBubble sort")

  for i := 0; i < len(array) - 1; i++ {
    swapped := false
    for j := 0; j < len(array)-1-i; j++ {
      fmt.Printf(" - Compare %d, %d", array[j], array[j+1])
      if array[j] > array[j+1] {
        temp := array[j]
        array[j] = array[j+1]
        array[j+1] = temp
        swapped = true
        fmt.Printf(" => swapped: %d, %d\n", array[j], array[j+1])
      } else {
        fmt.Println(" => not swap")
      }
    }
    if !swapped {
      break;
    }
    fmt.Printf("Iteration #%d: %v\n", i+1, array)
  }

  return array
}

func SelectionSort(array []int) []int  {
  fmt.Println("\nSelection sort")
  start := time.Now().UnixNano()
  fmt.Println(" - Start: ", start)

  for i := 0; i < len(array) - 1; i++ {
    min := i
    for j := i+1; j < len(array); j++ {
      if array[j] < array[min] {
        min = j
      }
    }
    if min != i {
      temp := array[i]
      array[i] = array[min]
      array[min] = temp
    }
  }

  end := time.Now().UnixNano()
  fmt.Println(" - End: ", end)
  fmt.Println(" - Time exec(ns): ", end - start)

  return array
}

func InsertionSort(array []int) []int  {
  fmt.Println("\nInsertion sort")
  start := time.Now().UnixNano()
  fmt.Println(" - Start: ", start)

  end := time.Now().UnixNano()
  fmt.Println(" - End: ", end)
  fmt.Println(" - Time exec(ns): ", end - start)

  return array
}
