package shell

import (
	"fmt"
)

type ShellNode struct {
	arrayList []int
}

func (s *ShellNode) sortList() {
	n := len(s.arrayList)
	h := 1
	for h<n/2 {
		h = h*2+1
	}
	for h>=1 {
		// fmt.Printf("h = %v\n", h)
		for i:=h; i<n; i++ {
			// fmt.Println("--111", i)
			for j:=i; j>=h; j-=h {
				// fmt.Println(j, j-h)
				if greater(s.arrayList[j-h], s.arrayList[j]) {
					s.exch(j, j-h)
				} else {
					break
				}
			}
		}
		h /= 2
	}
	// fmt.Println(h)
}

func greater(a, b int) bool {
	if a > b {
		return true
	} else {
		return false
	}
}

func (s *ShellNode) exch(i, j int) {
	temp := s.arrayList[i]
	s.arrayList[i] = s.arrayList[j]
	s.arrayList[j] = temp
}

func main() {
	b := &ShellNode{
		arrayList: []int{89,4,5,6,9,1,2,3},
	}
	b.sortList()
	fmt.Println(b)
}

