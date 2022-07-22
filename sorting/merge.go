package sorting

func MergeSort(list []int) (sorted []int) {

	var left = []int{}
	var right = []int{}

	if len(list) > 1 {

		i := 0
		j := len(list) - 1

		for i <= j {

			if i < j {
				left = append(left, list[i])
				right = append(right, list[j])
				i++
				j--
			} else if i == j {
				left = append(left, list[i])
				i++
			}
		}
		sorted = Merge(MergeSort(left), MergeSort(right))
	} else {
		sorted = list
	}

	return sorted
}

func Merge(left []int, right []int) (merged []int) {

	i := 0
	j := 0

	for len(left) > i {

		if len(right) > j {

			if left[i] <= right[j] {

				merged = append(merged, left[i])
				i++

			} else {
				merged = append(merged, right[j])
				j++
			}

		} else {

			merged = append(merged, left[i])
			i++
		}

	}

	if len(right) > j {

		for len(right) > j {
			merged = append(merged, right[j])
			j++
		}

	}

	return
}
