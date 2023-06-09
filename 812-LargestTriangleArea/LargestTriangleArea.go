package main

//最大三角形面積
func largestTriangleArea(points [][]int) (ans float64) {
	n := len(points)
	for i := 0; i < n-2; i++ { //以行列式公式求三角形面積
		for j := i + 1; j < n-1; j++ {
			for k := j + 1; k < n; k++ {
				x1, y1 := points[i][0], points[i][1]
				x2, y2 := points[j][0], points[j][1]
				x3, y3 := points[k][0], points[k][1]
				tg := x1*y2 + x2*y3 + x3*y1 - x1*y3 - x2*y1 - x3*y2
				if tg < 0 {
					tg *= -1
				}
				tga := float64(tg) / 2
				if tga > ans {
					ans = tga
				}
			}
		}
	}
	return
}
