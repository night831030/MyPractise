package main

//計算檔案最長絕對路徑
func lengthLongestPath(input string) (answer int) {
	ans := [99]int{}
	tan := 0 ///////////長度可改為保存路徑
	floor, now, filecheck := 0, 0, false
	for i, n := 0, len(input); i < n; i++ {
		switch input[i] { //分類項目條件
		case '\n': //結算前一層長度
			ans[floor] = tan
			tan = 0
			floor = 0
		case '\t': //計算檔案層數
			floor++
		case '.': //'.'副檔名前墜，確定為檔案
			filecheck = true
			now = floor
			tan++
		default:
			tan++
		}
		if i == n-1 { //絕對路徑結尾不會在遇到'\n'，直接結算
			ans[floor] = tan
			tan = 0
			floor = 0
		}
		if filecheck {
			if input[i] == '\n' || i == n-1 {
				length := 0
				for j := 0; j <= now; j++ { //加總檔案長度
					length += ans[j] + 1
				}
				length--
				if length > answer { //比較
					answer = length
				}
				filecheck = false
			}
		}
	}
	return answer
}
