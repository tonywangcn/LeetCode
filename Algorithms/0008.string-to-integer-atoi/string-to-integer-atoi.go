package problem0008

import (
	"math"
)

func myAtoi(s string) int {
	characters := []byte(s)
	sign := 1
	num := 0
	inConversion := false
	for i := 0; i < len(characters); i++ {
		c := characters[i]
		if c == ' ' && inConversion == false {
			continue
		}

		if inConversion == false {
			if c == '+' {
				inConversion = true
			} else if c == '-' {
				inConversion = true
				sign = -1
			} else if v, ok := toDigit(c); ok {
				num = num*10 + v
				inConversion = true
			} else {
				return 0
			}
		} else {
			if v, ok := toDigit(c); ok {
				num = num*10 + v
				numWithSign := num * sign
				if numWithSign > math.MaxInt32 {
					return math.MaxInt32
				} else if numWithSign < math.MinInt32 {
					return math.MinInt32
				}
			} else {
				break
			}
		}
	}
	return num * sign
}

func toDigit(c byte) (int, bool) {
	if '0' <= c && c <= '9' {
		return int(c - '0'), true
	}
	return 0, false
}
