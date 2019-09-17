package problem0011

func maxArea(height []int) int {
	maxA := 0
	for k, v1 := range height {
		for j, v2 := range height[k:] {
			area := j * min(v1, v2)
			maxA = max(maxA, area)
		}
	}
	return maxA
}

func maxAreaTwoPointer(height []int) int {
	maxA := 0
	left, right := 0, len(height)-1
	for left < right {
		area := (right - left) * min(height[left], height[right])
		maxA = max(area, maxA)
		if height[left] > height[right] {
			right--
		} else {
			left++
		}
	}
	return maxA
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
