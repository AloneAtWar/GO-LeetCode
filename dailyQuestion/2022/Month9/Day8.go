package main

import "fmt"

/**

 * @Author: AloneAtWar

 * @Date:   2022/9/11 10:59

 * @Note:	https://leetcode.cn/problems/beautiful-arrangement-ii/
			667. Beautiful Arrangement II
			667. 优美的排列 II
 **/

// 回溯
//func constructArray(n int, k int) ([]int) {
//	visited1:=make(map[int]int)
//	visited2:=make(map[int]bool)
//	curr:=[]int{1}
//	visited2[1]=true
//	var dfs func(i int)bool
//	dfs= func(i int) bool {
//		if i>=n && len(visited1)==k{
//			return true
//		}
//		if len(visited1)>k{
//			return false
//		}
//		for j := 1; j <= n; j++ {
//			if !visited2[j] {
//				length:=abs(j-curr[i-1])
//				visited1[length]++
//				curr=append(curr,j)
//				visited2[j]=true
//				if dfs(i+1){
//					return true
//				}
//				curr=curr[:i]
//				visited2[j]=false
//				visited1[length]--
//				if visited1[length]==0{
//					delete(visited1,length)
//				}
//			}
//		}
//		return false
//	}
//	dfs(1)
//	return curr
//}
//
//func abs(a int)int{
//	if a>0{
//		return a
//	}
//	return -a
//}
func constructArray(n, k int) []int {
	ans := make([]int, 0, n)
	for i := 1; i < n-k; i++ {
		ans = append(ans, i)
	}
	for i, j := n-k, n; i <= j; i++ {
		ans = append(ans, i)
		if i != j {
			ans = append(ans, j)
		}
		j--
	}
	return ans
}

func main() {
	fmt.Println(constructArray(3, 1))
}

//  先导题
// https://leetcode.cn/problems/beautiful-arrangement/submissions/
//  526. Beautiful Arrangement
//func countArrangement(n int) (ans int) {
//	vis := make([]bool, n+1)
//	match := make([][]int, n+1)
//	for i := 1; i <= n; i++ {
//		for j := 1; j <= n; j++ {
//			if i%j == 0 || j%i == 0 {
//				match[i] = append(match[i], j)
//			}
//		}
//	}
//
//	var backtrack func(int)
//	backtrack = func(index int) {
//		if index > n {
//			ans++
//			return
//		}
//		for _, x := range match[index] {
//			if !vis[x] {
//				vis[x] = true
//				backtrack(index + 1)
//				vis[x] = false
//			}
//		}
//	}
//	backtrack(1)
//	return
//}
