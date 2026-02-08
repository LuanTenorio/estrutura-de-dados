package dynamic

import "fmt"

func LcsLength(X, Y string) ([][]int, [][]rune) {
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

func PrintLcs(table [][]rune, X string, i, j int) {
	if i == 0 || j == 0 {
		return
	}

	switch table[i][j] {
	case '↖':
		PrintLcs(table, X, i-1, j-1)
		fmt.Printf("%c", X[i-1])
	case '←':
		PrintLcs(table, X, i, j-1)
	default:
		PrintLcs(table, X, i-1, j)
	}

	if i == len(X) {
		fmt.Println()
	}
}

func PrintArrowsTable(b [][]rune, X, Y string) {
	fmt.Print("      ∅ ")

	for _, char := range Y {
		fmt.Printf(" %c ", char)
	}

	fmt.Println()

	for i := 0; i < len(b); i++ {
		if i == 0 {
			fmt.Print(" ∅ | ")
		} else {
			fmt.Printf(" %c | ", X[i-1])
		}

		for j := 0; j < len(b[i]); j++ {
			val := b[i][j]
			if val == 0 {
				fmt.Print(" . ")
			} else {
				fmt.Printf(" %c ", val)
			}
		}

		fmt.Println()
	}
}
