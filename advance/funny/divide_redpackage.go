package funny

import (
	"math/rand"
	"sort"
	"time"
)

func DivideRedPackage(totalAmount int, nums int) []float64 {
	return lineSegmentCutting(totalAmount*100, nums)
}

func doubleAverage(totalAmount, nums int) []float64 {
	rand.Seed(time.Now().UnixNano())
	list := make([]float64, 0)
	for nums > 1 {
		amount := rand.Intn(totalAmount/nums*2-1) + 1
		amountInFloat := float64(amount) / 100
		list = append(list, amountInFloat)
		totalAmount -= amount
		nums--
	}
	list = append(list, float64(totalAmount)/100)
	return list
}

func lineSegmentCutting(totalAmount, nums int) []float64 {
	rand.Seed(time.Now().UnixNano())
	list := make([]float64, 0)
	idxs := make([]int, 1)
	set := make(map[int]bool)
	for nums > 1 {
		amount := rand.Intn(totalAmount-1) + 1
		if _, ok := set[amount]; !ok {
			set[amount] = true
			nums--
			idxs = append(idxs, amount)
		}
	}

	sort.Ints(idxs)

	for i := 0; i < len(idxs)-1; i++ {
		amount := idxs[i+1] - idxs[i]
		totalAmount -= amount
		list = append(list, float64(amount)/100)
	}
	list = append(list, float64(totalAmount)/100)
	return list
}
