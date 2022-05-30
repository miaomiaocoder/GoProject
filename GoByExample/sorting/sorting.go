package main

import (
	"fmt"
	"sort"
)

func main() {
	strs := []string{"c", "a", "b"}
	sort.Strings(strs)
	fmt.Println("Strings:", strs)

	ints := []int{7, 2, 4}

	// 检查是否有序
	s := sort.IntsAreSorted(ints)
	fmt.Println("Sorted:	", s)

	sort.Ints(ints)
	fmt.Println("Ints:	", ints)

	// 检查是否有序
	s2 := sort.IntsAreSorted(ints)
	fmt.Println("Sorted:	", s2)
}
