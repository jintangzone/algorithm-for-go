package array

func numberOfBoomerangs(points [][]int) int {
	size := len(points)
	res := 0
	for i := 0; i < size; i++ {
		records := make(map[int]int)
		for j := 0; j < size; j++ {
			if i != j {
				records[dis(points[i], points[j])]++
			}
		}
		for _, v := range records {
			res += v * (v - 1)
		}
	}
	return res
}

func dis(p1, p2 []int) int {
	return (p2[0]-p1[0])*(p2[0]-p1[0]) + (p2[1]-p1[1])*(p2[1]-p1[1])
}
