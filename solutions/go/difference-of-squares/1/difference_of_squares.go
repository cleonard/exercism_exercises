package differenceofsquares

import "math"

func SquareOfSum(n int) int {
	sum := 0
	for i := 1; i <= n; i++ {
		sum += i
	}
	result := math.Pow(float64(sum), float64(2))
	return int(result)
}

func SumOfSquares(n int) int {
	sum := 0
	for i := 1; i <= n; i++ {
		sum += int(math.Pow(float64(i), float64(2)))
	}
	return sum
}

func Difference(n int) int {
	return SquareOfSum(n) - SumOfSquares(n)
}
