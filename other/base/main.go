package main

import (
	"fmt"
)

func main() {
	a := []int{1, 3, 2, 4, 5}
	fmt.Println(QuickSort(a))
}

func quickSort(values []int, left int, right int) {
	key := values[left] //取出第一项
	p := left
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

		if values[i] <= key && i <= p {
			i++
		}

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
	//quickSort(values, 0, len(values)-1)
	quickSort2(values, 0, len(values)-1)
	return values
}

func quickSort2(source []int, leftIndex, rightIndex int) []int {
	if leftIndex >= rightIndex {
		return source
	}
	for i := leftIndex; i < rightIndex; {
		if source[leftIndex] < source[len(source)-1] {
			leftIndex++
		}
		if leftIndex >= rightIndex {
			source[leftIndex], source[len(source)-1] = source[len(source)-1], source[leftIndex]
			fmt.Println("the source is :", source)
			leftSource := source[0 : leftIndex-1]
			rightSource := source[leftIndex : len(source)-1]
			leftResult := quickSort2(leftSource, 0, 0)
			rightResult := quickSort2(rightSource, 0, 0)
			fmt.Println("the left is :", leftResult)
			fmt.Println("the right is :", rightResult)
			var result []int
			result = append(leftResult, rightResult...)
			return result
		}
		if source[rightIndex] > source[len(source)-1] {
			rightIndex--
		}
		i++
	}
	return nil
}
