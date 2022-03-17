package goutils

import (
	"fmt"
	"testing"
)

func TestQuickSort(t *testing.T) {
	arr := []int{3, 5, 49, 2, 4, 2, 1}
	qsort(arr)
	fmt.Println(arr)
}
