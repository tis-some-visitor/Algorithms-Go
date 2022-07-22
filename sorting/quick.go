package sorting

import (
	"errors"
	"math/rand"
)

//sorts an integer slice
func QuickSort(list []int, order string) []int {

	result := Sort(list)

	if order == "desc" {
		return Reverse(list)
	}

	return result
}

//InsertionSort sorts a slice of integers with O(log n) efficiency..
//Pass "asc" or "desc" in 2nd argument string for ascending or descending order.
func Sort(list []int) []int {

	if len(list) > 1 {

		var x int

		if len(list) > 100 {
			x = MedianOfThree(list)
		} else {
			x = list[rand.Intn(len(list)-1)]
		}

		//x := list[rand.Intn(len(list)-1)]

		i := 0
		k := len(list) - 1

		for list[i] != list[k] {

			if list[i] < x {
				i++
			} else if list[i] >= x && list[k] > x {
				k--
			} else if list[i] >= x && list[k] <= x {
				list[i], list[k] = list[k], list[i]
			}
		}

		Sort(list[:i])
		Sort(list[i+1:])
	}

	return list
}

func Reverse(list []int) []int {

	i := 0
	k := len(list) - 1

	for i < k {
		list[i], list[k] = list[k], list[i]
		i++
		k--
	}
	return list
}

//InsertionSort sorts a slice of integers with O(log n) efficiency..
//Pass "asc" or "desc" in 2nd argument string for ascending or descending order.
func QuickSortWithFunc(list []int, order string) ([]int, error) {

	var left func(int, int) bool
	var right func(int, int) bool

	if order == "asc" {
		left = func(a, b int) bool { return a < b }
		right = func(a, b int) bool { return a >= b }
	} else if order == "desc" {
		left = func(a, b int) bool { return a > b }
		right = func(a, b int) bool { return a <= b }
	} else {
		return []int{}, errors.New("Invalid order (must be asc or desc)")
	}

	if len(list) > 1 {

		var x int

		if len(list) > 1000 {
			x = MedianOfThree(list)
		} else {
			x = list[rand.Intn(len(list)-1)]
		}

		i := 0
		k := len(list) - 1

		for list[i] != list[k] {

			if left(list[i], x) {
				i++
			} else if right(list[i], x) && left(x, list[k]) {
				k--
			} else if right(list[i], x) && right(x, list[k]) {
				list[i], list[k] = list[k], list[i]
			}
		}

		QuickSort(list[:i], order)
		QuickSort(list[i+1:], order)
	}

	return list, nil
}

func MedianOfThree(list []int) int {

	a := list[rand.Intn(len(list)-1)]
	b := list[rand.Intn(len(list)-1)]
	c := list[rand.Intn(len(list)-1)]

	return CompairThree(a, b, c)
}

func CompairThree(a, b, c int) int {

	if (a >= b && b >= c) || (a <= b && b <= c) {
		return b
	} else if (b >= a && a >= c) || (b <= a && a <= c) {
		return a
	} else {
		return c
	}
}

func GenerateArr(max int) (res []int) {

	for i := 0; i < 100000; i++ {
		res = append(res, rand.Intn(max))
	}
	return
}
