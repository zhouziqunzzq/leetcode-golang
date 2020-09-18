package mycode

func maxProfit2(prices []int) int {
	n := len(prices)
	if n < 2 {
		return 0
	}

	ans := 0
	buyPrice := 0

	const LONG int8 = 1
	const FLAT int8 = 0
	var status int8 = 0

	for i, _ := range prices {
		if i == 0 {
			if prices[0] < prices[1] {
				buyPrice = prices[0]
				status = LONG
			}
		} else if i == n - 1 {
			if status == LONG {
				ans += prices[n - 1] - buyPrice
				status = FLAT
			}
		} else {
			if status == FLAT && prices[i] <= prices[i - 1] && prices[i] < prices[i + 1] {
				buyPrice = prices[i]
				status = LONG
			} else if status == LONG && prices[i] >= prices[i - 1] && prices[i] > prices[i + 1] {
				ans += prices[i] - buyPrice
				status = FLAT
			}
		}
	}

	return ans
}
