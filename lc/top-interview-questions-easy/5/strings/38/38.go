package _38

import "strings"

// 先暴力解决再考虑优化

func strStr(haystack string, needle string) int {
	if len(needle) == 0 {
		return 0
	}
	if len(needle) > len(haystack) {
		return -1
	}
	// 如果不存在返回-1
	for kn, n := range needle {
		var isExist bool = false
		for kh, h := range haystack {
			// 存在，寻找位置
			if h == n {
				// 如果子串只有一个字符，此时就可以返回正确的结果了
				isExist = true
				// 试图寻找子串的位置
				if len(needle) == 1 {
					return kh
				} else {
					isExist = false
					// 如果在一个对上之后的每一个都能对上，则完成了子串的查找
					for i, j := kn, kh; i < len(needle) && j < len(haystack); i, j = i+1, j+1 {
						if needle[i] != haystack[j] {
							break
						}
						if i == len(needle)-1 {
							isExist = true
						}
					}
					// 此时找到了完美匹配的子串
					if isExist {
						return kh
					}
				}
			}
		}
		// 没有找到匹配的位置
		if !isExist {
			return -1
		}
	}

	return 0
}

func strStrBetter(haystack string, needle string) int {
	if needle == "" {
		return 0
	}
	if len(haystack) < len(needle) {
		return -1
	}
	for i := 0; i <= len(haystack)-len(needle); i++ {
		// Ps.等一下，这个犯规了吧？？？
		if strings.HasPrefix(haystack[i:], needle) {
			return i
		}
	}
	return -1
}
