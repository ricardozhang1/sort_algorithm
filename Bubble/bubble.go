package bubble

import (
	"fmt"
)

type bubbles struct{
	arrayList []int
}


func (b *bubbles) sortList() {
	for i:=len(b.arrayList)-1; i>0; i-- {
		for j:=0; j<i; j++ {
			if greater(b.arrayList[j], b.arrayList[j+1]) {
				b.exch(j, j+1)
			}
		}
	}
}

func greater(a, b int) bool {
	if a > b {
		return true
	}else{
		return false
	}
}

func (b *bubbles) exch(v, w int) {
	tmp := b.arrayList[v]
	b.arrayList[v] = b.arrayList[w]
	b.arrayList[w] = tmp
}


