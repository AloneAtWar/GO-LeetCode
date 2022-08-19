package main

/**

 * @Author: AloneAtWar

 * @Date:   2022/8/19 17:52

 * @Note:	https://leetcode.cn/problems/sort-list/
			148. Sort List
			148. 排序链表
 **/

type ListNode struct {
	Val  int
	Next *ListNode
}

// 归并排序 链表
func sortList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	return mergeSort(head)
}

func mergeSort(root *ListNode) *ListNode {
	if root.Next == nil {
		return root
	}
	mid := findMid(root)
	mid.Next, mid = nil, mid.Next
	l := mergeSort(root)
	r := mergeSort(mid)
	fakeNode := &ListNode{}
	tail := fakeNode
	for l != nil && r != nil {
		if l.Val > r.Val {
			tail.Next = r
			tail = tail.Next
			next := r.Next
			r.Next = nil
			r = next
		} else {
			tail.Next = l
			tail = tail.Next
			next := l.Next
			l.Next = nil
			l = next
		}
	}
	if l != nil {
		tail.Next = l
	} else {
		tail.Next = r
	}
	return fakeNode.Next
}

func findMid(root *ListNode) *ListNode {
	if root.Next == nil {
		return root
	}
	fakeNode := &ListNode{}
	fakeNode.Next = root
	slow, quick := fakeNode, fakeNode
	for quick.Next != nil && quick.Next.Next != nil {
		slow = slow.Next
		quick = quick.Next.Next
	}
	return slow
}
