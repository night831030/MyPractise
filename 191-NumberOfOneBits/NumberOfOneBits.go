package main

func main() {

}

//二進制1總數
func HammingWeight(num uint32) int {
	ans := 0
	for num > 0 {
		if num%2 == 1 {
			ans++
		}
		num /= 2
	}
	return ans
}
