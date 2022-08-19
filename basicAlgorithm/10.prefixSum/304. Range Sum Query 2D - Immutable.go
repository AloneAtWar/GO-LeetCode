package main

// 这题挺无聊的 我就之前CP 我自己以前的代码了

type NumMatrix struct {
	table [][]int
	sum   [][]int
}

func Constructor(matrix [][]int) NumMatrix {
	m := len(matrix)
	n := len(matrix[0])
	table := make([][]int, m+1)
	table[0] = make([]int, n+1)
	for i := 0; i < m; i++ {
		table[i+1] = make([]int, n+1)
		for j := 0; j < n; j++ {
			table[i+1][j+1] = table[i][j+1] + table[i+1][j] - table[i][j] + matrix[i][j]
		}
	}

	return NumMatrix{matrix, table}
}

func (this *NumMatrix) SumRegion(row1 int, col1 int, row2 int, col2 int) int {
	return this.sum[row2+1][col2+1] + this.sum[row1][col1] - this.sum[row1][col2+1] - this.sum[row2+1][col1]
}
