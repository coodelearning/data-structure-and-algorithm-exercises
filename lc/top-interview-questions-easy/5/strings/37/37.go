package _37

import (
	"math"
	"strings"
)

func myAtoi(str string) int {
	// 从第一个非空格字符开始
	var startChar int32
	// 如果首字符不是数字或者负号的话就抛出错误
	for _, v := range str {
		if ('0' <= v && v <= '9') || v == '-' || v == '+' {
			startChar = v
			break
		}
		if v != ' ' && !('0' < v && v < '9') && v != '-' && v != '+' {
			return 0
		}
	}
	var result int64
	var intChars []int32
	var intStart bool = false
	for k, v := range str {
		if '0' <= v && v <= '9' {
			intStart = true
			intChars = append(intChars, v)
			continue
		}
		if k == len(str)-1 {
			break
		}
		if intStart {
			break
		}
		if v == ' ' || ((v == '-' || v == '+') && '0' <= str[k+1] && str[k+1] <= '9') {
			continue
		} else {
			break
		}
	}
	for i := 0; i < len(intChars); i++ {
		var sqrt int64 = 1
		for j := 1; j < len(intChars)-i; j++ {
			sqrt *= 10
		}
		//if result > 0 && math.MaxInt32/result < sqrt {
		//	if startChar == '-' {
		//		return math.MinInt32
		//	}
		//	return math.MaxInt32
		//}
		result += (int64(intChars[i]) - 48) * sqrt
	}
	// 判断正负
	if startChar == '-' && result > 0 {
		result = -result
	}
	// 处理越界
	if startChar != '-' && result < 0 {
		return math.MaxInt32
	}
	if result >= math.MaxInt32 {
		return math.MaxInt32
	}
	if result <= math.MinInt32 {
		return math.MinInt32
	}
	return int(result)
}

func myAtoiBetter(str string) int {
	return convertForBetter(cleanForBetter(str))
}

func cleanForBetter(s string) (sign int, abs string) {
	// 先去除首尾空格
	s = strings.TrimSpace(s)
	if s == "" {
		return
	}
	// 判断第一个字符
	switch s[0] {
	// 有效的
	case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
		sign, abs = 1, s
	// 有效的，正号
	case '+':
		sign, abs = 1, s[1:]
	// 有效的，负号
	case '-':
		sign, abs = -1, s[1:]
	// 无效的，当空字符处理，并且直接返回
	default:
		abs = ""
		return
	}
	for i, b := range abs {
		// 遍历第一波处理过的字符，如果直到第i个位置有效，那就取s[:i]，从头到这个有效的字符，剩下的就不管了，也就是break掉
		// 比如 s=123abc，那么就取123，也就是s[:3]
		if b < '0' || '9' < b {
			abs = abs[:i]
			// 一定要break，因为后面的就没用了
			break
		}
	}
	return
}

// 接收的输入是已经处理过的纯数字
func convertForBetter(sign int, absStr string) int {
	absNum := 0
	for _, b := range absStr {
		// b - '0' ==> 得到这个字符类型的数字的真实数值的绝对值
		absNum = absNum*10 + int(b-'0')
		// 检查溢出
		switch {
		case sign == 1 && absNum > math.MaxInt32:
			return math.MaxInt32
		// 这里和正数不一样的是，必须和负号相乘，也就是变成负数，否则永远走不到里面
		case sign == -1 && absNum*sign < math.MinInt32:
			return math.MinInt32
		}
	}
	return sign * absNum
}
