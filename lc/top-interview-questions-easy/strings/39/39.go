package _39

import (
	"fmt"
	"strconv"
)

// 可能要用递归

func countAndSay(n int) string {
	var tmp, result string
	switch {
	case n == 1:
		return "1"
	case n == 2:
		return "11"
	}

	result = countAndSay(n - 1)

	var count = 1
	for i := 1; i < len(result); i++ {
		if result[i] != result[i-1] {
			tmp += fmt.Sprintf("%s%s", strconv.Itoa(count), string(result[i-1]))
			count = 1
		} else {
			count++
		}
		if i == len(result)-1 {
			tmp += fmt.Sprintf("%s%s", strconv.Itoa(count), string(result[i]))
			result = tmp
			return result
		}
	}

	return result
}

func countAndSayBetter(n int) string {
	var r, s string
	var count int
	if n == 1 {
		return "1"
	}
	if n == 2 {
		return "11"
	}

	s = countAndSayBetter(n - 1)
	//初始化字符计数器
	count = 1
	//1. 对S从左到右遍历
	for i := 1; i <= len(s)-1; i++ {
		//1.1 如果s[i] != s[i-1] 则把前面的字符打印
		if s[i] != s[i-1] {
			r += fmt.Sprintf("%s%s", strconv.Itoa(count), string(s[i-1]))
			count = 1
		} else {
			//1.2 如果相等则计数器+1
			count++
		}
		//1.2 如果到末尾了，则打印当前字符
		if i+1 == len(s) {
			r += fmt.Sprintf("%s%s", strconv.Itoa(count), string(s[i]))
			s = r
			return s
		}
	}
	return s
}

func countAndSay2(n int) string {
	if n < 1 {
		return ""
	}

	ret := "1"
	for i := 2; i <= n; i++ {
		ret = getFor2(ret)
	}

	return ret
}

func getFor2(input string) string {
	curr := input[0:1]
	cnt := 1
	ret := ""
	for i := 1; i < len(input); i++ {
		t := input[i : i+1]
		if t == curr {
			cnt++
		} else {
			ret += fmt.Sprintf("%d%s", cnt, curr)
			curr = t
			cnt = 1
		}
	}

	ret += fmt.Sprintf("%d%s", cnt, curr)

	return ret
}
