package problem0013

func romanToInt(s string) int {
	//
	out := 0
	strs := map[string]int{"M": 1000, "CM": 900, "D": 500, "CD": 400, "C": 100, "XC": 90, "L": 50, "XL": 40, "X": 10, "IX": 9, "V": 5, "IV": 4, "I": 1}
	out += strs[string(s[len(s)-1])]
	for i := 0; i < len(s)-1; i++ {
		if strs[string(s[i])] >= strs[string(s[i+1])] {
			out += strs[string(s[i])]
		} else {
			out -= strs[string(s[i])]
		}
	}
	return out
}
