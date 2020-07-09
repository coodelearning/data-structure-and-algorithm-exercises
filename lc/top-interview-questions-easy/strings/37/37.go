package _37

import (
	"log"
	"math"
	"strings"
)

// 需要注意的地方在于每一次累加都要判断是否有发生越界
// 如果只在最后才进行判断可能会错过一些很大的数据
// 但是好像每一次判断也不能自信的说每次就一定可以避免越界？

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
	// 按顺序提取数字
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
	// 将整形数组元素累加成整数
	var result int
	for i := 0; i < len(intChars); i++ {
		result = result*10 + int(intChars[i]-'0')
		// 处理越界
		switch {
		case startChar != '-' && result > math.MaxInt32:
			return math.MaxInt32
		case startChar == '-' && -result < math.MinInt32:
			return math.MinInt32
		}
	}
	// 判断正负
	if startChar == '-' && result > 0 {
		result = -result
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
	log.Println("The param is :", sign, absStr)
	absNum := 0
	for _, b := range absStr {
		// b - '0' ==> 得到这个字符类型的数字的真实数值的绝对值
		absNum = absNum*10 + int(b-'0')
		log.Println("The abs is ", absNum)
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
