package dynamicprogramming

/*
188. Best Time to Buy and Sell Stock IV
https://leetcode.com/problems/best-time-to-buy-and-sell-stock-iv/
*/

func maxProfit(k int, prices []int) int {
	l := len(prices)

	type State struct {
		idx int
		isBought bool
		profit int
		transactionsCount int
	}

	result := 0

	var search func(state State) 

	search = func(s State) {
		if s.idx == l {
			if !s.isBought && s.profit > result {
				result = s.profit
			}
			return
		}
		price := prices[s.idx]
		if s.isBought {
			s.idx += 1
			search(s)
			if s.transactionsCount < 2*k {
				s.isBought = false
				s.profit += price
				s.transactionsCount += 1
				search(s)	
			}		
		} else {
			s.idx += 1
			search(s)
			if s.transactionsCount < 2*k {
				s.isBought = true
				s.profit -= price
				s.transactionsCount += 1
				search(s)	
			}
		}
	}

	search(State{0, false, 0, 0})

	return result	
}

func MaxProfit4(k int, prices []int) int {
	return maxProfit(k, prices)
}