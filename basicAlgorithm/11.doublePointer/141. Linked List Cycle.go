package main

/**

 * @Author: AloneAtWar

 * @Date:   2022/8/28 20:38

 * @Note:	https://leetcode.cn/problems/linked-list-cycle/
			141. Linked List Cycle
			141. 环形链表
 **/

func hasCycle(head *ListNode) bool {
	quick, slow := head, head
	for quick != nil {
		quick = quick.Next
		if quick == slow {
			return true
		} else if quick != nil {
			quick = quick.Next
		}
		slow = slow.Next
	}
	return false
}
