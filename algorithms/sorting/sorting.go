package sorting

import "fmt"

func BubbleSort(array []int) []int  {
  fmt.Println("\nBubble sort")
  fmt.Println("Source array: ", array)
  for i := 0; i < len(array) - 1; i++ {
    swapped := false
    for j := 0; j < len(array) - 1 - i; j++ {
      fmt.Printf(" Compare %d, %d", array[j], array[j + 1])
      if array[j] > array[j + 1] {
        swap(array, j, j + 1)
        swapped = true
        fmt.Printf(" => swapped: %d, %d\n", array[j], array[j + 1])
      } else {
        fmt.Println(" => not swap")
      }
    }
    if !swapped {
      break;
    }
    fmt.Printf("Iteration #%d: %v\n", i + 1, array)
  }

  return array
}

func SelectionSort(array []int) []int  {
  fmt.Println("\nSelection sort")
  fmt.Println("Source array: ", array)
  for i := 0; i < len(array) - 1; i++ {
    minIndex := i
    for j := i + 1; j < len(array); j++ {
      if array[j] < array[minIndex] {
        minIndex = j
      }
    }
    fmt.Printf(" Min value:%d at index %d", array[minIndex], minIndex)
    if minIndex != i {
      swap(array, i, minIndex)
      fmt.Printf(" => swapped %d, %d\n", array[i], array[minIndex])
    } else {
      fmt.Println(" => not swap")
    }
    fmt.Printf("Iteration #%d: %v\n", i + 1, array)
  }

  return array
}

func MergeSort(array []int) []int  {
  fmt.Println("\nMerge sort")
  fmt.Println("Source array: ", array)
  tempArray := make([]int, len(array))
  doMergeSort(array, tempArray, 0, len(array) - 1)
  fmt.Println("Result array: ", array)
  return array
}

func doMergeSort(array, tempArray []int, indexLow, indexHigh int)  {
  if indexLow < indexHigh {
    indexMidle := indexLow + (indexHigh - indexLow) / 2
    doMergeSort(array, tempArray, indexLow, indexMidle)
    doMergeSort(array, tempArray, indexMidle + 1, indexHigh)
    mergeParts(array, tempArray, indexLow, indexMidle, indexHigh)
  }
}

func mergeParts(array, tempArray []int, indexLow, indexMidle, indexHigh int)  {
  fmt.Println(array, tempArray)
  fmt.Println(indexLow, indexMidle, indexHigh)
  for i := indexLow; i <= indexHigh; i++ {
    tempArray[i] = array[i]
  }
  i := indexLow
  j := indexMidle + 1
  k := indexLow
  for ; i <= indexMidle && j <= indexHigh; k++ {
    if tempArray[i] <= tempArray[j] {
      array[k] = tempArray[i]
      i++
    } else {
      array[k] = tempArray[j]
      j++
    }
  }
  for ; i <= indexMidle; i++ {
    array[k] = tempArray[i]
    k++
  }
}

func QuickSort(array []int) []int  {
  fmt.Println("\nQuick sort")
  fmt.Println("Source array: ", array)
  doQuickSort(array, 0, len(array) - 1)
  return array
}

func doQuickSort(array []int, left, right int)  {
  if left < right {
    pivot := array[right]
    fmt.Printf(" Left: %d, Right: %d, Pivot: %d\n", array[left], array[right], pivot)
    partitionPoint := partition(array, left, right, pivot)
    doQuickSort(array, left, partitionPoint - 1)
    doQuickSort(array, partitionPoint + 1, right)
  }
  fmt.Println(array)
}

func partition(array []int, left, right, pivot int) int  {
  leftPointer := left-1
  rightPointer := right
  for {
    for array[leftPointer] < pivot {
      leftPointer++
    }
    for array[rightPointer] >  pivot {
      rightPointer--
    }
    if leftPointer >= rightPointer {
      break;
    } else{
      swap(array, leftPointer, rightPointer)
      fmt.Printf(" swapped: %d, %d\n", array[leftPointer], array[rightPointer])
    }
  }
  swap(array, leftPointer, rightPointer)
  fmt.Printf(" Pivot swapped: %d, %d\n", array[leftPointer], array[right])
  return leftPointer
}

func swap(array []int, index1, index2 int)  {
  temp := array[index1]
  array[index1] = array[index2]
  array[index2] = temp
}
