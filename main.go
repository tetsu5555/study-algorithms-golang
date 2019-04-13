package main

import (
	"fmt"

	"./bubble"
	"./insertion"
	"./merge"
	"./quick"
)

func main() {
	fmt.Println("Bubble sort started")
	bubble.BubbleSort(bubble.ToBeSorted)
	fmt.Println("Bubble sort end")

	fmt.Println("Insertion sort started")
	slice1 := insertion.GenerateSlice(20)
	sorted := insertion.InsertionSort(slice1)
	fmt.Println(sorted)
	fmt.Println("Insertion sort end")

	slice2 := merge.GenerateSlice(20)
	fmt.Println("\n-- Unsorted --- \n\n", slice2)
	fmt.Println("\n-- Sorted --- \n\n", merge.MergeSort(slice2))

	slice3 := merge.GenerateSlice(20)
	// ソートして結果を出力する
	fmt.Println("\n-- Unsorted --- \n\n", slice3)
	fmt.Println("\n-- Sorted --- \n\n", quick.QuickSort(slice3))
}
