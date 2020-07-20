package main

import "fmt"

/*
Josephus Problem: https://www.geogebra.org/m/ExvvrBbR
There are n soldiers in a circle going from 1 to n. The 1st soldier kills the one next to it
The 3rd soldier kills the one next to it and so on till only one soldier is left
Find the position of this soldier

In mathematical proof:
n = 2^m + l
where l < 2^m (i.e. 2^m is the highest power of 2 such that 2^m < n)
then, f(n) = 2l + 1
 */

func main() {
	var initialList []int
	for i := 1; i <= 6; i++ {
		initialList = append(initialList, i)
	}
	i := 0
	skipValue := i + 1
	for len(initialList) > 1 {
		initialList = splice(initialList, skipValue)
		// i was the last value before splice, so for the next round it is the 0th i
		if i == len(initialList) {
			i = 0
		} else {
			i = (i + 1) % len(initialList)
		}
		skipValue = (i + 1) % len(initialList)
	}
	fmt.Printf("safe position is %v \n", initialList[0])
}

func splice(list []int, skipValue int) []int {
	return append(list[:skipValue], list[skipValue + 1:]...)
}
