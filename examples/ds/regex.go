package main

import (
	"fmt"
	"math"
)

var setIndex int
var originalLength int
func findMatchVal(text string, pattern string) bool {
	if text == pattern || pattern == "*" || len(pattern) == 0 {
		return true
	}
	if len(text) == 0 {
		return false
	}
	if string(pattern[0]) == "*" {
		return findMatchVal(text, pattern[1:]) || findMatchVal(text[1:], pattern)
	}

	if text[0] == pattern[0] {
		if setIndex > (originalLength - len(text)) {
			setIndex = originalLength - len(text)
		}
		return findMatchVal(text[1:], pattern[1:])
	} else if text[0] != pattern[0] {
		return findMatchVal(text[1:], pattern)
	}
	return false
}

func findMatchIndex(s string, p string) int32 {
	setIndex = math.MaxInt16
	originalLength = len(s)
	ans := findMatchVal(s, p)
	if !ans {
		return -1
	}
	return int32(setIndex)
}

func isSubstring(s string, p string, f bool) bool {
	if s == p || len(p) == 0 {
		return true
	}
	if len(s) <= len(p) {
		return false
	}
	if s[0] == p[0] {
		return isSubstring(s[1:], p[1:], true) || isSubstring(s[1:], p, false)
	}
	if f {
		return false
	}
	return isSubstring(s[1:], p, false)
}

func isSubstringAlternate(s string, p string) bool {
	if len(p) > len(s) {
		return false
	}
	if s[:len(p)] != p {
		return isSubstringAlternate(s[1:], p)
	}
	return true
}

func main() {
	fmt.Printf("abcdef a**c*e %v \n", findMatchIndex("abcdef", "a**c*e"))
	fmt.Printf("abcdef c*e %v \n", findMatchIndex("abcdef", "c*e"))
	fmt.Printf("abf b*e %v \n", findMatchIndex("abf", "b*e"))
	fmt.Printf("zzzzwewei w* %v \n", findMatchIndex("zzzzwewei", "w*"))
}
/*
Everything is a recursive structure:
strings, arrays, trees, graphs
 */