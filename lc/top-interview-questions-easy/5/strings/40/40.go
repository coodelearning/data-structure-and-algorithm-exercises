package _40

import (
	"fmt"
)

// 所有输入只包含小写字母 `a-z` 。

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}

	if len(strs) == 1 {
		return strs[0]
	}

	return getS(strs)
}

func getS(strs []string) string {
	var index int
	minLen := len(strs[0])
	for _, str := range strs {
		if len(str) < minLen {
			minLen = len(str)
		}
	}

	for j := 0; j < minLen; j++ {
		if index > 0 {
			break
		}
		for i := 1; i < len(strs); i++ {
			preChar := strs[0][j]
			if strs[i][j] != preChar {
				if j == 0 {
					return ""
				}
				index = j
				break
			}
			// 保证单个元素的队列完成循环
			if j == 0 && i < len(strs)-1 {
				continue
			}
			// 当判断到最短元素的长度时还没有不同的字符则可以直接返回最短长度去聚合字符
			if j == minLen-1 && i == len(strs)-1 {
				index = minLen
				break
			}
		}
	}

	var result string
	for i := 0; i < index; i++ {
		result = fmt.Sprintf("%s%s", result, string(strs[0][i]))
	}
	return result
}

func longestCommonPrefixBetter(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	str1 := []byte(strs[0])
	j := 0
	for {
		for i := 0; i < len(strs); i++ {
			// 发现不同元素，返回从开始到该元素之前一个元素的子字符串，确实比我的做法要高明很多
			if len(strs[i])-1 < j || strs[i][j] != str1[j] {
				return string(str1[:j])
			}
		}
		j++
	}
}
