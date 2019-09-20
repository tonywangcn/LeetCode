package problem0015

import (
	"sort"
)

// 首先将 nums 升序排序，然后设定指针 i，l，r，其中
// l, r = i + 1, length -1，i 为最左的指针，及元组中最小的，需保证其永远小于或等于 0.
// 当 l < r 时，计算 total 的数值，当其小于 0 时，则需 l 向右移，当大于 0 时，则需 r 向左移。
// 完成后分别对 l ++, r--，计算下一位。

func threeSum(nums []int) [][]int {
	out := make([][]int, 0)
	sort.Ints(nums)
	length := len(nums)
	for i := 0; i < length-1; i++ {
		if nums[i] > 0 {
			break
		}
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		l, r := i+1, length-1
		for l < r {
			total := nums[i] + nums[l] + nums[r]
			if total < 0 {
				l++
			} else if total > 0 {
				r--
			} else {
				out = append(out, []int{nums[i], nums[l], nums[r]})
				for l < r && nums[l] == nums[l+1] {
					l++
				}
				for l < r && nums[r] == nums[r-1] {
					r--
				}
				l++
				r--
			}
		}
	}
	return out
}
