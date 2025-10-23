package logicaltest

func Fibonacci(n int) []int {
	var result []int = []int{0, 1}
	if n < 2 {
		return []int{}
	}

	for i := 2; i <= n+1; i++ {
		newVal := result[i-1] + result[i-2]
		result = append(result, newVal)
	}
	return result

}
