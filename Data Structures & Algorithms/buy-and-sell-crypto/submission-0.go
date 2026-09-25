func maxProfit(prices []int) int {
	maxProfit := 0
	minPrice := prices[0]

	for _, price := range prices {
		profit := price - minPrice
		if profit > maxProfit {
			maxProfit = profit
		}
		if price < minPrice {
			minPrice = price
		}
	} 
	return maxProfit
}
