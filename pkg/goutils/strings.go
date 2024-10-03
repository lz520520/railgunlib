package goutils

import "strings"

func CompareIgnoreCase(s1, s2 string) bool {
	return strings.ToLower(s1) == strings.ToLower(s2)
}

func ContainIgnoreCase(s, substr string) bool {
	return strings.Contains(strings.ToLower(s), strings.ToLower(substr))
}
