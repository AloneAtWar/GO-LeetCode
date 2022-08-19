package main

/**

 * @Author: AloneAtWar

 * @Date:   2022/8/19 15:37

 * @Note:	https://leetcode.cn/problems/longest-substring-with-at-most-two-distinct-characters/
			159. Longest Substring with At Most Two Distinct Characters
			159. 至多包含两个不同字符的最长子串
 **/

func lengthOfLongestSubstringTwoDistinct(s string) (result int) {
	left := 0
	table := make(map[byte]int)
	count := 0
	for i, v := range s {
		thisByte := byte(v)
		if table[thisByte] == 0 {
			count++
		}
		table[thisByte]++
		for count > 2 {
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
