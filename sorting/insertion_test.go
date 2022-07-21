package sorting

import (
	"reflect"
	"testing"
)

func TestInsertionSort(t *testing.T) {

	unsorted := []int{8, 54, 9, 1, 16}

	assertSort := func(t testing.TB, got []int, want []int) {
		t.Helper()

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v, want %v", got, want)
		}
	}

	t.Run("asc", func(t *testing.T) {
		want := []int{1, 8, 9, 16, 54}
		got, _ := InsertionSort(unsorted, "asc")

		assertSort(t, got, want)

	})

	t.Run("desc", func(t *testing.T) {

		want := []int{54, 16, 9, 8, 1}
		got, _ := InsertionSort(unsorted, "desc")

		assertSort(t, got, want)
	})

	t.Run("wrong order argument", func(t *testing.T) {

		got, err := InsertionSort(unsorted, "wrong order argument")
		want := []int{}
		if err == nil {
			t.Error("wanted an error but didn't get one")
		}
		assertSort(t, got, want)
	})
}
