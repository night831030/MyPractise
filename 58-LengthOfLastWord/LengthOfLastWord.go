package main

//計算最後一個單字長度
func lengthOfLastWord(s string) int {
	length := 0
	i := len(s) - 1
	for s[i] == ' ' {
		i--
	}
	for i >= 0 && s[i] != ' ' {
		length++
		i--
	}
	return length
}
