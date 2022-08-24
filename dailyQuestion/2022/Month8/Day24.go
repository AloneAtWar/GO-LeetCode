package main

/**

 * @Author: AloneAtWar

 * @Date:   2022/8/24 23:17

 * @Note:

 **/

func canBeEqual(target []int, arr []int) bool {
	if len(target) != len(arr) {
		return false
	}
	table := make(map[int]int)
	for i := 0; i < len(target); i++ {
		table[target[i]]++
		if table[target[i]] == 0 {
			delete(table, target[i])
		}
		table[arr[i]]--
		if table[arr[i]] == 0 {
			delete(table, arr[i])
		}
	}
	return len(table) == 0
}
