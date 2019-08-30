package main

import "sort"

func qs(array []int, i, j int) {
	if i >= j {
		return
	}
	key := array[i]
	left, right := i, j

	for left < right {
		for left < right && array[right] >= key {
			right--
		}
		array[left] = array[right]
		for left < right && array[left] <= key {
			left++
		}
		array[right] = array[left]
	}

	array[left] = key

	qs(array, i, left-1)
	qs(array, left+1, j)
}

func bisect(array []int, elem int) int {
	left, right := 0, len(array)-1

	for left <= right {
		mid := (right + left) / 2
		if array[mid] > elem {
			right = mid - 1
		} else if array[mid] < elem {
			left = mid + 1
		} else {
			return mid
		}
	}
	return -1
}

func main() {
	array := []int{1, 2, -1, 3, 8, 9, 0}
	var array1 []int

	for _, elem := range array {
		array1 = append(array1, elem)
	}

	sort.Ints(array1)
	qs(array, 0, len(array)-1)

	for i := range array {
		if array[i] != array1[i] {
			panic("not equal")
		}
	}
}
