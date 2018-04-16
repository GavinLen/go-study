package main

import "fmt"

type Book struct {
	id   int
	name string
}

func main() {

	// 基本类型
	i := 1
	iPtr := &i
	fmt.Println(i)
	fmt.Println(iPtr)
	fmt.Println(*iPtr)

	// String 类型
	str := "abc"
	strPtr := &str
	fmt.Println(str)
	fmt.Println(strPtr)
	fmt.Println(*strPtr)

	// 数组类型
	arr := []int{1, 2, 3}
	arrPtr := &arr
	fmt.Println(arr)
	fmt.Println(arrPtr)

	// slice 类型
	sli := make([]int, 2, 5)
	sliPtr := &sli
	fmt.Println(sli)
	fmt.Println(sliPtr)
	fmt.Println(*sliPtr)

	// 结构体
	book := Book{1, "abc"}
	bookPtr := &book
	fmt.Println(book)
	fmt.Println(bookPtr)
	fmt.Println(*bookPtr)

https://books.studygolang.com/gobyexample/structs/

}
