package problem0005

func longestPalindrome(s string) string {
	longestDict := make(map[int]string, len(s))
	longestNum := 0
	if len(s) == 1 {
		return s
	}
	for i, _ := range s {
		for j, _ := range s[i:] {
			target := s[i:][:j+1]
			ok := isPalindrome(target)
			if ok {
				if longestNum < len(target) {
					longestDict[len(target)] = target
					longestNum = len(target)
				}
			}
		}
	}
	return longestDict[longestNum]
}

func isPalindrome(s string) bool {
	length := len(s)
	if length == 0 {
		return false
	}
	median := length / 2
	for k, v := range s[:median] {
		if byte(v) != s[length-k-1] {
			return false
		}
	}
	return true
}

func longestPalindromeDP(s string) string {
	n := len(s)
	subTable := make([][]bool, n)
	maxLength := 1
	i := 0
	for i < n {
		subTable[i] = make([]bool, n)
		subTable[i][i] = true
		i++
	}
	var start int
	i = 0
	for i < n-1 {
		if s[i] == s[i+1] {
			subTable[i][i+1] = true
			start = i
			maxLength = 2
		}
		i++
	}
	k := 3
	for k <= n {
		i = 0
		for i < (n - k + 1) {
			j := i + k - 1
			if subTable[i+1][j-1] && s[i] == s[j] {
				subTable[i][j] = true
				if k > maxLength {
					start = i
					maxLength = k
				}
			}
			i++
		}
		k++
	}
	return s[start : start+maxLength]

}
