package main

/**

 * @Author: AloneAtWar

 * @Date:   2022/8/18 0:18

 * @Note:	https://leetcode.cn/problems/maximum-equal-frequency/
			1224. Maximum Equal Frequency
			1224. 最大相等频率
 **/

// 关键词是前缀 这个在题目信息上没有提及
// 记录 每个数字的出现次数 与 记录最大出现的频次
// 皆有删除元素的可能性 分类讨论 借由数学关系可以得出
func maxEqualFreq(nums []int) (result int) {
	// 记录每个数字出现次数的 map
	countTable := map[int]int{}
	// 记录最大出现频次
	maxFre := 0
	// 记录频次所对应的数字个数的
	freTable := map[int]int{}
	for i, num := range nums {
		if countTable[num] != 0 {
			freTable[countTable[num]]--
		}
		countTable[num]++
		freTable[countTable[num]]++
		if countTable[num] > maxFre {
			maxFre = countTable[num]
		}
		if maxFre == 1 || (freTable[maxFre] == 1 && maxFre+freTable[maxFre-1]*(maxFre-1) == (i+1)) ||
			(freTable[1] == 1 && 1+freTable[maxFre]*maxFre == (i+1)) {
			result = i + 1
		}
	}
	return
}
