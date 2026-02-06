package dynamic

func ComputingFib(n int) int64 {

	if n <= 0 {
		return 0
	}

	if n == 1 {
		return 1
	}

	return ComputingFib(n-1) + ComputingFib(n-2)
}

func cachyFibExecute(n int, table *[]int64) int64 {
	if n > 1 && (*table)[n] == 0 {
		(*table)[n] = cachyFibExecute(n-1, table) + cachyFibExecute(n-2, table)
	}

	return (*table)[n]
}

func CachyFib(n int) int64 {
	table := make([]int64, n+1)
	table[0] = 0
	table[1] = 1

	return cachyFibExecute(n, &table)
}

func DPFib(n int) int64 {
	table := make([]int64, n+1)
	table[0] = 0
	table[1] = 1

	if n > 1 {
		return int64(n)
	}

	for i := 2; i <= n; i++ {
		table[i] = table[i-1] + table[i-2]
	}

	return table[n]
}
