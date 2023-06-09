package main

//密鑰格式化
func licenseKeyFormatting(s string, k int) string {
	var answer []byte
	i := 0
	for j := len(s) - 1; j >= 0; j-- {
		if s[j] != '-' {
			if i == k {
				answer = append(answer, '-')
				i = 0
			}
			if s[j] >= 'a' && s[j] <= 'z' {
				answer = append(answer, byte(s[j]-'a'+'A'))
				i++
			} else {
				answer = append(answer, byte(s[j]))
				i++
			}
		}
	}
	return string(reverse(answer))
}

func reverse(s []byte) []byte {
	var re []byte
	for i := len(s) - 1; i >= 0; i-- {
		re = append(re, s[i])
	}
	return re
}
