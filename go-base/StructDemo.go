package main

import "fmt"

/*
结构体 Demo
 */

/* 结构体 */
type person struct {
	name string
	age  int
}

func main() {

	// 先定义后赋值
	personA := person{}
	personA.name = "Tom"
	personA.age = 12

	printPerson(personA)
	printPersonInfo(personA)

	// 定义时全部赋值
	personB := person{"Lisa", 12}
	printPerson(personB)
	printPersonInfo(personB)

	// 定义时部分赋值
	personC := person{name: "Jim"}
	personC.age = 100
	printPerson(personC)
	printPersonInfo(personC)

	// 指针指向
	personD := &person{"张三", 12}
	fmt.Println(personD)

	personD.age = 100
	printPersonInfo(personB)

	fmt.Println(&personA)

}

/**
输出 person
 */
func printPerson(person person) {
	fmt.Println(person)
}

/* 输出 person Info */
func printPersonInfo(person person) {
	fmt.Printf("person Info: Name [%s], Age [%d]\n", person.name, person.age)
}
