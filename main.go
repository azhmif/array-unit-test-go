package main

import "fmt"

var (
	a        = []int{1, 2, 3, 8, 9, 3, 2, 1}
	b        = []int{1, 2, 1, 2, 2, 1}
	c        = []int{7, 1, 2, 9, 7, 2, 1}
	maxA int = 3
	maxB int = 2
	maxC int = 2
)

func main() {
	maxA := findMaxRow(a)
	fmt.Println("Max: ", maxA)
	maxB := findMaxRow(b)
	fmt.Println("Max: ", maxB)
	maxC := findMaxRow(c)
	fmt.Println("Max: ", maxC)

}

func findMaxRow(a []int) (max int) {
	var nextValue int
	var rangeValue int
	var rangeValueTemp int
	for _, value := range a {
		if value == nextValue {
			if rangeValueTemp > rangeValue {
				rangeValue = rangeValueTemp
				max = value
			}
		} else {
			rangeValueTemp = 0
		}
		nextValue = value + 1
		rangeValueTemp++
	}
	return max
}
