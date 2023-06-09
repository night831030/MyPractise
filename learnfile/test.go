package main

import (
	"fmt"
	"log"
	"os"
	"reflect"
	"strings"
)

func main() {

}

func Calc(n int) int {
	x, y := 0, 0 // x,y分別代表1、2的次數
	for x <= n { //當1等於n結束
		if (n-x)%2 == 0 { //n-1的次數求2的次數
			y = (n - x) / 2
		}
		if x+2*y == n { //當1、2的次數相加等於n，則為其中一種解
			fmt.Printf("move 1 square * %d + 2 square * %d \n", x, y)
		}
		x++
	}

	return (n / 2) + 1
}

func IsPair(s string) bool {
	if len(s)%2 != 0 { //若字串數量為奇數則直接回傳false
		return false
	}
	start, end := 0, len(s)-1 //字串頭尾標籤
	for start < end {         //依序比對頭尾字元是否相同
		if s[start] != s[end] { //不同直接回傳false
			return false
		}
		start++ //移動標籤
		end--
	}
	return true //若過程中無直接回傳false，則比對正確，回傳true
}

func ReverseInteger(i int) int {
	x, y := 0, i //設y為臨時變數，x為解
	if i < 0 {   //若i為負數，先將負數拿掉
		y = -i
	}
	for y >= 1 { //將y的個位數給x，並將y/10、x*10，依次將數調換
		x = (x * 10) + (y % 10)
		y /= 10
	}
	if i < 0 { //若i為負數，將負數取回
		x *= -1
	}
	return x
}

func RemoveDuplicated(input []int) []int {
	var ans []int             //ans為解
	for _, i := range input { //讀取input陣列
		k := true               //先設陣列皆不重複
		for _, j := range ans { //比對input將放入的數與ans中是否重複
			if i == j {
				k = false
			}
		}
		if k { //比對後無重複則將數放入ans
			ans = append(ans, i)
		}
	}
	return ans
}

//求質數
func p() {
	i := 2
	var n [101]int
	for ; i <= len(n); i++ {
		for j := i; j <= len(n); j++ {
			if i*j <= len(n) {
				n[i*j] = 1
			}
		}
	}
	for j := 1; j <= len(n); j++ {
		if n[j] == 0 {
			println(j)
		}
	}
}

//刪除重複
func ftest() {
	in := []int{1, 2, 3, 3, 4, 5, 1, 6, 1, 5, 1, 7, 1, 2}
	var ans []int
	for i := 0; i <= len(in)-1; i++ {
		k := true
		for j := 0; j <= len(ans)-1; j++ {
			if in[i] == ans[j] {
				k = false
			}
		}

		if k == true {
			ans = append(ans, in[i])
		}
	}

	for a, b := range ans {
		println(a, b)
	}
}

func learn() {
	println("YA\nY\tA, \"YA\", \\YA\\") //\n-換行、\t-跳格、\\前後斜線
	fmt.Println('A')                    //''轉換為ASCII
	println(reflect.TypeOf(1))          //顯示指標
	fmt.Println(strings.Title("2.75,ya i am small word"))
	fileInfo, err := os.Stat("my.txt") //讀取檔案my.txt
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(fileInfo.Size()) //檔案大小
}

func test() {
	amap := map[int]int{1: 1, 2: 2}
	var value int
	var ok bool
	value, ok = amap[3]
	println(value, ok)
	value = amap[1]
	_, ok = amap[2]
	println(value, ok)
}
