package _34

import (
	"log"
	"strings"
)

func firstUniqChar(s string) int {
	charS := strings.Split(s, "")
	if len(charS) == 0 {
		return -1
	}
	if len(charS) == 1 {
		return 0
	}
	log.Println(charS)
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
	log.Println("The repCI is :", repeatCharIndex)
	for i := 0; i < len(charS); i++ {
		isRepat := false
		for _, v := range repeatCharIndex {
			if i == v {
				isRepat = true
				break
			}
		}
		if !isRepat {
			return i
		}
	}
	return -1
}
