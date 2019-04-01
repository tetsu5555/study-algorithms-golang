package main

import (
	"fmt"
)

var toBeSorted [10]int = [10]int{1, 3, 2, 4, 8, 6, 7, 2, 3, 0}

func bubbleSort(input [10]int) {
	// listの中の要素の数
	n := 10
	// swappedにtrueをセット
	swapped := true
	// loop
	for swapped {
		// swappedにfalseをセット
		swapped = false
		// iterate through all of the elements in our list
		for i := 1; i < n; i++ {
			// 現在の要素が次の要素よりも大きければ入れ替える
			if input[i-1] > input[i] {
				fmt.Println("Swapping")
				// Go's tuple assignmentを使って入れ替える
				input[i], input[i-1] = input[i-1], input[i]
				swapped = true
			}
		}
	}
	// ソートされたリストをprintする
	fmt.Println(input)
}

func main() {
	fmt.Println("Hello World")
	bubbleSort(toBeSorted)
}
