package main

/**

 * @Author: AloneAtWar

 * @Date:   2022/8/19 16:06

 * @Note:	https://leetcode.cn/problems/minimum-window-substring/
			76. Minimum Window Substring
			76. 最小覆盖子串
 **/

func minWindow(s string, t string) string {
	n := len(s)
	need := [52]int{}
	for _, v := range t {
		need[findIndex(int(v))]++
	}
	curr := [52]int{}
	real := [52]int{}
	left := 0
	result := ""
	for i := 0; i < n; i++ {
		index := findIndex(int(s[i]))
		real[index]++
		if curr[index] < need[index] {
			curr[index]++
		}
		for curr == need {
			length := i - left + 1
			if len(result) == 0 || len(result) > length {
				result = s[left : i+1]
			}
			delIndex := findIndex(int(s[left]))
			real[delIndex]--
			if real[delIndex] < need[delIndex] {
				curr[delIndex]--
			}
			left++
		}
	}

	return result
}

func findIndex(v int) int {
	currIndex := 0
	if v >= 'a' {
		currIndex = int(v - 'a')
	} else {
		currIndex = int(v - 'A' + 26)
	}
	return currIndex
}

func main() {
	minWindow("ADOBECODEBANC", "ABC")
}
