package main

import (
	"fmt"
	"strconv"
)

/**
Map Demo
 */
func main() {

	// 创建 Map
	m := make(map[string]int)
	for i := 0; i < 5; i++ {
		key := "k_" + strconv.Itoa(i)
		m[key] = i
	}
	printMap(m)

	// 声明时赋值
	mA := map[string]int{"A": 1, "B": 2}
	printMap(mA)

	// add
	m["key_add"] = 100
	printMap(m)

	// delete
	delete(m, "k_1")
	printMap(m)

	// get Value
	b := m["key_1add"]
	fmt.Printf("Value:%d", b)

}

/* 输出 Map */
func printMap(m map[string]int) {
	fmt.Println(m)
}
