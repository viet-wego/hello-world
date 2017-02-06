package sorting

import (
  "fmt"
  "time"
  )

func BubbleSort(array []int) []int  {
  fmt.Println("\nBubble sort")
  start := time.Now().UnixNano()
  fmt.Println(" - Start: ", start)

  for i := 0; i < len(array) - 1; i++ {
    for j := i; j < len(array); j++ {
      if array[i] > array[j] {
        temp := array[i]
        array[i] = array[j]
        array[j] = temp
      }
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
