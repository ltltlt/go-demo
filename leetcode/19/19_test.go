package main

import "testing"

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	if n == 0 {
		return head
	}
	var array []*ListNode

	var current = head
	for current != nil {
		array = append(array, current)
		current = current.Next
	}
	array = append(array, (*ListNode)(nil))
	var pos = len(array) - n - 2
	if pos == -1 {
		return array[1]
	}
	array[pos].Next = array[pos+1].Next
	return head
}

func collect(head *ListNode) []int {
	var result []int
	for current := head; current != nil; current = current.Next {
		result = append(result, current.Val)
	}
	return result
}

func equalSlice(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	var i int
	for i = 0; i < len(a); i++ {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func buildList(a []int) *ListNode {
	var head, current, prev *ListNode
	for i, v := range a {
		current = &ListNode{v, nil}
		if prev != nil {
			prev.Next = current
		}
		prev = current
		if i == 0 {
			head = current
		}
	}
	return head
}

func TestRemoveNthFromEnd(t *testing.T) {
	if !equalSlice(collect(removeNthFromEnd(buildList([]int{1, 2, 3, 4, 5}), 2)), []int{1, 2, 3, 5}) {
		t.Fail()
	}
	if !equalSlice(collect(removeNthFromEnd(buildList([]int{1}), 1)), []int{}) {
		t.Fail()
	}
}
