package main

/**

 * @Author: AloneAtWar

 * @Date:   2022/8/28 20:37

 * @Note:

 **/

type ListNode struct {
	Val  int
	Next *ListNode
}

func detectCycle(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	quick, slow := head, head
	for quick.Next != nil && quick.Next.Next != nil {
		quick = quick.Next.Next
		slow = slow.Next
		if quick == slow {
			slow = head
			for slow != quick {
				quick = quick.Next
				slow = slow.Next
			}
			return slow
		}
	}
	return nil
}
