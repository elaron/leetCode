package dp

import (
	"fmt"
)

func main_coinChange() {
	//coins := []int{3,7,405,436}
	//amount := 8839
	//coins := []int{288,160,10,249,40,77,314,429}
	//amount := 9208
	//coins := []int{470,18,66,301,403,112,360}
	//amount := 8235
	//coins := []int{1,2147483647}
	//amount := 2
	coins := []int{456, 117, 5, 145}
	amount := 1459
	//coins := []int{1}
	//amount := 0
	ret := coinChange(coins, amount)
	fmt.Println(ret)
}

func coinChange(coins []int, amount int) int {
	if len(coins) == 0 {
		return -1
	}
	dp := make([]int, amount+1)
	availableCoins := make([]int, 0)
	for _, coin := range coins {
		if coin > amount {
			continue
		}
		availableCoins = append(availableCoins, coin)
		dp[coin] = 1
	}
	for i := 1; i < len(dp); i++ {
		if dp[i] != 0 {
			continue
		}
		dp[i] = -1
		for _, coin := range coins {
			lastIdx := i - coin
			if lastIdx <= 0 || dp[lastIdx] == -1 {
				continue
			}
			if dp[i] == -1 || (dp[i] != -1 && dp[lastIdx]+1 < dp[i]) {
				dp[i] = dp[lastIdx] + 1
			}
		}
	}
	return dp[amount]
}
