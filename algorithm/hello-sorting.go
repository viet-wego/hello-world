package main

import (
  "fmt"
  "math/rand"
  "github.com/viettd/hello-world/algorithm/sorting"
)

const MAX_VALUE = 10000000
const LENGTH = 1000

func main()  {
  srcArray := make([]int, LENGTH)
  for i := 0; i < len(srcArray); i++ {
    srcArray[i] = rand.Intn(MAX_VALUE)
  }
  fmt.Println("Source array: ", srcArray)
  arrayToBeSort := make([]int, LENGTH)

  copy(arrayToBeSort, srcArray)
  fmt.Println("Result: ", sorting.BubbleSort(arrayToBeSort))

  copy(arrayToBeSort, srcArray)
  fmt.Println("Result: ", sorting.InsertionSort(arrayToBeSort))


}
