package sorting

import (
	"reflect"
	"testing"
)

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

func TestMergeSortLoop(t *testing.T) {

	got := MergeSortLoop([]int{8, 54, 9, 1, 16, 209, 64, 3, 40})
	want := []int{1, 3, 8, 9, 16, 40, 54, 64, 209}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, want %v", got, want)
	}

	got = MergeSortLoop([]int{8, 54, 9, 1, 16, 64, 3, 40})
	want = []int{1, 3, 8, 9, 16, 40, 54, 64}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, want %v", got, want)
	}

}

func TestMergeSortReq(t *testing.T) {

	got := MergeSortReq([]int{8, 54, 9, 1, 16, 209, 64, 3, 40})
	want := []int{1, 3, 8, 9, 16, 40, 54, 64, 209}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, want %v", got, want)
	}

	got = MergeSortReq([]int{8, 54, 9, 1, 16, 64, 3, 40})
	want = []int{1, 3, 8, 9, 16, 40, 54, 64}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, want %v", got, want)
	}

}

//cpu: Intel(R) Core(TM) i7-10750H CPU @ 2.60GHz
//BenchmarkMergeSortLoop-12    	      37	  30609189 ns/op	62305231 B/op	  400418 allocs/op
func BenchmarkMergeSortLoop(b *testing.B) {

	for n := 0; n < b.N; n++ {
		MergeSortLoop(Arr)
	}
}

//BenchmarkMergeSort3-12    	      32	  36297773 ns/op	78269924 B/op	  762564 allocs/op
func BenchmarkMergeSortReq(b *testing.B) {

	for n := 0; n < b.N; n++ {
		MergeSortReq(Arr)
	}
}
