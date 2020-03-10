package questions

import (
	"math/rand"
)

type LessFunc func(interface{}, interface{}) bool

func Quicksort(array []int, lessFunc LessFunc) {
	if len(array) < 2 {
		return
	}
	left, right := 0, len(array)-1
	pivot := rand.Int() % len(array)

	swap(array, pivot, right)
	for i := range array {
		if lessFunc(array[i], array[right]) {
			swap(array, i, left)
			left++
		}
	}
	swap(array, left, right)
	Quicksort(array[:left], lessFunc)
	Quicksort(array[left+1:], lessFunc)
}

func swap(array []int, i, j int) {
	temp := array[j]
	array[j] = array[i]
	array[i] = temp
}
