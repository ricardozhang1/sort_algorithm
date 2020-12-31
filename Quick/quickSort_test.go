package quickSort

import (
	"testing"
	"reflect"
)

func TestQuickSort(t *testing.T) {
	b := &quickNode{
		arrayList: []int{3,2,1,5,6,4},
	}

	c := &quickNode{
		arrayList: []int{1,2,3,4,5,6},
	}

	b.sortList()

	result := reflect.DeepEqual(b, c)
	if !result {
		t.Fatal("出现错误！")
	}
}
