package main

import (
	"log"
)

// 编写一个函数，其作用是将输入的字符串反转过来。输入字符串以字符数组 char[] 的形式给出。

// 不要给另外的数组分配额外的空间，你必须原地修改输入数组、使用 O (1) 的额外空间解决这一问题。

// 你可以假设数组中的所有字符都是 ASCII 码表中的可打印字符。(是不是一个提示)

func main() {
	ss := "hello"
	log.Println(reverseString([]byte(ss)))
}

func reverseString(s []byte) string {
	j := len(s) - 1
	switch len(s) % 2 {
	case 0:
		for i := 0; i < len(s)/2; i++ {
			tmp := s[i]
			s[i] = s[j]
			s[j] = tmp
			j--
		}
	case 1:
		for i := 0; i < (len(s)-1)/2; i++ {
			tmp := s[i]
			s[i] = s[j]
			s[j] = tmp
			j--
		}
	}
	str2 := string(s[:])
	return str2
}
