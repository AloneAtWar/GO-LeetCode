package main

/**

 * @Author: AloneAtWar

 * @Date:   2022/8/28 21:59

 * @Note:	https://leetcode.cn/problems/longest-substring-with-at-most-k-distinct-characters/
			340. Longest Substring with At Most K Distinct Characters
			340. 至多包含 K 个不同字符的最长子串
 **/

// 老题 滑动窗口
func lengthOfLongestSubstringKDistinct(s string, k int) (result int) {
	left := 0
	table := map[int]int{}
	for i := 0; i < len(s); i++ {
		table[int(s[i])]++
		if len(table) > k {
			for len(table) > k {
				table[int(s[left])]--
				if table[int(s[left])] == 0 {
					delete(table, int(s[left]))
				}
				left++
			}
		} else {
			length := i - left + 1
			if length > result {
				result = length
			}
		}
	}
	return
}
