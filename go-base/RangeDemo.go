package main

import "fmt"

/**
Range Demo
 */
func main() {

	// range Slice
	sli := []int{101, 102, 103, 104, 105, 106}
	printSlice(sli)

	fmt.Println("Range Slice:")
	for index, value := range sli {
		fmt.Printf("Index [%d], Value [%d]\n", index, value)
	}

	// range Map
	m := map[string]string{"A": "a", "B": "b"}
	fmt.Println("Range map:")
	for key, value := range m {
		fmt.Printf("Key [%s] : Value [%s]\n", key, value)
	}

	// range string
	str := "ABCDEFG"

	fmt.Println("Range stirng:")
	for index, value := range str {
		fmt.Printf("Index [%d], value [%c]\n", index, value)
	}

}

/* 输出 Slice */
func printSlice(sli []int) {
	fmt.Println(sli)
}
