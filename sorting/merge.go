package sorting

/*
merge sort is valuable for very large sets of
data because it divides the data in predictable ways. This allows us to divide the
data into more manageable pieces ourselves, use merge sort to sort them, and
then perform as many merges as necessary without having to keep the entire set
of data in memory all at once.
*/

//non-reqursive
func MergeSortLoop(in []int) (out []int) {

	var separated [][]int

	//break an input slice into a slice of slices, each length 1
	for _, elm := range in {

		i := []int{elm}
		separated = append(separated, i)
	}

	//until we have only 1 slice in our slice of slices:
	//- cut off the first two elements
	//- merge them
	//- append our slice of slices with the result
	for len(separated) > 1 {

		merged := Merge(separated[0], separated[1])
		separated = separated[2:]
		separated = append(separated, merged)

	}

	return separated[0]
}

//reqursive
func MergeSortReq(list []int) (sorted []int) {

	if len(list) > 1 {

		j := len(list) / 2
		return Merge(MergeSortReq(list[:j]), MergeSortReq(list[j:]))
	} else {
		return list
	}

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
