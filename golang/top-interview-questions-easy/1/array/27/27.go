package _27

import "log"

func plusOne(digits []int) []int {
	// 空数组直接结束
	if len(digits) == 0 {
		return nil
	}
	var result []int
	// 如果为9直接返回10
	if len(digits) == 1 && digits[0] == 9 {
		return []int{1, 0}
	}
	for k, v := range digits {
		// 判断末尾是不是9
		if v == 9 && k == len(digits)-1 {
			// 是9的化开始进位操作
			goto next
		}
	}
	// 不是9的话常规末尾加一就可以返回了
	result = digits
	result[len(result)-1] = result[len(result)-1] + 1
	return result
next:
	result = digits
	log.Println("result", result)
	for i := len(result) - 1; i >= 0; i-- {
		// 末尾为9,进位
		if result[i] >= 9 && i > 0 {
			result[i-1] = result[i-1] + 1
			result[i] = 0
			// 如果进位之后高位没有到10则可以结束了
			// 若高位为10继续进位
			if result[i-1] != 10 {
				break
			}
		}
	}
	// 如果最高位为10 则添加一位,将其后的值补入结果
	if result[0] == 10 {
		result[0] = 0
		var newResult []int
		newResult = append(newResult, 1)
		newResult = append(newResult, result...)
		return newResult
	}
	return result
}
