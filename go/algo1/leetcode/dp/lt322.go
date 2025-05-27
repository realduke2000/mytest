package dp

import (
	"math"
	"sort"
)

func coinChange2(coins []int, amount int, rem map[int]int) int {
	if amount == 0 {
		return 0
	}
	if coins[0] > amount {
		return -1
	}
	if v, ok := rem[amount]; ok {
		return v
	}

	if amount > coins[0] && amount < coins[len(coins)-1] {
		for i := 0; i < len(coins); i++ {
			if coins[i] == amount {
				rem[amount] = 1
				return 1
			}
		}
	}

	_min := math.MaxInt
	for i := len(coins) - 1; i >= 0; i-- {
		if tmp := coinChange2(coins, amount-coins[i], rem); tmp != -1 {
			_min = min(_min, tmp+1)
		}
	}
	if _min == math.MaxInt {
		rem[amount] = -1
		return -1
	} else {
		rem[amount] = _min
		return _min
	}
}

func coinDp(coins []int, amount int) int {
	dp := make([]int, amount+1)
	dp[0] = 0
	for i := 1; i < amount; i++ {

	}
	return dp[amount]
}

func coinChange(coins []int, amount int) int {
	sort.Ints(coins)

	if amount == 0 {
		return 0
	}
	if amount < coins[0] {
		return -1
	}
	i := 0
	for ; i < len(coins); i++ {
		if coins[i] > amount {
			break
		}
	}
	rem := make(map[int]int)
	return coinChange2(coins[:i], amount, rem)
}
