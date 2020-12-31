package shell

import (
	"testing"
	"reflect"
)

func TestShellSort(t *testing.T) {
	b := &ShellNode{
		arrayList: []int{89,4,5,6,9,1,2,3},
	}

	bc := &ShellNode{
		arrayList: []int{1,2,3,4,5,6,9,89},
	}

	// fmt.Println(b)
	b.sortList()

	result := reflect.DeepEqual(b, bc)
	if !result {
		t.Fatal("出现错误！")
	}
}




