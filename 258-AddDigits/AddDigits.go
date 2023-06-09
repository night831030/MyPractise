package main

//各位數相加
func addDigits(num int) int {
	if num < 10 {
		return num
	}
	ans := 0
	for num >= 10 {
		ans += num % 10
		num /= 10
		if num < 10 {
			ans += num
		}
	}
	if ans < 10 {
		return ans
	} else {
		return addDigits(ans)
	}
}

// {
// 	for num >= 10 {
// 		ans := 0
// 		for ; num > 0; num /= 10 {
// 			ans += num % 10
// 		}
// 		num = ans
// 	}
// 	return num
// }
