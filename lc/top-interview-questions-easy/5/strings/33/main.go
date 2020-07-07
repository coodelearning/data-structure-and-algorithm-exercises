package _33

func reverse(x int) int {
	if -10 < x && x < 10 {
		return x
	}
	for x%10 == 0 {
		x = x / 10
	}
	var result int
	for !(-10 < x && x < 10) {
		result = result*10 + x%10
		x = x / 10
	}
	result = result*10 + x%10
	if !(-(1<<31) <= result && result <= (1<<31)-1) {
		return 0
	}
	return result
}
