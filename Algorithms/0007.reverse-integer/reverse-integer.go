package problem0007

import (
	"math"
)

func reverse(x int) int {
	var tar int
	var pop int
	var rev bool
	if x < 0 {
		rev = false
	} else {
		rev = true
	}
	x = Abs(x)
	for x > 0 {
		pop = x % 10
		if tar > math.MaxInt32/10 || (tar == math.MaxInt32/10 && pop > 7) {
			return 0
		}
		if tar < math.MinInt32/10 || (tar == math.MinInt32/10 && pop < -8) {
			return 0
		}
		tar = tar*10 + pop
		x = x / 10
	}
	if !rev {
		tar = -tar
	}
	return tar
}

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
