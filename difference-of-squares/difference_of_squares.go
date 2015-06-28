package diffsquares

// SumOfSquares returns the sum of the
// squares of the first n natural numbers
func SumOfSquares(n int) int {
	total := 0
	for i := 1; i <= n; i++ {
		total += i * i
	}

	return total
}

// SquareOfSums returns the square of the
// sum of the first n natural numbers
func SquareOfSums(n int) int {
	total := 0
	for i := 1; i <= n; i++ {
		total += i
	}

	return total * total
}

// Difference returns the difference between
// the square of the sums and the sum of the
// squares of the first n natural numbers
func Difference(n int) int {
	return SquareOfSums(n) - SumOfSquares(n)
}
