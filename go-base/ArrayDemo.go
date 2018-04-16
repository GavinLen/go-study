package main

import "fmt"

/**
Array Demo
 */
func main() {

	// 先定义后赋值
	var arr [5]int
	for i := 0; i < 5; i++ {
		arr[i] = i + 100
	}
	fmt.Println(arr)

	// 定义并赋值
	nums := []int{1, 2, 3, 4, 5}
	fmt.Println(nums)
}
