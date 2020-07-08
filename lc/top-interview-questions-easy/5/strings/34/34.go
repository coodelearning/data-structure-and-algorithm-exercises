package _34

import (
	"strings"
)

// 该版本字符串太长时耗时过长
func firstUniqChar(s string) int {
	// 切分字符串
	charS := strings.Split(s, "")
	// 边界条件
	if len(charS) == 0 {
		return -1
	}
	if len(charS) == 1 {
		return 0
	}
	//log.Println(charS)
	// 储存出现过重复的索引
	mapS := make(map[string]int)
	var repeatCharIndex []int
	for index, char := range charS {
		_, ok := mapS[char]
		if ok {
			repeatCharIndex = append(repeatCharIndex, mapS[char], index)
		} else {
			mapS[char] = index
		}
	}
	//log.Println("The repCI is :", repeatCharIndex)
	// 寻找没有出现过重复的索引，并将其返回
	for i := 0; i < len(charS); i++ {
		isRepeat := false
		for _, v := range repeatCharIndex {
			// 如果发现了重复索引，直接开始寻找下一个
			if i == v {
				isRepeat = true
				break
			}
		}
		if !isRepeat {
			return i
		}
	}
	return -1
}

func firstUniqChar2(s string) int {
	// 切分字符串
	charS := strings.Split(s, "")
	// 边界条件
	if len(charS) == 0 {
		return -1
	}
	if len(charS) == 1 {
		return 0
	}

	// 储存出现过重复的索引
	type charCount struct {
		count      uint
		firstIndex int
	}
	mapS := make(map[string]charCount)

	for index, char := range charS {
		_, ok := mapS[char]
		if ok {
			newCC := charCount{
				count:      mapS[char].count + 1,
				firstIndex: mapS[char].firstIndex,
			}
			mapS[char] = newCC
		} else {
			newCC := charCount{
				count:      1,
				firstIndex: index,
			}
			mapS[char] = newCC
		}
	}
	// map 数组化
	var firstIndexArr []int
	for _, v := range mapS {
		if v.count == 1 {
			firstIndexArr = append(firstIndexArr, v.firstIndex)
		}
	}
	// 寻找没有出现过重复的索引，并将其返回
	if len(firstIndexArr) == 0 {
		return -1
	}
	if len(firstIndexArr) == 1 {
		return firstIndexArr[0]
	}
	min := firstIndexArr[0]
	for i := 1; i < len(firstIndexArr); i++ {
		if firstIndexArr[i] > min {
			continue
		} else {
			min = firstIndexArr[i]
		}
	}

	return min
}

// 由于字母一共就26个，所以建立一个长度为26的数组
// 第一次遍历，记录每个字母最后一次出现的位置
// 第二次遍历，比较字母第一次出现的索引是不是最后一次出现的位置
// 如果是，就返回这个元素的index
// 如果不是，就在这个位置写入-1，标记这个元素已经是重复的了，不满足要求
// 如果遍历完没有找到，就返回-1
func firstUniqCharBetter(s string) int {
	last := [26]int{}

	// 记录字母最后出现的位置
	for i, v := range s {
		last[v-'a'] = i
	}

	for i, v := range s {
		if last[v-'a'] == i {
			return i
		} else {
			last[v-'a'] = -1
		}
	}

	return -1
}
