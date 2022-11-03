package util

import "strings"

// strings.Cut, but starting from last character, found is either empty or seperator
func RevCut(s, sep string) (leftOf, rightOf string, found bool) {
	if i := strings.LastIndex(s, sep); i >= 0 {
		return s[:i], s[i+len(sep):], true
	}
	return s, "", false
}
