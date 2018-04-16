package main

/* 图形 */
type Shape interface {
	// 面积
	area()
	// 周长
	perimeter() float32
}

// https://www.yiibai.com/go/go_interfaces.html
/* 圆 */
type Circle struct {
	x, y, radius float32
}

/* 矩形 */
type Rectangle struct {
	width, height float32
}

func main() {

}
