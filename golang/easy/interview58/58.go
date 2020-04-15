package main

import "log"

func main(){
		log.Println("test abcdefg",reverseLeftWords("abcdefg",2))
		log.Println("test lrloseumgh",reverseLeftWords("lrloseumgh",6))
}

func reverseLeftWords(s string, n int) string {
	s1:=s[:n]
	s2:=s[n:]
	s3:=s2+s1
	return s3
}