package main

/**

 * @Author: AloneAtWar

 * @Date:   2022/8/19 15:41

 * @Note:	https://leetcode.cn/problems/longest-substring-with-at-most-k-distinct-characters/
			340. Longest Substring with At Most K Distinct Characters
			340. 至多包含 K 个不同字符的最长子串
 **/

func lengthOfLongestSubstringKDistinct(s string, k int) (result int) {
	left := 0
	table := make(map[byte]int)
	count := 0
	for i, v := range s {
		thisByte := byte(v)
		if table[thisByte] == 0 {
			count++
		}
		table[thisByte]++
		for count > k {
			table[s[left]]--
			if table[s[left]] == 0 {
				count--
			}
			left++
		}
		if i-left+1 > result {
			result = i - left + 1
		}
	}
	return
}
