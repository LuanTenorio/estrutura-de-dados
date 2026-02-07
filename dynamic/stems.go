package dynamic

import (
	"math"
	"slices"
)

func MaxStems(prices []int, n int) int {
	if n == 0 {
		return 0
	}

	var q int = math.MinInt

	for i := 1; i <= n; i++ {
		q = max(q, prices[i-1]+MaxStems(prices, n-i))
	}

	return q
}

func memoizedMaxStemsAux(prices []int, n int, mem []int) int {
	if n == 0 {
		return 0
	}

	if mem[n] != math.MinInt {
		return mem[n]
	}

	var q int = math.MinInt

	for i := 1; i <= n; i++ {
		q = max(q, prices[i-1]+memoizedMaxStemsAux(prices, n-i, mem))
	}

	mem[n] = q

	return q
}

func MemoizedMaxStems(prices []int, n int) int {
	mem := slices.Repeat([]int{math.MinInt}, n+1)

	return memoizedMaxStemsAux(prices, n, mem)
}

func BottomToTopMaxSteams(prices []int, n int) int {
	mem := slices.Repeat([]int{math.MinInt}, n+1)

	mem[0] = 0

	for i := 1; i <= n; i++ {
		q := math.MinInt
		for j := 1; j <= i; j++ {
			q = max(q, prices[j-1]+mem[i-j])
		}

		mem[i] = q
	}

	return mem[n]
}
