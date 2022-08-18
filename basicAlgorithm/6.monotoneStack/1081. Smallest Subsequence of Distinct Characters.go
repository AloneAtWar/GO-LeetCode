package main

/**

 * @Author: AloneAtWar

 * @Date:   2022/8/18 19:41

 * @Note:	https://leetcode.cn/problems/smallest-subsequence-of-distinct-characters/
			1081. Smallest Subsequence of Distinct Characters
			1081. 不同字符的最小子序列
 **/

func smallestSubsequence(s string) string {
	freq := map[byte]int{}
	for i, v := range s {
		thisByte := byte(v)
		freq[thisByte] = i
	}
	table := make(map[byte]int)
	var result []byte
	for i, num := range s {
		thisByte := byte(num)
		if table[thisByte] == 0 {
			for len(result) != 0 && result[len(result)-1] > thisByte && freq[result[len(result)-1]] > i {
				table[result[len(result)-1]] = 0
				result = result[:len(result)-1]
			}
			result = append(result, thisByte)
			table[thisByte] = 1
		}
	}
	return string(result)
}
