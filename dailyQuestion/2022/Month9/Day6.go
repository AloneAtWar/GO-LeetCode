package main

/**

 * @Author: AloneAtWar

 * @Date:   2022/9/6 18:23

 * @Note:	https://leetcode.cn/problems/count-unique-characters-of-all-substrings-of-a-given-string/
			828. Count Unique Characters of All Substrings of a Given String
			828. 统计子串中的唯一字符
 **/

// 超时
//func uniqueLetterString(s string) (result int) {
//	n:=len(s)
//	dp := make([]int, n+1)
//	for i, _:= range s {
//		add:=0
//		table:=map[byte]int{}
//		has:=map[byte]bool{}
//		for j := i; j >=0 ; j-- {
//			table[s[j]]++
//			if table[s[j]]>1{
//				delete(has,s[j])
//			}else{
//				has[s[j]]=true
//			}
//			add+=len(has)
//		}
//		dp[i+1]=dp[i]+add
//	}
//	return dp[n]
//}
func uniqueLetterString(s string) (ans int) {
	idx := map[rune][]int{}
	for i, c := range s {
		idx[c] = append(idx[c], i)
	}
	for _, arr := range idx {
		arr = append(append([]int{-1}, arr...), len(s))
		for i := 1; i < len(arr)-1; i++ {
			ans += (arr[i] - arr[i-1]) * (arr[i+1] - arr[i])
		}
	}
	return
}

func main() {
	uniqueLetterString("ABA")
}
