package merge

import (
	"math/rand"
	"time"
)

func GenerateSlice(size int) []int {
	slice := make([]int, size, size)
	// UNIX 時間をシードにして乱数生成器を用意する
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < size; i++ {
		slice[i] = rand.Intn(999) - rand.Intn(999)
	}
	return slice
}

func MergeSort(items []int) []int {
	var num = len(items)

	if num == 1 {
		return items
	}

	middle := int(num / 2) // この書き方何？
	var (
		left  = make([]int, middle)
		right = make([]int, num-middle)
	)
	for i := 0; i < num; i++ {
		if i < middle {
			left[i] = items[i]
		} else {
			right[i-middle] = items[i]
		}
	}
	return merge(MergeSort(left), MergeSort(right))
}

// 再帰関数はこうなる
// return merge(
// 	merge(
// 		items,
// 		merge(items, items),
// 	),
// 	merge(
// 		items,
// 		merge(items, items),
// 	),
// )

func merge(left, right []int) (result []int) {
	result = make([]int, len(left)+len(right))

	i := 0
	for len(left) > 0 && len(right) > 0 {
		if left[0] < right[0] {
			result[i] = left[0]
			left = left[1:]
		} else {
			result[i] = right[0]
			right = right[1:]
		}
		i++
	}
	for j := 0; j < len(left); j++ {
		result[j] = left[j]
		i++
	}
	for j := 0; j < len(right); j++ {
		result[j] = right[j]
		i++
	}
	return
}
