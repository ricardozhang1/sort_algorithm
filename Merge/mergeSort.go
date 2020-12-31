package main

import (
	"fmt"
)

type MergeNode struct {
	arrayList []int
}

var assist = make([]int, 8)

func mergeSort(arr []int, start, end int) {
	if start >= end {
		return
	}
	mid := (start + end)/2
	mergeSort(arr, start, mid)
	mergeSort(arr, mid+1, end)
	merge(arr, start, mid, end)
}

// conquer
func merge(arr []int, start, mid, end int) {
	i, p1, p2 := start, start, mid+1
	for p1 <= mid && p2 <= end {
		if !greater(arr[p1], arr[p2]) {
			assist[i] = arr[p1]
			p1++
		} else {
			assist[i] = arr[p2]
			p2++
		}
		i++
	}
	for p1<=mid {
		assist[i] = arr[p1]
		p1++
	}
	for p2<=end {
		assist[i] = arr[p2]
		p2++
	}
	for index:=start; index<=end; index++ {
		arr[index] = assist[index]
	}
}

func (m *MergeNode) sortList() {
	mergeSort(m.arrayList, 0, len(m.arrayList)-1)
}

func greater(a, b int) bool {
	if a > b {
		return true
	} else {
		return false
	}
}

func (s *MergeNode) exch(i, j int) {
	temp := s.arrayList[i]
	s.arrayList[i] = s.arrayList[j]
	s.arrayList[j] = temp
}

func main() {
	b := &MergeNode{
		arrayList: []int{89,4,5,6,9,1,2,3},
	}
	b.sortList()
	fmt.Println(b)
}

