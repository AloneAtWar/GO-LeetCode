package main

import "unicode"

/**

 * @Author: AloneAtWar

 * @Date:   2022/8/19 23:10

 * @Note:	https://leetcode.cn/problems/valid-word-abbreviation/
			408. Valid Word Abbreviation
			408. 有效单词缩写
 **/

func validWordAbbreviation(word string, abbr string) bool {
	m, n := len(word), len(abbr)
	i, j := 0, 0
	for i < m && j < n {
		if unicode.IsDigit(rune(abbr[j])) {
			if abbr[j] == '0' {
				return false
			}
			num := 0
			for j < n && unicode.IsDigit(rune(abbr[j])) {
				num = num*10 + int(abbr[j]-'0')
				j++
			}
			i += num
		} else if word[i] == abbr[j] {
			j++
			num := 0
			i += num + 1
		} else {
			return false
		}
	}
	if i != m || j != n {
		return false
	}
	return true
}
