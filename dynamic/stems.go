package dynamic

import (
	"math"
)

func MaxStems(prices []int, n int) int {
	if n == 0 {
		return 0
	}

	var q int = math.MinInt

	for i := range n {
		q = max(q, prices[i]+MaxStems(prices, n-(i+1)))
	}

	return q
}
