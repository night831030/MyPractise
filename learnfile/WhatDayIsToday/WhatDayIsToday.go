package main

import "fmt"

func main() {
	year := 2552
	week := 6
	testtime := 6
	monthday := [][]int{{15, 20}, {8, 30}, {2, 29}, {11, 31}, {0, 13}, {9, 19}}
	i := 0
	for testtime > 0 {
		fmt.Printf("%d ", yeartest(year, monthday[i][0], monthday[i][1], week))
		i++
		testtime--

	}
}

func yeartest(year int, month int, day int, week int) int {
	if month > 12 || month < 1 {
		return -1
	}
	if !daycheck(year, month, day) {
		return -2
	}
	daysum := week
	for m := 1; m < month; m++ {
		switch {
		case m == 2:
			if year%4 == 0 {
				daysum += 29
			} else {
				daysum += 28
			}
			break
		case m == 4 || m == 6 || m == 9 || m == 11:
			daysum += 30
			break
		default:
			daysum += 31
		}
	}
	daysum += (day - 1)
	return daysum % 7
}

func daycheck(year int, month int, day int) bool {
	out := true
	if day < 1 {
		out = false
	}
	switch {
	case month == 1 || month == 3 || month == 5 || month == 7 || month == 8 || month == 10 || month == 12:
		if day > 31 || day < 0 {
			out = false
		}
	case month == 2:
		if year%4 == 0 {
			if day >= 30 || day < 1 {
				out = false
			}
		}
		if year%4 != 0 {
			if day >= 29 || day < 1 {
				out = false
			}
		}
	case month == 2 || month == 4 || month == 6 || month == 9 || month == 11:
		if day > 30 || day < 1 {
			out = false
		}
	}
	return out
}

// year := 1492
// week := 0
// testtime := 7
// monthday := [][]int{{10, 12}, {-1, 0}, {12, -12}, {8, 3}, {6, 31}, {4, 9}, {11, 3}}
// 5 -1 -2 5 -2 1 6

// year := 1911
// week := 0
// testtime := 8
// monthday := [][]int{{4, 27}, {2, 29}, {12, 25}, {7, -5}, {11, 6}, {2, 0}, {10, 10}, {3, 100}}
// 4 -2 1 -2 1 -2 2 -2

// year := 2552
// week := 6
// testtime := 6
// monthday := [][]int{{15, 20}, {8, 30}, {2, 29}, {11, 31}, {0, 13}, {9, 19}}
// -1 3 2 -2 -1 2

// year := 2005
// week := 6
// testtime := 7
// monthday := [][]int{{-77, -132}, {12, 25}, {8, 31}, {7, 6}, {1, 1}, {1024, 122}, {2, 29}}
// -1 0 3 3 6 -1 -2

// year := 2000
// week := 6
// testtime := 10
// monthday := [][]int{{2, 29}, {1300, 1}, {11, 19}, {6, 1272}, {8, 31}, {4, 9}, {9, 23}, {2, 28}, {1, 1}, {10, 1}}
// 2 -1 0 -2 4 0 6 1 6 0
