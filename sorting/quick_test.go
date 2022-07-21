package sorting

import (
	"reflect"
	"testing"
)

func TestQuickOrder(t *testing.T) {

	assertResult := func(t testing.TB, got []int, want []int) {
		t.Helper()
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v, want %v", got, want)
		}
	}

	got, _ := QuickSortWithFunc([]int{8, 54, 9, 1, 16, 209, 64, 3, 40}, "asc")
	want := []int{1, 3, 8, 9, 16, 40, 54, 64, 209}
	assertResult(t, got, want)

	got, _ = QuickSortWithFunc([]int{8, 54, 9, 1, 16, 209, 64, 3, 40}, "desc")
	want = []int{209, 64, 54, 40, 16, 9, 8, 3, 1}
	assertResult(t, got, want)

	got = Sort([]int{8, 64, 9, 1, 16, 209, 64, 3, 40})
	want = []int{1, 3, 8, 9, 16, 40, 64, 64, 209}
	assertResult(t, got, want)

	got = Sort([]int{8, 64, 9, 1, 16, 209, 64, 3, 40})
	want = []int{1, 3, 8, 9, 16, 40, 64, 64, 209}
	assertResult(t, got, want)

	got = Reverse([]int{3, 8, 9, 16, 40, 54, 64, 209})
	want = []int{209, 64, 54, 40, 16, 9, 8, 3}
	assertResult(t, got, want)

	got = QuickSort([]int{8, 54, 9, 1, 16, 209, 64, 3, 40}, "asc")
	want = []int{1, 3, 8, 9, 16, 40, 54, 64, 209}
	assertResult(t, got, want)

	got = QuickSort([]int{8, 54, 9, 1, 16, 209, 64, 3, 40}, "desc")
	want = []int{209, 64, 54, 40, 16, 9, 8, 3, 1}
	assertResult(t, got, want)
}

func TestCompairThree(t *testing.T) {
	assert := func(t testing.TB, got int, want int) {
		t.Helper()
		if got != want {
			t.Errorf("got %v, want %v", got, want)
		}
	}
	got := CompairThree(1, 2, 3)
	want := 2
	assert(t, got, want)

	got = CompairThree(208, 11, 75)
	want = 75
	assert(t, got, want)

	got = CompairThree(3, 3, 2)
	want = 3
	assert(t, got, want)
}

var Arr = GenerateArr(100000)

//cpu: Intel(R) Core(TM) i7-10750H CPU @ 2.60GHz
//BenchmarkQuick-12    	     147	   7849662 ns/op	       0 B/op	       0 allocs/op
func BenchmarkQuickSortWithFunc(b *testing.B) {

	for n := 0; n < b.N; n++ {
		QuickSortWithFunc(Arr, "asc")
	}
}

//cpu: Intel(R) Core(TM) i7-10750H CPU @ 2.60GHz
//BenchmarkSort2-12    	     376	   3017635 ns/op	       0 B/op	       0 allocs/op
func BenchmarkQuickSortAsc(b *testing.B) {

	for n := 0; n < b.N; n++ {
		QuickSort(Arr, "asc")
	}

}

//cpu: Intel(R) Core(TM) i7-10750H CPU @ 2.60GHz
//BenchmarkSort22-12    	     190	   7306411 ns/op	       0 B/op	       0 allocs/op
func BenchmarkQuickSortDesc(b *testing.B) {

	for n := 0; n < b.N; n++ {
		QuickSort(Arr, "desc")
	}
}
