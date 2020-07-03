package main

import (
	"leetcode/other/base/sort"
	"log"
)

func main() {
	values := []int{7, 3, 2, 5, 4}
	//log.Println(QuickSort(values))
	log.Println(sort.QuickSort(values))
	//log.Println(quickSort2(values, 0, len(values)-1))
}

func quickSort(values []int, left int, right int) {
	key := values[left] // 取出第一项
	p := left           // 用第一个元素做基准
	i, j := left, right

	for i <= j {
		//由后开始向前搜索(j--)，找到第一个小于key的值values[j]
		for j >= p && values[j] >= key {
			j--
		}

		//第一个小于key的值 赋给 values[p]
		if j >= p {
			values[p] = values[j]
			p = j
		}

		// Ps.从前往后搜索大于key的值
		if values[i] <= key && i <= p {
			i++
		}

		// Ps.第一个大于key的值 赋给 values[p]
		if i < p {
			values[p] = values[i]
			p = i
		}

		values[p] = key
		if p-left > 1 {
			quickSort(values, left, p-1)
		}
		if right-p > 1 {
			quickSort(values, p+1, right)
		}
	}

}

func QuickSort(values []int) []int {
	quickSort(values, 0, len(values)-1)
	//quickSort2(values, 0, len(values)-1)
	return values
}

// 这是一个S(N)的版本
func quickSort2(source []int, left, right int) []int {
	leftIndex := left
	rightIndex := right - 1
	log.Println("the source is :", source)

	if leftIndex >= rightIndex {
		return source
	}

	var isChanged = false

	for !isChanged {
		for i := left; i < right; {
			if source[leftIndex] < source[rightIndex] && leftIndex < rightIndex {
				leftIndex++
			} else if leftIndex >= rightIndex {
				isChanged = true
				source[leftIndex], source[right] = source[right], source[leftIndex]
			}
			i++
		}

		if !isChanged {
			for i := left; i < right; {
				if source[rightIndex] > source[right] && leftIndex < rightIndex {
					rightIndex--
				} else if leftIndex >= rightIndex {
					source[leftIndex], source[right] = source[right], source[leftIndex]
				}
				i++
			}
		}

		if leftIndex >= rightIndex {
			source[leftIndex], source[rightIndex] = source[rightIndex], source[leftIndex]
			leftSource := source[0:leftIndex]
			leftResult := quickSort2(leftSource, 0, len(leftSource)-1)
			var rightSource []int
			var rightResult []int
			if right-leftIndex > 1 {
				rightSource = source[leftIndex:right]
				rightResult = quickSort2(rightSource, 0, len(rightSource)-1)
			} else {
				rightResult = append(rightResult, source[right])
			}
			log.Println("the left is :", leftSource)
			log.Println("the right is :", rightSource)
			log.Println("the left is :", leftResult)
			log.Println("the right is :", rightResult)
			var result []int
			result = append(leftResult, source[leftIndex])
			result = append(leftResult, rightResult...)
			return result
		}

	}

	log.Println("return with nil")
	return nil
}
