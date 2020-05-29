package _27

func plusOne(digits []int) []int {
	if len(digits) == 0 {
		return nil
	}
	var result []int
	for _, v := range digits {
		if v == 9 {
			goto next
		}
	}
	result = digits
	result[len(digits)-1] = result[len(digits)-1] + 1
	return result
next:
	for i := len(digits) - 1; i >= 0; i-- {

	}
	return nil
}
