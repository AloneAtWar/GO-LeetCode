package main

import "strings"

/**

 * @Author: AloneAtWar

 * @Date:   2022/8/21 14:46

 * @Note:	https://leetcode.cn/problems/check-if-a-word-occurs-as-a-prefix-of-any-word-in-a-sentence/
			1455. Check If a Word Occurs As a Prefix of Any Word in a Sentence
			1455. 检查单词是否为句中其他单词的前缀
 **/

//API 调用大师
func isPrefixOfWord(sentence string, searchWord string) int {
	strs := strings.Split(sentence, " ")
	for i, str := range strs {
		if strings.HasPrefix(str, searchWord) {
			return i
		}
	}
	return -1
}
