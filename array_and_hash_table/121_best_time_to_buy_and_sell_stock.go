package arrays_and_hashing

// link to the task on leetcode
// https://leetcode.com/problems/best-time-to-buy-and-sell-stock/description/

func maxProfit(prices []int) int {
	minPrice := prices[0]
	maxProfit := 0
	curPrice := 0

	for i := range prices {
		curPrice = prices[i] - minPrice
		if curPrice > maxProfit {
			maxProfit = curPrice
		}

		if prices[i] < minPrice {
			minPrice = prices[i]
		}
	}

	return maxProfit
}
