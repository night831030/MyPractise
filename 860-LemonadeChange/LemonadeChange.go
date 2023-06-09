package main

//檸檬水找零
func lemonadeChange(bills []int) bool {
	five, ten := 0, 0
	for _, bill := range bills {
		switch bill {
		case 5:
			five++
		case 10:
			ten++
			if five >= 1 {
				five--
			} else {
				return false
			}
		case 20:
			if five >= 1 && ten >= 1 {
				ten--
				five--
			} else if five >= 3 {
				five -= 3
			} else {
				return false
			}
		}
	}
	return true
}
