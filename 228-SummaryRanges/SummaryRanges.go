package main

import "fmt"

//範圍彙整
func summaryRanges(nums []int) (answer []string) {
	for i := 0; i < len(nums); {
		start := i
		for i++; i < len(nums) && nums[i] == nums[i-1]+1; i++ {
		}
		if start == i-1 { //須扣除循環結束時的i++
			answer = append(answer, fmt.Sprintf("%d", nums[start]))
		} else {
			answer = append(answer, fmt.Sprintf("%d->%d", nums[start], nums[i-1]))
		}
	}

	return
}
