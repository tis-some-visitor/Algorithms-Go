package sorting

import (
	"reflect"
	"testing"
)

func TestMergeSort(t *testing.T) {

	assertResult := func(t testing.TB, got []int, want []int) {
		t.Helper()
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v, want %v", got, want)
		}
	}

	got := MergeSort([]int{8, 54, 9, 1, 16, 209, 64, 3, 40})
	want := []int{1, 3, 8, 9, 16, 40, 54, 64, 209}
	assertResult(t, got, want)

	got = MergeSort([]int{8, 54, 9, 1, 16, 64, 3, 40})
	want = []int{1, 3, 8, 9, 16, 40, 54, 64}
	assertResult(t, got, want)

	got = MergeSort([]int{8})
	want = []int{8}
	assertResult(t, got, want)
}

func TestMerge(t *testing.T) {

	assertResult := func(t testing.TB, got []int, want []int) {
		t.Helper()
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v, want %v", got, want)
		}
	}

	got := Merge([]int{8, 9}, []int{16, 40})
	want := []int{8, 9, 16, 40}
	assertResult(t, got, want)

	got = Merge([]int{8, 9, 89}, []int{16, 40})
	want = []int{8, 9, 16, 40, 89}
	assertResult(t, got, want)

	got = Merge([]int{8}, []int{})
	want = []int{8}
	assertResult(t, got, want)
}

func BenchmarkMergeSort(b *testing.B) {

	for n := 0; n < b.N; n++ {
		MergeSort(Arr)
	}
}
