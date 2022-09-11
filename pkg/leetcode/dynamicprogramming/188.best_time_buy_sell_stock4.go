package dynamicprogramming

/*
188. Best Time to Buy and Sell Stock IV
https://leetcode.com/problems/best-time-to-buy-and-sell-stock-iv/
*/

const MinInt = -int(^uint(0)>>1) - 1

func maxProfitDPSlow(k int, prices []int) int {

	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	numDays := len(prices)
	if k == 0 || numDays == 0 {
		return 0
	}
	numTransactions := k + 1
	t := make([][]int, numTransactions)
	for i := 0; i < numTransactions; i++ {
		t[i] = make([]int, numDays)
	}
	for transaction := 1; transaction < numTransactions; transaction++ {
		for day := 1; day < numDays; day++ {
			maxVal := t[transaction][day-1]
			for m:=0; m<day; m++ {
				maxVal = max(maxVal, prices[day] - prices[m] + t[transaction-1][m])	
			}
			t[transaction][day] = maxVal
		}
	}
	return t[numTransactions-1][numDays-1]
}

func maxProfitDP(k int, prices []int) int {

	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	numDays := len(prices)
	if k == 0 || numDays == 0 {
		return 0
	}
	numTransactions := k + 1
	t := make([][]int, numTransactions)
	for i := 0; i < numTransactions; i++ {
		t[i] = make([]int, numDays)
	}
	for transaction := 1; transaction < numTransactions; transaction++ {
		maxDiff := -prices[0]
		for day := 1; day < numDays; day++ {
			t[transaction][day] = max(t[transaction][day-1], prices[day]+maxDiff)
			maxDiff = max(maxDiff, t[transaction-1][day]-prices[day])
		}
	}
	return t[numTransactions-1][numDays-1]
}

func maxProfitRecursion(k int, prices []int) int {

	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	n := len(prices)
	memo := make(map[bool][][]int)
	memo[true] = make([][]int, k+1)
	memo[false] = make([][]int, k+1)
	for i := 0; i < k+1; i++ {
		memo[true][i] = make([]int, n)
		memo[false][i] = make([]int, n)
		for j := 0; j < n; j++ {
			memo[false][i][j] = -1
			memo[true][i][j] = -1
		}
	}

	var recursion func(pos int, bought bool, transCount int) int

	recursion = func(pos int, bought bool, transCount int) int {
		if transCount == 0 || pos >= n {
			return 0
		}
		if memo[bought][transCount][pos] != -1 {
			return memo[bought][transCount][pos]
		}
		sum := recursion(pos+1, bought, transCount)
		if bought {
			sum = max(sum, recursion(pos+1, false, transCount-1)+prices[pos])
		} else {
			sum = max(sum, recursion(pos+1, true, transCount)-prices[pos])
		}
		memo[bought][transCount][pos] = sum
		return sum
	}

	return recursion(0, false, k)
}

func maxProfitBacktracking(k int, prices []int) int {
	l := len(prices)

	type State struct {
		idx               int
		isBought          bool
		profit            int
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
				if s.profit < 0 {
					return
				}
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

func MaxProfitDPSlow4(k int, prices []int) int {
	return maxProfitDPSlow(k, prices)
}

func MaxProfitDP4(k int, prices []int) int {
	return maxProfitDP(k, prices)
}

func MaxProfitRecursion4(k int, prices []int) int {
	return maxProfitRecursion(k, prices)
}

func MaxProfit4Backtracking(k int, prices []int) int {
	return maxProfitBacktracking(k, prices)
}
