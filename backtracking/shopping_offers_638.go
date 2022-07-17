package backtracking

/*
638. Shopping Offers
https://leetcode.com/problems/shopping-offers/
*/

import (
	"fmt"
)

func shoppingOffersDPMemo(price []int, special [][]int, needs []int) int {

	createKey := func(s []int) string {
		return fmt.Sprintf("%q", s)
	}

	dot := func(a []int, b []int) int {
		if len(a) != len(b) {
			panic("len(a) != len(b)")
		}
		res := 0
		for i := 0; i < len(a); i++ {
			res += a[i] * b[i]
		}
		return res
	}

	subtract := func(a []int, b []int) []int {
		c := make([]int, len(a))
		copy(c, a)
		for i := 0; i < len(a); i++ {
			c[i] -= b[i]
		}
		return c
	}

	canUseOffer := func(a []int, b []int) bool {
		for i := 0; i < len(a); i++ {
			if a[i] < b[i] {
				return false
			}
		}
		return true
	}

	min := func(x int, y int) int {
		if x > y {
			return y
		} else {
			return x
		}
	}

	memo := make(map[string]int)

	var shopping func(price []int, special [][]int, needs []int) int

	shopping = func(price []int, special [][]int, needs []int) int {

		key := createKey(needs)

		if val, ok := memo[key]; ok {
			return val
		}

		res := dot(price, needs)

		for i := 0; i < len(special); i++ {
			specialOffer := special[i]
			if canUseOffer(needs, specialOffer) {
				newNeeds := subtract(needs, specialOffer)
				priceWithSpecialOffer := specialOffer[len(specialOffer)-1] + shopping(price, special, newNeeds)
				res = min(res, priceWithSpecialOffer)
			}
		}
		memo[key] = res
		return res
	}
	return shopping(price, special, needs)
}

func shoppingOffersDP1(price []int, special [][]int, needs []int) int {

	dot := func(a []int, b []int) int {
		if len(a) != len(b) {
			panic("len(a) != len(b)")
		}
		res := 0
		for i := 0; i < len(a); i++ {
			res += a[i] * b[i]
		}
		return res
	}

	subtract := func(a []int, b []int) []int {
		c := make([]int, len(a))
		copy(c, a)
		for i := 0; i < len(a); i++ {
			c[i] -= b[i]
		}
		return c
	}

	canUseOffer := func(a []int, b []int) bool {
		for i := 0; i < len(a); i++ {
			if a[i] < b[i] {
				return false
			}
		}
		return true
	}

	min := func(x int, y int) int {
		if x > y {
			return y
		} else {
			return x
		}
	}

	res := dot(price, needs)
	for i := 0; i < len(special); i++ {
		specialOffer := special[i]
		if canUseOffer(needs, specialOffer) {
			newNeeds := subtract(needs, specialOffer)
			priceWithSpecialOffer := specialOffer[len(specialOffer)-1] + shoppingOffersDP1(price, special, newNeeds)
			res = min(res, priceWithSpecialOffer)
		}
	}
	return res
}

func ShoppingOffersDP1(price []int, special [][]int, needs []int) int {
	return shoppingOffersDP1(price, special, needs)
}

func ShoppingOffersDPMemo(price []int, special [][]int, needs []int) int {
	return shoppingOffersDPMemo(price, special, needs)
}
