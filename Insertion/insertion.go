package insertion

import (
	"fmt"
)

type InsertionNode struct{
	arrayList []int
}


func (s *InsertionNode) sortList() {
	for i:=1; i<len(s.arrayList); i++ {
		fmt.Println(i)
		for j:=i; j>0; j-- {
			fmt.Println(j)
			if greater(s.arrayList[j-1], s.arrayList[j]) {
				s.exch(j-1, j)
			}
		}
	}
}


func greater(a, b int) bool {
	if a > b {
		return true
	}else {
		return false
	}
}

func (s *InsertionNode) exch(v, w int) {
	temp := s.arrayList[v]
	s.arrayList[v] = s.arrayList[w]
	s.arrayList[w] = temp
}

func main() {
	b := &InsertionNode{
		arrayList: []int{89,4,5,6,9,1,2,3},
	}
	b.sortList()
	fmt.Println(b)
}


