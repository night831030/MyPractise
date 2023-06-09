package main

func main() {
}

//回文數
func IsPalindrome(x int) bool {
	if x < 0 || x != 0 && x%10 == 0 {
		return false
	}
	ans := false
	y, z := 0, x
	for x > y {
		y = (y * 10) + (z % 10)
		z /= 10
	}
	if x == y {
		ans = true
	}
	return ans
}
