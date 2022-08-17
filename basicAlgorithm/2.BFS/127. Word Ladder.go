package main

import "fmt"

/**

 * @Author: AloneAtWar

 * @Date:   2022/8/17 19:42

 * @Note:	https://leetcode.cn/problems/word-ladder/
			127. Word Ladder
			127. 单词接龙
 **/

func ladderLength(beginWord string, endWord string, wordList []string) int {
	if beginWord == endWord {
		return 0
	}
	table := map[string]bool{}
	for _, s := range wordList {
		table[s] = true
	}
	if !table[endWord] {
		return 0
	}
	list := []string{beginWord}
	var next []string
	step := 0
	visited := map[string]bool{beginWord: true}
	for len(list) != 0 {
		for _, out := range list {
			for i, _ := range out {
				for j := 0; j < 26; j++ {
					change := fmt.Sprintf("%s%c%s", out[:i], 'a'+j, out[i+1:])
					if table[change] && !visited[change] {
						visited[change] = true
						next = append(next, change)
						if change == endWord {
							return step
						}
					}
				}
			}
		}
		step++
		list = next
		next = []string{}
	}

	return 0
}
