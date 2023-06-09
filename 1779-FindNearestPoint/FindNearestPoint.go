package main

import "math"

// import "fmt"

func main() {
}

//找最近同軸座標
func NearestValidPoint(x int, y int, points [][]int) int {
	z := -1                            //預設為無同一線上座標，回傳-1
	ans := math.MaxInt64               //最近距離預設int最大值
	for i := 0; i < len(points); i++ { //以迴圈比對X,Y軸座標
		if points[i][0] == x || points[i][1] == y {
			a := abs(points[i][0]-x) + abs(points[i][1]-y) //a為計算最小距離
			if a < ans {                                   //最小距離為答案，如小於現有最小值則取代
				ans = a
				z = i
			}
		}
	}
	return z

}

//距離無負數
func abs(x int) int {
	if x < 0 {
		return -x
	} else {
		return x
	}
}
