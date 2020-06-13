package array

func fourSumCount(A []int, B []int, C []int, D []int) int {

	records := make(map[int]int)
	res := 0

	for i := 0; i < len(A); i++ {
		for j := 0; j < len(B); j++ {
			records[A[i]+B[j]]++
		}
	}

	for i := 0; i < len(C); i++ {
		for j := 0; j < len(D); j++ {
			if v, ok := records[0-C[i]-D[j]]; ok {
				res += v
			}
		}
	}

	return res
}
