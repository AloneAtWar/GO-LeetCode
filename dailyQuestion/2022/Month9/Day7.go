package main

import "strings"

/**

 * @Author: AloneAtWar

 * @Date:   2022/9/11 12:33

 * @Note:	https://leetcode.cn/problems/rearrange-spaces-between-words/
			1592. Rearrange Spaces Between Words
			1592. 重新排列单词间的空格
 **/

func reorderSpaces(s string) (ans string) {
	words := strings.Fields(s)
	space := strings.Count(s, " ")
	lw := len(words) - 1
	if lw == 0 {
		return words[0] + strings.Repeat(" ", space)
	}
	return strings.Join(words, strings.Repeat(" ", space/lw)) + strings.Repeat(" ", space%lw)
}
