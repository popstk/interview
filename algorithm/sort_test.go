package algorithm

import (
	"sort"
	"testing"
)

func testSort(t *testing.T, f func([]int)) {
	data := []int{5, 6, 3, 9, 8, 4, 1, 7, 2}
	cp := make([]int, len(data))
	copy(cp, data)
	sort.Ints(cp)

	f(data)

	t.Log("-> ", data)

	for i := 0; i < len(data); i++ {
		if data[i] != cp[i] {
			t.Error(data)
			break
		}
	}
}

func TestQSort(t *testing.T) {
	testSort(t, QSort)
}

func TestHeapSort(t *testing.T) {
	testSort(t, HeapSort)
}

func TestMergeSort(t *testing.T) {
	testSort(t, MergeSort)
}
