package main

import "fmt"

func main() {
	IsEvenOrOdd(creatNumber(10))
}

//creat any number to slice
//新增數字到切片
func creatNumber(m int) []int {
	n := []int{}
	for i := 0; i <= m; i++ {
		n = append(n, i)
	}
	return n
}

//check the number even or odd
//分辨奇數或偶數
func IsEvenOrOdd(n []int) {
	for _, dif := range n {
		if dif%2 == 0 {
			fmt.Printf("%d is even\n", dif)
		} else {
			fmt.Printf("%d is odd\n", dif)
		}
	}
}
