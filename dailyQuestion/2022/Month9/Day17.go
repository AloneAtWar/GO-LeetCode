package main

/**

 * @Author: AloneAtWar

 * @Date:   2022/9/17 10:35

 * @Note:	https://leetcode.cn/problems/largest-substring-between-two-equal-characters/submissions/
			1624. Largest Substring Between Two Equal Characters
			1624. 两个相同字符之间的最长子字符串
 **/

func maxLengthBetweenEqualCharacters(s string) int {
	//记录每个字符串第一个出现以及最后一次出现
	left := [26]int{}
	right := [26]int{}
	for i := 0; i < 26; i++ {
		left[i] = -1
		right[i] = -1
	}
	for i, v := range s {
		index := v - 'a'
		if left[index] == -1 {
			left[index], right[index] = i, i
		} else {
			right[index] = i
		}
	}
	result := -1
	for i := 0; i < 26; i++ {
		if left[i] != -1 {
			this := right[i] - left[i] - 1
			if this > result {
				result = this
			}
		}
	}
	return result
}
