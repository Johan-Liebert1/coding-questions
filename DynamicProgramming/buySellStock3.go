package dynamicprogramming

func profit3(prices []int, index, lastBought, remainingTransactions int, cache [][][]int) int {
	if index >= len(prices) {
		return 0
	}

	if cache[lastBought][remainingTransactions][index] != -1 {
		return cache[lastBought][remainingTransactions][index]
	}

	nothing := profit3(prices, index+1, lastBought, remainingTransactions, cache)

	buy := 0
	sell := 0

	if lastBought == 0 && remainingTransactions > 0 {
		buy = -prices[index] + profit3(prices, index+1, 1, remainingTransactions-1, cache)
	}

	if lastBought != 0 {
		sell = prices[index] + profit3(prices, index+1, 0, remainingTransactions, cache)
	}

	profit := max(buy, sell, nothing)

	cache[lastBought][remainingTransactions][index] = profit

	return profit
}

func buySellStock3(prices []int) int {
	cache := make([][][]int, 2) // for lastBought (0 or 1)

	for i := range 2 {
		cache[i] = make([][]int, 3) // for remaining transactions 0, 1, 2

		for j := range cache[i] {
            cache[i][j] = make([]int, len(prices))

            for k := range cache[i][j] {
                cache[i][j][k] = -1
            }
		}
	}

	return profit3(prices, 0, 0, 2, cache)
}
