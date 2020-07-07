package main

import "log"

func main() {
	log.Println("test 14", numberOfSteps(14))
	log.Println("test 8", numberOfSteps(8))
}

func numberOfSteps(num int) int {
	var count int
	for true {
		if num == 0 {
			break
		}
		count++
		switch num % 2 {
		case 0:
			num = num / 2
		case 1:
			num--
		}
	}
	return count
}
