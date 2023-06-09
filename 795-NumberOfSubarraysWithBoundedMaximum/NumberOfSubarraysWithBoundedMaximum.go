package main

//計算區間子數總數
//在n時，所有符合條件的總數為->目前達成最小條件數量-最後超額的數量
//當前超額則為0，不足則為前一個符合-超額
func numSubarrayBoundedMax(nums []int, left int, right int) (Answer int) {
	over, on := 0, 0 //超過最大值及在最小值之上的標籤
	for n, i := range nums {
		if i > right { //超額時為第幾個
			over = n + 1
		}
		if i >= left { //達成最小值為第幾個
			on = n + 1
		}
		Answer += on - over
	}
	return
}
