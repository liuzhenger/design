package design

import (
	"fmt"
	"sort"
	"testing"
)

func TestSort(t *testing.T) {
	arr := []int{2, 1, 5, 4}
	sort.Ints(arr)
	fmt.Println(arr)
}

func Test_25_1(t *testing.T) {
	s1 := new(MinSort)
	s2 := new(BigSort)

	c := new(Context)
	c.Init(s1)

	arr := []int{2, 4, 2, 1, 5, 6, 9, 8, 0}
	c.DoSort(arr)
	fmt.Println(arr)

	c.Init(s2)
	c.DoSort(arr)
	fmt.Println(arr)
}
