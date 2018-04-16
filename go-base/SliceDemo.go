package main

import (
	"fmt"
)

/**
Slice Demo
 */
func main() {
	arr := []int{1, 2, 3, 4, 5}

	// 指定元素个数
	sliA := make([] int, len(arr))
	printSliceInfo(sliA)

	// 指定元素个数与容量
	sliB := make([] int, 5, 11)
	printSliceInfo(sliB)

	// copy
	copy(sliB, arr)
	printSliceInfo(sliB)
	printElement(sliB)

	// append
	sliB = append(sliB, 2)
	printSliceInfo(sliB)
	printElement(sliB)
}

/** 输出 Slice 元素 */
func printElement(sli []int) {
	for index, value := range sli {
		fmt.Printf("Index[%d]-Value[%d]\n", index, value)
	}
}

/** 输出 Slice Info */
func printSliceInfo(sli []int) {
	fmt.Printf("Slice: len [%d], cap [%d]\n", len(sli), cap(sli))
}
