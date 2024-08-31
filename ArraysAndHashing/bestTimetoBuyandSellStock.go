package arrays_and_hashing

// link to the task on leetcode
// https://leetcode.com/problems/best-time-to-buy-and-sell-stock/description/

func MaxProfit(prices []int) int {
	minPrice := prices[0]
	maxProfit := 0

	for i := range prices {
		if prices[i]-minPrice > maxProfit {
			maxProfit = prices[i] - minPrice
		}

		if prices[i] < minPrice {
			minPrice = prices[i]
		}
	}

	return maxProfit
}
