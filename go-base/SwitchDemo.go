package main

import (
	"fmt"
	"time"
)

/**
Switch Demo
 */
func main() {

	// 成绩等级判断
	grade := 100
	var class = class(grade)
	fmt.Printf("成绩[%d],等级[%s]\n", grade, class)

	// 周末判断
	t := time.Now()
	b := weekend(t)
	if b {
		fmt.Printf("Yes\n")
	} else {
		fmt.Printf("No\n")
	}

	// 楼层判断
	fmt.Printf("楼层: %s\n", floor(2))

}

/** 楼层判断 */
func floor(num int) string {
	switch num {
	case 1:
		return "One"
	case 2:
		return "two"
	case 3:
		return "three"
	default:
		return "unknown"
	}
}

/** 周末判断 */
func weekend(t time.Time) bool {
	switch t.Weekday() {
	case time.Saturday, time.Sunday:
		return true
	default:
		return false
	}
}

/** 成绩等级判断 */
func class(grade int) string {
	switch {
	case grade > 90:
		return "A"
	case grade > 80:
		return "B"
	case grade > 70:
		return "C"
	case grade > 60:
		return "D"
	default:
		return "E"
	}
}
