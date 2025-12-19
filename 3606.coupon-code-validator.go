package leetcode

import (
	"sort"
	"unicode"
)

// @leet start
func isValidCode(s string) bool {
	if len(s) == 0 {
		return false
	}

	for _, r := range s {
		if !(unicode.IsLetter(r) || unicode.IsDigit(r) || r == '_') {
			return false
		}
	}
	return true
}
func isCouponValid(code string, businessLine string, isActive bool) string {
	if !isActive {
		return "invalid"
	}

	if !(businessLine == "electronics" || businessLine == "grocery" || businessLine == "pharmacy" || businessLine == "restaurant") {
		return "invalid"
	}

	if isValidCode(code) == false {
		return "invalid"
	}
	return "valid"
}

func validateCoupons(code []string, businessLine []string, isActive []bool) []string {
	n := len(code)
	electronicsResult := []string{}
	groceryResult := []string{}
	pharmacyResult := []string{}
	restaurantResult := []string{}

	for i := 0; i < n; i++ {
		if isCouponValid(code[i], businessLine[i], isActive[i]) == "valid" {
			switch businessLine[i] {
			case "electronics":
				electronicsResult = append(electronicsResult, code[i])
			case "grocery":
				groceryResult = append(groceryResult, code[i])
			case "pharmacy":
				pharmacyResult = append(pharmacyResult, code[i])
			case "restaurant":
				restaurantResult = append(restaurantResult, code[i])
			}
		}
	}

	sort.Strings(electronicsResult)
	sort.Strings(groceryResult)
	sort.Strings(pharmacyResult)
	sort.Strings(restaurantResult)

	return append(append(append(electronicsResult, groceryResult...), pharmacyResult...), restaurantResult...)
}
// @leet end

// Keynold
