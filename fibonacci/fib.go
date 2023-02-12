package fibonacci

// Fib berechnet den n-ten Wert der Fibonacci-Folge.
func Fib(n int) int {
	if n <= 1 {
		return 1
	}
	return Fib(n-1) + Fib(n-2)
}
