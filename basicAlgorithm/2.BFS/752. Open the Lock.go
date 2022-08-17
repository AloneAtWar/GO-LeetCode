package main

import "fmt"

/**

 * @Author: AloneAtWar

 * @Date:   2022/8/17 19:55

 * @Note:	https://leetcode.cn/problems/open-the-lock/
			752. Open the Lock
			752. 打开转盘锁
 **/

func openLock(deadends []string, target string) int {
	if target == "0000" {
		return 0
	}
	table := map[string]bool{}
	for _, deadend := range deadends {
		table[deadend] = true
	}
	if table[target] || table["0000"] {
		return -1
	}
	visited := map[string]bool{"0000": true}
	step := 0
	list := []string{"0000"}
	var next []string
	for len(list) != 0 {
		for _, s := range list {
			for i := 0; i < 4; i++ {
				j := '0' + (s[i]-'0'+1)%10
				this := fmt.Sprintf("%s%c%s", s[:i], j, s[i+1:])
				if !table[this] && !visited[this] {
					visited[this] = true
					next = append(next, this)
					if this == target {
						return step + 1
					}
				}
				j = '0' + (s[i]-'0'+9)%10
				this = fmt.Sprintf("%s%c%s", s[:i], j, s[i+1:])
				if !table[this] && !visited[this] {
					visited[this] = true
					next = append(next, this)
					if this == target {
						return step + 1
					}
				}
			}

		}
		step++
		list = next
		next = []string{}
	}
	return -1
}
