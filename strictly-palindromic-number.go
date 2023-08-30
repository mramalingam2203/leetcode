// https://leetcode.com/problems/strictly-palindromic-number/

func isStrictlyPalindromic(n int) bool {

}

func main() {
	decimalToNnary(10, 2)
}

func decimalToNnary(decimal int, base int) string {
	nnnary := ""

	if decimal == 0 {
		return "0"
	}

	for decimal > 0 {
		remainder := decimal % base
		nnary = strconv.Itoa(remainder) + nnary
		decimal /= 2
	}

	return nnary
}

