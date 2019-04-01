package bubble

import (
	"testing"
)

func TestCalculate(t *testing.T) {
	var toBeSorted [10]int = [10]int{1, 3, 2, 4, 8, 6, 7, 2, 3, 0}
	sorted2 := BubbleSort(toBeSorted)
	if sorted2 != [10]int{0, 1, 2, 2, 3, 3, 4, 6, 7, 8} {
		t.Error("Expected to be Sorted")
	}
}
