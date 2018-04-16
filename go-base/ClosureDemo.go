package main

import "fmt"

/*
闭包 Demo
 */
func main() {

	funcA := intSeq()
	fmt.Println(funcA())
	fmt.Println(funcA())

	funcB := intSeq()
	fmt.Println(funcB())
	fmt.Println(funcB())
	fmt.Println(funcB())
	fmt.Println(funcB())
	fmt.Println(funcB())
}

func intSeq() func() int {
	index := 0

	return func() int {
		index += 1
		return index
	}

}
