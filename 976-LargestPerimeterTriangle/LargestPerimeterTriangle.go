package main

import "sort"

func main() {

}

//三角形最大周長
func LargestPerimeter(nums []int) int {
	sort.Ints(nums)
	for i := len(nums) - 1; i >= 2; i-- {
		if nums[i-1]+nums[i-2] > nums[i] {
			return nums[i] + nums[i-1] + nums[i-2]
		}
	}
	return 0
}
