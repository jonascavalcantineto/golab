package main

import (
	"fmt"
	"golabs/algo"
)

func main() {
	arr := []int{5, 50, 1, 2}
	algo.SelectionSort(arr)
	fmt.Println(arr)

}