package dynamic

func LcsLength(X string, Y string) ([][]int, [][]rune) {
	m, n := len(X), len(Y)
	c, b := make([][]int, m+1), make([][]rune, m+1)

	for i := range c {
		c[i], b[i] = make([]int, n+1), make([]rune, n+1)
	}

	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if X[i-1] == Y[j-1] {
				c[i][j] = c[i-1][j-1] + 1
				b[i][j] = '↖'
			} else if c[i-1][j] >= c[i][j-1] {
				c[i][j] = c[i-1][j]
				b[i][j] = '↑'
			} else {
				c[i][j] = c[i][j-1]
				b[i][j] = '←'
			}
		}
	}

	return c, b
}
