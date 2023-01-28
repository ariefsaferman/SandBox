package fibo

func RecursiveFibonacci(n int) int {
	if n <= 1 {
		return n
	}

	return RecursiveFibonacci(n-1) + RecursiveFibonacci(n-2)
}

func SequentialFibonacci(n int) int {
	if n <= 1 {
		return int(n)
	}

	var n2, n1 int = 0, 1
	for i := int(2); i < n; i++ {
		n2, n1 = n1, n1+n2
	}

	return n2 + n1
}
