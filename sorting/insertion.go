package sorting

import (
	"errors"
)

//InsertionSort sorts a slice of integers with O(n log n) efficiency..
//Pass "asc" or "desc" in 2nd argument string for ascending or descending order.
func InsertionSort(data []int, order string) ([]int, error) {

	var eval func(int, int, []int) bool
	var err error

	if order == "asc" {
		eval = func(a, b int, x []int) bool { return a < x[b-1] }
	} else if order == "desc" {
		eval = func(a, b int, x []int) bool { return a > x[b-1] }
	} else {
		return []int{}, errors.New("Invalid order (must be asc or desc)")
	}

	if len(data) > 1 {

		for i, elm := range data {

			for j := i; j > 0; j-- {

				if eval(elm, j, data) {
					data[j], data[j-1] = data[j-1], data[j]
				}
			}
		}
	}
	return data, err
}
