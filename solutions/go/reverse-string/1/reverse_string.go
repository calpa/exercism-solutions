package reverse

func Reverse(input string) string {
	ans := ""

	for _, r := range(input) {
        ans = string(r) + ans
    }

    return ans
}
