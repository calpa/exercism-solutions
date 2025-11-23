package diffsquares

func SquareOfSum(n int) int {
	ans := 0
    for i := 1; i <= n; i++ {
        ans += i
    }

    return ans * ans
}

func SumOfSquares(n int) int {
	ans := 0
    for i := 1; i <= n; i++ {
        ans += i * i
    }

    return ans
}

func Difference(n int) int {
	return SquareOfSum(n) - SumOfSquares(n)
}
