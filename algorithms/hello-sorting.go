package main

import (
  "fmt"
  "math/rand"
  "github.com/viettd/hello-world/algorithms/sorting"
)

const MAX_VALUE = 10
const LENGTH = 10

func main()  {
  srcArray := make([]int, LENGTH)
  for i := 0; i < len(srcArray); i++ {
    srcArray[i] = rand.Intn(MAX_VALUE)
  }
  fmt.Println("Source array: ", srcArray)
  arrayToBeSort := make([]int, LENGTH)

  copy(arrayToBeSort, srcArray)
  sorting.BubbleSort(arrayToBeSort)

  copy(arrayToBeSort, srcArray)
  fmt.Println("Result: ", sorting.SelectionSort(arrayToBeSort))

  //copy(arrayToBeSort, srcArray)
  //fmt.Println("Result: ", sorting.InsertionSort(arrayToBeSort))


}
