package main

/**

 * @Author: AloneAtWar

 * @Date:   2022/8/19 15:42

 * @Note:	https://leetcode.cn/problems/longest-repeating-character-replacement/
			424. Longest Repeating Character Replacement
			424. 替换后的最长重复字符
 **/
func characterReplacement(s string, k int) (result int) {
	n := len(s)
	count := make([]int, 26)
	left := 0
	for i := 0; i < n; i++ {
		count[s[i]-'A']++
		for i-left+1-findMax(count) > k {
			count[s[left]-'A']--
			left++
		}
		if i-left+1 > result {
			result = i - left + 1
		}
	}
	return
}

func findMax(ints []int) int {
	result := 0
	for i := 1; i < len(ints); i++ {
		if ints[result] < ints[i] {
			i = result
		}
	}
	return ints[result]
}

// 前面那个会超时 下面这个思路更加巧妙 假设26种可能性
func characterReplacement(s string, k int) (result int) {
	for i := 0; i < 26; i++ {
		left := 0
		currK := k
		for j, v := range s {
			if int(v) != 'A'+i {
				if currK > 0 {
					currK--
				} else {
					for int(s[left]) == 'A'+i && left != j {
						left++
					}
					// 为什么这里 currK 没有增加的原因是 增加了一个非假设元素 但也减少了一个非假设元素
					//if left!=j{
					//	currK++
					//}
					left++
				}
			}
			if result < j-left+1 {
				result = j - left + 1
			}
		}
	}
	return
}
