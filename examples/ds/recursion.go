package main

import "fmt"

var arr []string
func searchWords(str string, count int) {
	if len(str) == 0 {
		return
	}
	arr = append(arr, str)
	searchWords(str[1:], count + 1)
	fmt.Printf("arr: %v, count: %v \n", arr, count)
}

func main() {
	arr = make([]string, 0)
	searchWords("fine", 0)
}
