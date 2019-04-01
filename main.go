package main

import (
	"fmt"

	"./bubble"
	"./insertion"
)

func main() {
	fmt.Println("Bubble sort started")
	bubble.BubbleSort(bubble.ToBeSorted)
	fmt.Println("Bubble sort end")

	fmt.Println("Insertion sort started")
	slice := insertion.GenerateSlice(20)
	sorted := insertion.InsertionSort(slice)
	fmt.Println(sorted)
	fmt.Println("Insertion sort end")
}
