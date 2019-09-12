package problem0006

func convert(s string, numRows int) string {
	direct := 1
	lineIndex := 0
	outStr := make([]string, len(s))
	for i := 0; i < len(s); i++ {
		outStr[lineIndex] = outStr[lineIndex] + string(s[i])
		if numRows > 1 {
			lineIndex = lineIndex + direct
			if lineIndex == 0 || lineIndex == numRows-1 {
				direct *= -1
			}
		}
	}
	outputStr := ""
	for j := 0; j < numRows; j++ {
		outputStr = outputStr + outStr[j]
	}
	return outputStr
}
