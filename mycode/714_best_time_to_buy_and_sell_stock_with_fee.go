package mycode

// https://leetcode.com/problems/best-time-to-buy-and-sell-stock-with-transaction-fee/solution/
// Elegant DP solution with O(N) time and O(1) Space, LEARN FROM IT!
func maxProfit(prices []int, fee int) int {
	// cash[i]: maximum money in hand at day i if you don't own a share
	// hold[i]: maximum money in hand at day i if you own a share
	// both cash[i] and hold[i] only depends on cash[i-1]/hold[i-1]
	// so space optimization is applied
	cash, hold := 0, -prices[0]

	for i := 1; i < len(prices); i++ {
		cash = maxInt(cash, hold+prices[i]-fee) // do nothing / sell the share to become flat
		hold = maxInt(hold, cash-prices[i])     // do nothing / buy one share to become long
	}

	return cash // at the end we don't want to be long (holding a share)
}
