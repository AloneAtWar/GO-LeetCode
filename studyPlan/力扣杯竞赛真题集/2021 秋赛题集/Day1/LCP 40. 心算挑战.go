package main

import "sort"

/**

 * @Author: AloneAtWar

 * @Date:   2022/9/4 5:00

 * @Note:

 **/

////10 9 5 2 1
////func maxmiumScore(cards []int, cnt int) int {
////	sort.Slice(cards, func(i, j int) bool {
////		return cards[i]>cards[j]
////	})
////	n:=len(cards)
////	dp := make([][][2]int, n+1)
////	for i := 0; i < n+1; i++ {
////		dp[i]=make([][2]int,cnt+1)
////	}
////
////	for i := 1; i < n+1; i++ {
////		for j := 1; j <= cnt && j<=i; j++ {
////			if cards[i-1]%2==0{
////				if j==1 || dp[i-1][j-1][0]!=0{
////					dp[i][j][0] = max(dp[i-1][j][0], dp[i-1][j-1][0]+cards[i-1])
////				}else{
////					dp[i][j][0]=dp[i-1][j][0]
////				}
////				if j!=1 && dp[i-1][j-1][1]!=0{
////					dp[i][j][1] = max(dp[i-1][j][1], dp[i-1][j-1][1]+cards[i-1])
////				}else{
////					dp[i][j][1]=dp[i-1][j][1]
////				}
////
////			}else{
////				if j==1 || dp[i-1][j-1][0]!=0{
////					dp[i][j][1] = max(dp[i-1][j][1], dp[i-1][j-1][0]+cards[i-1])
////				}else{
////					dp[i][j][1]=dp[i-1][j][1]
////				}
////				if j!=1 && dp[i-1][j-1][1]!=0{
////					dp[i][j][0] = max(dp[i-1][j][0], dp[i-1][j-1][1]+cards[i-1])
////				}else{
////					dp[i][j][0]=dp[i-1][j][0]
////				}
////			}
////			//fmt.Printf("前%d个元素 选%d 最大值为 %d %d\n",i,j,dp[i][j][0],dp[i][j][1])
////		}
////	}
////	return dp[n][cnt][0]
////}
//
//func maxmiumScore(cards []int, cnt int) int {
//	n:=len(cards)
//	if cnt>n{
//		return 0
//	}
//	sort.Slice(cards, func(i, j int) bool {
//		return cards[i]>cards[j]
//	})
//	dp:=make([][2]int,cnt+1)
//
//	for i := 1; i < n+1; i++ {
//		var j int =cnt
//		if j>i{
//			j=i
//		}
//		for ; j>=1; j-- {
//			if cards[i-1]%2==0{
//				if j==1 || dp[j-1][0]!=0{
//					dp[j][0] = max(dp[j][0], dp[j-1][0]+cards[i-1])
//				}else{
//					dp[j][0]=dp[j][0]
//				}
//				if j!=1 && dp[j-1][1]!=0{
//					dp[j][1] = max(dp[j][1], dp[j-1][1]+cards[i-1])
//				}else{
//					dp[j][1]=dp[j][1]
//				}
//
//			}else{
//				if j==1 || dp[j-1][0]!=0{
//					dp[j][1] = max(dp[j][1], dp[j-1][0]+cards[i-1])
//				}else{
//					dp[j][1]=dp[j][1]
//				}
//				if j!=1 && dp[j-1][1]!=0{
//					dp[j][0] = max(dp[j][0], dp[j-1][1]+cards[i-1])
//				}else{
//					dp[j][0]=dp[j][0]
//				}
//			}
//			//fmt.Printf("前%d个元素 选%d 最大值为 %d %d\n",i,j,dp[j][0],dp[j][1])
//		}
//	}
//	return dp[cnt][0]
//}
//func main() {
//	maxmiumScore([]int{1,10,5,2,9},4)
//}
//func max(ints ...int)int{
//	result:=ints[0]
//	for _, v := range ints {
//		if v>result{
//			result=v
//		}
//	}
//	return result
//}

func maxmiumScore(cards []int, cnt int) (ans int) {
	sort.Sort(sort.Reverse(sort.IntSlice(cards)))
	sum := 0
	for _, v := range cards[:cnt] {
		sum += v
	}
	if sum&1 == 0 {
		return sum
	}
	// 在 cards[cnt:] 中找一个最大的且奇偶性和 x 不同的元素，替换 x
	replace := func(x int) {
		for _, v := range cards[cnt:] {
			if v&1 != x&1 {
				ans = max(ans, sum-x+v)
				break
			}
		}
	}
	replace(cards[cnt-1]) // 替换 cards[cnt-1]
	for i := cnt - 2; i >= 0; i-- {
		if cards[i]&1 != cards[cnt-1]&1 { // 找一个最小的且奇偶性不同于 cards[cnt-1] 的元素，将其替换掉
			replace(cards[i])
			break
		}
	}
	return
}

func max(a, b int) int {
	if b > a {
		return b
	}
	return a
}
