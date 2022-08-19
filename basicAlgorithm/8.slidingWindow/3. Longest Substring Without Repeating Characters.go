package main

/**

 * @Author: AloneAtWar

 * @Date:   2022/8/19 15:31

 * @Note:	https://leetcode.cn/problems/longest-substring-without-repeating-characters/
			3. Longest Substring Without Repeating Characters
			3. 无重复字符的最长子串
 **/

func lengthOfLongestSubstring(s string) (result int) {
	left := 0
	table := make(map[byte]bool)
	for i, v := range s {
		thisByte := byte(v)
		for table[thisByte] {
			table[s[left]] = false
			left++
		}
		table[thisByte] = true
		if i-left+1 > result {
			result = i - left + 1
		}
	}
	return
}
