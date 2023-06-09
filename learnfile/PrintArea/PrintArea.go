package main

import "fmt"

func main() {
	square1 := square{20}
	triangle1 := triangle{25, 40}
	PrintArea(square1)
	PrintArea(triangle1)
}

//面積共通介面
type Area interface {
	getArea() float64
}

//方形結構
type square struct {
	sideLength float64
}

//三角型結構
type triangle struct {
	height float64
	base   float64
}

//計算方形面積
func (s square) getArea() float64 {
	return s.sideLength * s.sideLength
}

//計算三角形面積
func (t triangle) getArea() float64 {
	return 0.5 * t.height * t.base
}

//輸出
func PrintArea(a Area) {
	fmt.Println("Area is ", a.getArea())
}
