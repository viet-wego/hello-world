package sorting

import "fmt"

func BubbleSort(array []int) []int  {
  fmt.Println("\nBubble sort")
  fmt.Println("Source array: ", array)
  for i := 0; i < len(array) - 1; i++ {
    swapped := false
    for j := 0; j < len(array)-1-i; j++ {
      fmt.Printf(" Compare %d, %d", array[j], array[j+1])
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
  fmt.Println("Source array: ", array)
  for i := 0; i < len(array) - 1; i++ {
    minIndex := i
    for j := i+1; j < len(array); j++ {
      if array[j] < array[minIndex] {
        minIndex = j
      }
    }
    fmt.Printf(" Min value:%d at index %d", array[minIndex], minIndex)
    if minIndex != i {
      temp := array[i]
      array[i] = array[minIndex]
      array[minIndex] = temp
      fmt.Printf(" => swapped %d, %d\n", array[i], array[minIndex])
    } else {
      fmt.Println(" => not swap")
    }
    fmt.Printf("Iteration #%d: %v\n", i+1, array)
  }

  return array
}

func InsertionSort(array []int) []int  {
  fmt.Println("\nInsertion sort")

  return array
}
