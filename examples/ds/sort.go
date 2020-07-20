package main

func main() {
	s := []int{5, 2, 6, 3, 1, 4} // unsorted
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}
