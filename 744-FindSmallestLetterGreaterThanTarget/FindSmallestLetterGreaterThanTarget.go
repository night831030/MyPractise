package main

//尋找比目標大的最小字母
func nextGreatestLetter(letters []byte, target byte) (ans byte) {
	for _, i := range letters {
		if i > target {
			return i
		}
	}
	return letters[0]
}
