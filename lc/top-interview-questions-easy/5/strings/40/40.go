package _40

import (
	"fmt"
	"log"
)

// 所有输入只包含小写字母 `a-z` 。

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}

	if len(strs) == 1 {
		return strs[0]
	}

	var index int

	for j := 0; j < len(strs[0]); j++ {
		if index != 0 {
			break
		}
		for i := 1; i < len(strs); i++ {
			if j > len(strs[i])-1 {
				index = j - 1
				break
			}
			preChar := strs[0][j]
			// TODO:缺少了兼容完全匹配的逻辑
			if j < len(strs[i]) && strs[i][j] != preChar {
				if j == 0 && len(strs[i]) != 1 {
					log.Println("Im return in here 1 and len(strs[i]) =", len(strs[i]))
					return ""
				}
				index = j
				break
			}
			if j == len(strs[i])-1 {
				index = len(strs[0])
				break
			}
		}
	}

	if index < 0 {
		log.Println("Im return in here 2")
		return ""
	}

	var result string
	for i := 0; i < index; i++ {
		result = fmt.Sprintf("%s%s", result, string(strs[0][i]))
	}
	return result
}
