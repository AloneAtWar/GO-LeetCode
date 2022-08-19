package main

/**

 * @Author: AloneAtWar

 * @Date:   2022/8/19 23:34

 * @Note:

 **/

type SparseVector struct {
	table map[int]int
}

func Constructor(nums []int) SparseVector {
	table := make(map[int]int)
	for i, num := range nums {
		if num != 0 {
			table[i] = num
		}
	}
	return SparseVector{table}
}

// Return the dotProduct of two sparse vectors
func (this *SparseVector) dotProduct(vec SparseVector) int {
	result := 0
	for i, v := range vec.table {
		result += this.table[i] * v
	}
	return result
}
