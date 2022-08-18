package main

/**

 * @Author: AloneAtWar

 * @Date:   2022/8/18 16:57

 * @Note:	https://leetcode.cn/problems/next-greater-node-in-linked-list/
			1019. Next Greater Node In Linked List
			1019. 链表中的下一个更大节点
 **/

type ListNode struct {
	Val  int
	Next *ListNode
}

func nextLargerNodes(head *ListNode) (result []int) {
	node := flipList(head)
	var stack []int
	for node != nil {
		for len(stack) != 0 && stack[len(stack)-1] <= node.Val {
			stack = stack[:len(stack)-1]
		}
		if len(stack) == 0 {
			result = append(result, 0)
		} else {
			result = append(result, stack[len(stack)-1])
		}
		stack = append(stack, node.Val)
		node = node.Next

	}
	return flipArray(result)
}

func flipArray(arr []int) []int {
	left, right := 0, len(arr)-1
	for left < right {
		arr[left], arr[right] = arr[right], arr[left]
		left++
		right--
	}
	return arr
}

func flipList(l *ListNode) *ListNode {
	if l.Next == nil {
		return l
	}
	next := l.Next
	l.Next = nil
	node := flipList(next)
	next.Next = l
	return node
}
