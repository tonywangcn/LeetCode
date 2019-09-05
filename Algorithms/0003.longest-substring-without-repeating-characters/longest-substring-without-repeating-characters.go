package problem0003

func lengthOfLongestSubstring_BruteForce(s string) int {
	// 设置两个指针，start,final 从起始点开始，判断 slice[start: k+1] 是否为 unique，若否则 start ++
	var start int
	var final int
	MaxCount := 1
	CurrentCount := 1
	slice := []rune(s)
	// fmt.Println(slice)
	if len(slice) == 0 {
		return 0
	}
	for k := 1; k < len(slice); k++ {
		if isUnique(slice[start : k+1]) {
			final = k
		} else {
			start++
		}
		CurrentCount = final - start + 1
		if MaxCount < CurrentCount {
			MaxCount = CurrentCount
		}
	}
	return MaxCount

}

func isUnique(slice []rune) bool {
	container := make(map[rune]int, len(slice))
	for _, v := range slice {
		if container[v] == 0 {
			container[v] = 1
		} else {
			return false
		}
	}
	return true

}

func lengthOfLongestSubstring_HashMap(s string) int {
	// 创建一个 hashmap，记录最大值，当前值和起始位，遍历字符串
	dict := make(map[rune]int, len(s))
	max := 0
	curr := 0
	start := 0
	for k, v := range []rune(s) {
		// fmt.Println(k, v)
		if dict[v] != 0 && dict[v] >= start {
			if max < curr {
				max = curr
			}
			curr = k - dict[v] + 1
			start = dict[v]
		} else {
			curr++
		}
		dict[v] = k + 1
	}
	if max < curr {
		max = curr
	}
	return max
}
