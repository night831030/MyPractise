package main

//單字拼湊
func countCharacters(words []string, chars string) (ans int) {
	for _, word := range words { //遍歷字串
		char := chars
		have := false
		for wn := 0; wn < len(word); wn++ { //遍歷單字
			t := false
			for cn := 0; cn < len(char); cn++ { //比對字元
				if word[wn] == char[cn] {
					t = true
					if cn == len(char)-1 { //移除已使用字元
						char = char[:cn]
					} else {
						char = char[:cn] + char[cn+1:]
					}
					break
				}
			}
			if !t {
				break
			}
			if wn == len(word)-1 {
				have = true
			}
		}
		if have {
			ans += len(word)
		}
	}
	return
}
