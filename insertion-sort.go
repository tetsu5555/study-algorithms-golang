package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// make([]int, length, capacity)
	slice := generateSlice(20)
	fmt.Println("\n--- Unsorted --- \n\n", slice)
	insertionSort(slice)
	fmt.Println("\n--- Sorted ---\n\n", slice, "\n")
}

func generateSlice(size int) []int {
	slice := make([]int, size, size)
	// 時刻を使って初期のSeedを設定する
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < size; i++ {
		slice[i] = rand.Intn(999) - rand.Intn(999)
	}
	return slice
}

func insertionSort(items []int) []int {
	n := len(items)
	for i := 1; i < n; i++ {
		j := i
		for j > 0 {
			if items[j - 1] > items[j] {
				items[j - 1], items[j] = items[j], items[j - 1]
			}
			j = j- 1
		}
	}
	return items
}
