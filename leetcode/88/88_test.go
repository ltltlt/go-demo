package main

import "testing"

func merge(nums1 []int, m int, nums2 []int, n int) {
	var lastIndex = m + n - 1
	begin1, begin2 := m-1, n-1
	for begin1 >= 0 || begin2 >= 0 {
		for begin1 >= 0 && (begin2 < 0 || nums1[begin1] >= nums2[begin2]) {
			nums1[lastIndex] = nums1[begin1]
			lastIndex--
			begin1--
		}
		for begin2 >= 0 && (begin1 < 0 || nums1[begin1] < nums2[begin2]) {
			nums1[lastIndex] = nums2[begin2]
			lastIndex--
			begin2--
		}
	}
}

func equalSlice(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
func TestMerge(t *testing.T) {
	var slice1 = []int{1, 2, 3, 0, 0, 0}
	var slice2 = []int{2, 5, 6}
	merge(slice1, 3, slice2, 3)
	if !equalSlice(slice1, []int{1, 2, 2, 3, 5, 6}) {
		t.Fail()
	}
	slice1 = []int{2, 0}
	slice2 = []int{1}
	merge(slice1, 1, slice2, 1)
	if !equalSlice(slice1, []int{1, 2}) {
		t.Fail()
	}
}
