package selection

import (
	"fmt"
)

type SelectionNode struct {
	arrayList []int
}

func (s *SelectionNode) sortList() {
	for i:=0; i<len(s.arrayList)-1; i++ {
		// fmt.Println("i = ", i)
		minNum := i
		for j:=i+1; j<len(s.arrayList); j++ {
			if !greater(s.arrayList[j], s.arrayList[minNum]) {
				minNum = j
			}
		}
		// fmt.Println("minNum: ", minNum, "i:", i)
		s.exch(minNum, i)
	}
}

func greater(a, b int) bool {
	if a > b {
		return true
	}else {
		return false
	}
}

func (s *SelectionNode) exch(v, w int) {
	temp := s.arrayList[v]
	s.arrayList[v] = s.arrayList[w]
	s.arrayList[w] = temp
}

func main() {
	b := &SelectionNode{
		arrayList: []int{89,4,5,6,9,1,2,3},
	}
	b.sortList()
	fmt.Println(b)
}



