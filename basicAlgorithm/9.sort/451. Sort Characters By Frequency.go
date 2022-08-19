package main

import "sort"

/**

 * @Author: AloneAtWar

 * @Date:   2022/8/19 19:40

 * @Note:	https://leetcode.cn/problems/sort-characters-by-frequency/
			451. Sort Characters By Frequency
			451. 根据字符出现频率排序
 **/

// 计数排序??
func frequencySort(s string) string {
	var result []byte
	table := make(map[int]int)
	for _, v := range s {
		table[int(v)]++
	}
	var ints []int
	for i, _ := range table {
		ints = append(ints, i)
	}
	sort.Slice(ints, func(i, j int) bool {
		return table[ints[i]] > table[ints[j]]
	})
	for _, v := range ints {
		for i := 0; i < table[v]; i++ {
			result = append(result, byte(v))
		}
	}
	return string(result)
}
