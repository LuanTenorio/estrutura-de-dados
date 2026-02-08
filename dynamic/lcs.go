package dynamic

func LcsLength(X string, Y string) [][]int {
	m := len(X)
	n := len(Y)

	c := make([][]int, m+1)
	for i := range c {
		c[i] = make([]int, n+1)
	}

	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if X[i-1] == Y[j-1] {
				c[i][j] = c[i-1][j-1] + 1
			} else {
				c[i][j] = max(c[i-1][j], c[i][j-1])
			}
		}
	}

	return c
}
