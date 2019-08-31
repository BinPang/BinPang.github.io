package main

import "fmt"

func main() {
	println(maxProfit(2, []int{1, 2, 4, 2, 5, 7, 2, 4, 9, 0, 9}))
	//println(maxProfit(3, []int{1, 2, 4, 2, 5, 7, 2, 4, 9, 0, 9}))
}

//func maxProfit(k int, prices []int) int {
//	l := len(prices)
//	if l == 0 || k <= 0 {
//		return 0 //如果不合法，直接返回0
//	}
//	dp := make([][]int, k+1)
//	for i := 0; i <= k; i++ {
//		dp[i] = make([]int, l) //分配(k+1)*l内存空间
//	}
//	tmp := 0 //这里没啥含义，就是数学式推导出来，只是临时记录使用的
//	for kk := 1; kk <= k; kk++ {
//		tmp = -prices[0]
//		for i := 1; i < l; i++ {
//			tmp = max(tmp, dp[kk-1][i-1]-prices[i-1])
//			dp[kk][i] = max(dp[kk][i-1], prices[i] + tmp)
//		}
//	}
//	fmt.Println(dp)
//	return dp[k][l-1]
//}

//func maxProfit(k int, prices []int) int {
//	l := len(prices)
//	if l == 0 || k <= 0 {
//		return 0 //如果不合法，直接返回0
//	}
//	dp := make([][]int, k+1)
//	for i := 0; i <= k; i++ {
//		dp[i] = make([]int, l) //分配(k+1)*l内存空间
//	}
//	for kk := 1; kk <= k; kk++ {
//		for i := 1; i < l; i++ {
//			dp[kk][i] = dp[kk][i-1]
//			for j := 0; j < i; j++ {
//				if j == 0 {
//					dp[kk][i] = max(dp[kk][i], prices[i]-prices[j]) //如果第0天有买入，直接用prices[i]-prices[j]
//				} else {
//					//prices[i]-prices[j]:第j天买入并且在第i天卖出的收益，再加上第j-1天且kk-1次交易的收益
//					dp[kk][i] = max(dp[kk][i], prices[i]-prices[j]+dp[kk-1][j-1])
//				}
//			}
//		}
//	}
//	fmt.Println(dp)
//	return dp[k][l-1]
//}

func maxProfit(k int, prices []int) int {
	l := len(prices)
	if l == 0 || k <= 0 {
		return 0
	}
	dp := make([][]int, 2)
	for i := 0; i < 2; i++ {
		dp[i] = make([]int, l+1)
	}
	tmp := 0
	for kk := 1; kk <= k; kk++ {
		tmp = dp[kk&1^1][0] - prices[0]
		for ll := 1; ll <= l; ll++ {
			tmp = max(tmp, dp[kk&1^1][ll-1]-prices[ll-1])
			dp[kk%2][ll] = max(dp[kk%2][ll-1], tmp+prices[ll-1])
		}
		if dp[0][l] == dp[1][l] {
			return dp[0][l]
		}
	}
	println(fmt.Sprintf("%+v", dp))
	return dp[k%2][l]
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
