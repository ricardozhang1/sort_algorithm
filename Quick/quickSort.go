package quickSort

type quickNode struct{
	arrayList []int
}

func (q *quickNode) sortList() {
	// lo hi 下标元素
	lo := 0
	hi := len(q.arrayList) - 1
	sortted(&q.arrayList, lo, hi)
}

func sortted(q *[]int, lo, hi int) {
	if hi <= lo {
		return
	}
	// 对q数组中，从lo到hi元素进行切分
	part := partition(q, lo, hi)
	// 对左边分组中的元素进行排序
	// 对右边分组中的元素进行排序
	sortted(q, lo, part-1)
	sortted(q, part+1, hi)
}

func partition(q *[]int, lo, hi int) int {
	key := (*q)[lo]  // 把最左边的元素当做基准值
	left := lo    // 定义一个左侧指针，初始指向最左边的元素
	right := hi // 定义一个右侧指针，初始指向最右边的元素
	// 进行切分
	for {
		// 先从右往左扫描，找到一个比基准值小的元素
		for {
			if (*q)[right] < key {
				break
			}else if right == lo {
				break
			}else {
				right--
			}
		}


		// 再从左往右扫描，找一个比基准值大的元素
		for {
			if (*q)[left] > key {
				break
			}else if left == hi {
				break
			}else {
				left++
			}
		}

		if left >= right {
			// 扫描完了所有元素
			break
		}else {
			// 交换left和right索引出的元素
			exch(q, left, right)
		}
	}

	// 交换最后的right索引处和基准处的索引值
	exch(q, lo, right)

	// right就是切分的界限
	return right
}

func greater(a, b int) bool {
	if a > b {
		return true
	}else {
		return false
	}
}

func exch(s *[]int, v, w int) {
	temp := (*s)[v]
	(*s)[v] = (*s)[w]
	(*s)[w] = temp
}


// func main() {
// 	b := &quickNode{
// 		arrayList: []int{6, 1, 2, 7, 9, 3, 4, 5, 8},
// 	}
// 	b.sortList()
// 	fmt.Println(b)
// }
