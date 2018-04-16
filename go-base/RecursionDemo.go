package main

import "fmt"

/**
递归 Demo
 */
func main() {

	num := factorial(-1)
	fmt.Println("阶乘结果：", num)
}

/* 阶乘 */
func factorial(num int) int {
	if num < 0 {
		return 0
	}

	if num == 1 || num == 0 {
		return 1
	} else {
		return num * factorial(num-1)
	}
}
