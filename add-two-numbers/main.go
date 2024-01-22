package main

import (
	"fmt"
	"math"
)

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	// Example 1:
	// Input: l1 = [2,4,3], l2 = [5,6,4]
	// Output: [7,0,8]
	// Explanation: 342 + 465 = 807.

	// Example 2:
	// Input: l1 = [0], l2 = [0]
	// Output: [0]

	// Example 3:
	// Input: l1 = [9,9,9,9,9,9,9], l2 = [9,9,9,9]
	// Output: [8,9,9,9,0,0,0,1]

	// Example 1:
	l1 := ListNode{2, &ListNode{4, &ListNode{3, nil}}}
	l2 := ListNode{5, &ListNode{6, &ListNode{4, nil}}}
	l3 := addTwoNumbers(&l1, &l2)
	fmt.Println(
		"Value:",
		*l3,
		*l3.Next,
		*l3.Next.Next,
	)

	// Example 2:
	l1 = ListNode{0, nil}
	l2 = ListNode{0, nil}
	l3 = addTwoNumbers(&l1, &l2)
	fmt.Println("Value:", *l3)

	// Example 3:
	l1 = ListNode{9, &ListNode{9, &ListNode{9, &ListNode{9, &ListNode{9, &ListNode{9, &ListNode{9, nil}}}}}}}
	l2 = ListNode{9, &ListNode{9, &ListNode{9, &ListNode{9, nil}}}}
	l3 = addTwoNumbers(&l1, &l2)
	fmt.Println(
		"Value:",
		*l3,
		*l3.Next,
		*l3.Next.Next,
		*l3.Next.Next.Next,
		*l3.Next.Next.Next.Next,
		*l3.Next.Next.Next.Next.Next,
		*l3.Next.Next.Next.Next.Next.Next,
		*l3.Next.Next.Next.Next.Next.Next.Next,
	)
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	sum := parseListToInt(l1) + parseListToInt(l2)
	return parseIntToList(sum)
}

func parseListToInt(l *ListNode) int {
	var index, sum int
	ltemp := l
	for {
		sum += ltemp.Val * int(math.Pow(10, float64(index)))
		if ltemp.Next == nil {
			break
		}
		index++
		ltemp = ltemp.Next
	}
	return sum
}

func parseIntToList(i int) *ListNode {
	num := i
	var l, llast *ListNode
	for {
		if i == 0 {
			l = &ListNode{0, nil}
			break
		} else if num == 0 {
			break
		} else {
			numTemp := num % 10
			if l == nil {
				l = &ListNode{numTemp, nil}
				llast = l
			} else if llast.Next == nil {
				llast.Next = &ListNode{numTemp, nil}
				llast = llast.Next
			} else {
				llast.Next = &ListNode{numTemp, nil}
			}
			num -= numTemp
			num /= 10
		}
	}
	return l
}
