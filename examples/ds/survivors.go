package main

import "fmt"

func main() {
	var initialList []int
	for i := 1; i <= 100; i++ {
		initialList = append(initialList, i)
	}
	i := 0
	skipValue := 0
	for len(initialList) > 1 {
		initialList = splice(initialList, skipValue)
		skipValue = (skipValue + 1 + i) % len(initialList)
		i++
	}
	fmt.Printf("initialList: %v \n",initialList)
}

func splice(list []int, skipValue int) []int {
	return append(list[:skipValue], list[skipValue + 1:]...)
}
