package main

import "fmt"

/*
函数 Demo
 */
func main() {

	sum := sum(1, 2, 3, 4, 5, 6)
	fmt.Printf("Sum:%d\n", sum)
}

/* 求和 */
func sum(nums ...int) int {
	fmt.Println(nums)

	sum := 0
	for index, value := range nums {
		fmt.Printf("Index [%d], Value [%d]\n", index, value)
		sum += value
	}
	return sum
}

/* 交换 */
func exchange(a int, b int) (int, int) {
	return b, a
}

/* 求和 */
func maxNum(a int, b int) int {
	if a >= b {
		return a
	} else {
		return b
	}
}
