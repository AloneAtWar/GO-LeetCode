package main

/**

 * @Author: AloneAtWar

 * @Date:   2022/9/1 15:29

 * @Note:	https://leetcode.cn/problems/alien-dictionary/
			269. Alien Dictionary
			269. 火星词典
 **/

func alienOrder(words []string) string {
	graph := make([][]int, 26)
	n := len(words)
	inDegree := make([]int, 26)
	table := make(map[int]bool)
	for _, word := range words {
		for _, v := range word {
			table[int(v-'a')] = true
		}
	}
	for i := 0; i < n-1; i++ {
		word1, word2 := words[i], words[i+1]
		if word1 == word2 {
			continue
		}
		n, m := len(word1), len(word2)
		i, j := 0, 0
		for i < n && j < m && word1[i] == word2[j] {
			i++
			j++
		}
		if j >= m {
			return ""
		}
		if i >= n {
			continue
		}
		char1, char2 := int(word1[i]-'a'), int(word2[j]-'a')
		graph[char1] = append(graph[char1], char2)
		inDegree[char2]++
	}
	var list []int
	var out []byte
	for i, v := range inDegree {
		if table[i] && v == 0 {
			list = append(list, i)
		}
	}
	for len(list) != 0 {
		this := list[0]
		list = list[1:]
		out = append(out, byte(this+'a'))
		for _, i := range graph[this] {
			inDegree[i]--
			if inDegree[i] == 0 {
				list = append(list, i)
			}
		}
	}
	if len(out) != len(table) {
		return ""
	} else {
		return string(out)
	}
}
