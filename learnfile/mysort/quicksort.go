package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	var date []int
	n := 19
	seedNum := time.Now().UnixNano()
	for i := 0; i < n; i++ {
		rand.Seed(seedNum)
		date = append(date, rand.Intn(1001))
		seedNum++
	}
	fmt.Println(date)
	QuickSort(date, 0, len(date)-1)
	fmt.Println(date)
}

//快速排序法
func QuickSort(date []int, left, right int) {
	if left < right {
		mark := date[right]
		var rs []int
		l := left
		for i := left; i < right; i++ {
			if date[i] > mark {
				rs = append(rs, date[i])
			} else {
				date[l] = date[i]
				l++
			}
		}
		date[l] = mark
		copy(date[l+1:], rs)
		if left < l-1 {
			QuickSort(date, left, l-1)
		}
		if l+1 < right {
			QuickSort(date, l+1, right)
		}
	}
}
