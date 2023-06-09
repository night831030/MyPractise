package main

//計算矩陣中的負數
func countNegatives(grid [][]int) (ans int) {
	for _, i := range grid {
		for _, j := range i {
			if j < 0 {
				ans++
			}
		}
	}
	return
}

// func countNegatives(grid [][]int) (ans int) {
// 	var j int
// 	for _, i := range grid {
// 		j = 0
// 		for ; j < len(i) && i[j] >= 0; j++ {
// 		}
// 		ans += len(i) - j
// 	}
// 	return
// }
