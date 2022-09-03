package main

/**

 * @Author: AloneAtWar

 * @Date:   2022/9/4 3:42

 * @Note:	https://leetcode.cn/problems/special-positions-in-a-binary-matrix/
			1582. Special Positions in a Binary Matrix
			1582. 二进制矩阵中的特殊位置
 **/

func numSpecial(mat [][]int) (result int) {
	m, n := len(mat), len(mat[0])
	row := make([]int, m)
	column := make([]int, n)
	for i, ints := range mat {
		for j, v := range ints {
			row[i] += v
			column[j] += v
		}
	}
	for i, ints := range mat {
		for j, v := range ints {
			if row[i] == 1 && column[j] == 1 && v == 1 {
				result++
			}
		}
	}
	return result
}
