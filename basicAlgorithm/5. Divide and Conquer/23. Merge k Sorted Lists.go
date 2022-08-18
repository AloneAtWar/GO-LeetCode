package main

/**

 * @Author: AloneAtWar

 * @Date:   2022/8/18 15:07

 * @Note:	https://leetcode.cn/problems/merge-k-sorted-lists/
			23. Merge k Sorted Lists
			23. 合并K个升序链表
 **/

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeKLists(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}
	for len(lists) != 1 {
		out1, out2 := lists[0], lists[1]
		lists = append(lists[2:], merge(out1, out2))
	}
	return lists[0]
}

func merge(list1 *ListNode, list2 *ListNode) *ListNode {
	fakeHead := &ListNode{}
	pre := fakeHead
	for list1 != nil && list2 != nil {
		if list1.Val < list2.Val {
			pre.Next = &ListNode{Val: list1.Val}
			list1 = list1.Next
		} else {
			pre.Next = &ListNode{Val: list2.Val}
			list2 = list2.Next
		}
		pre = pre.Next
	}
	if list1 != nil {
		pre.Next = list1
	}
	if list2 != nil {
		pre.Next = list2
	}
	return fakeHead.Next
}
