package _6

import (
	"strings"
)

// 可以忽略大小写 故而可以考虑一下在ascii码中固定的间隔关系
// 问题之一体现在空格是该如何处理的

func isPalindrome(s string) bool {
	if len(s) <= 1 {
		return true
	}
	// 处理字符串，只保留字母和数字
	var charStrings []string
	for _, v := range s {
		if (65 <= v && v <= 90) || (97 <= v && v <= 122) || (48 <= v && v <= 57) {
			charStrings = append(charStrings, string(v))
		}
	}

	var checkString string
	checkString = strings.Join(charStrings, checkString)

	newS := checkString
	sLen := len(checkString)
	for i, j := 0, sLen-1; i < sLen/2; i, j = i+1, j-1 {
		// 数字必须相同
		if 48 <= checkString[i] && checkString[i] <= 57 && checkString[i] != newS[j] {
			return false
		}
		// 字母考虑大小写
		if checkString[i] != newS[j] && checkString[i] != newS[j]-32 && checkString[i] != newS[j]+32 {
			return false
		}
	}
	return true
}

// 时间优势体现在不用将字符串来回进行处理
// 顺便学习了strings包中的ToLower用法
func isPalindromeBetter(s string) bool {
	s = strings.ToLower(s)
	left := 0
	right := len(s) - 1

	for left < right {
		for left < len(s) && !isalnumForBetter(s[left]) {
			left++
		}
		for right >= 0 && !isalnumForBetter(s[right]) {
			right--
		}
		if left < right {
			if s[left] == s[right] {
				left++
				right--
			} else {
				return false
			}
		}
	}
	return true
}

func isalnumForBetter(ch byte) bool {
	return (ch >= 'A' && ch <= 'Z') || (ch >= 'a' && ch <= 'z') || (ch >= '0' && ch <= '9')
}
