package main

/**

 * @Author: AloneAtWar

 * @Date:   2022/8/18 19:06

 * @Note:	https://leetcode.cn/problems/remove-duplicate-letters/
			316. Remove Duplicate Letters
			316. 去除重复字母
 **/

// 这题真的忘记咯 ！！！！

func removeDuplicateLetters(s string) string {
	// 用来记录每个字母最后一次出现的索引
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
