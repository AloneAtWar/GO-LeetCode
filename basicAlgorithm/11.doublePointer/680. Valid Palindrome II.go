package main

/**

 * @Author: AloneAtWar

 * @Date:   2022/8/19 23:04

 * @Note:	https://leetcode.cn/problems/valid-palindrome-ii/
			680. Valid Palindrome II
			680. 验证回文字符串 Ⅱ
 **/

func validPalindrome(s string) bool {
	return valid(s, 0, len(s)-1, true)
}

func valid(s string, left, right int, canDel bool) bool {
	for left < right {
		if s[left] != s[right] {
			return canDel && (valid(s, left, right-1, false) || valid(s, left+1, right, false))
		}
		left++
		right--
	}
	return true
}
