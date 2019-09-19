package problem0014

import (
	"strings"
)

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	comm := strs[0]
	for _, v := range strs {
		for strings.Index(v, comm) != 0 {
			comm = comm[:len(comm)-1]
		}
	}
	return comm
}
