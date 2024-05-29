package main

import (
	"fmt"
	. "project_dict/laba2_array/array"
)

func main() {
	arrayInt1 := []int{12, 214142, 1231, 123, 1, 2342, 12, 123, 123}
	arrayInt2 := []int{12, 214142, 1231, 123, 1, 2342, 12, 123, 123}
	arrayInt3 := []int{12, 214142, 1231, 123, 1, 2342, 12, 123, 123}

	QuickSort(arrayInt1, 0, len(arrayInt1)-1, '+')
	QuickSort(arrayInt2, 0, len(arrayInt2)-1, '-')
	QuickSortShuffle(arrayInt3, 0, len(arrayInt3)-1)

	fmt.Println(arrayInt1)
	fmt.Println(arrayInt2)
	fmt.Println(arrayInt3)
}
