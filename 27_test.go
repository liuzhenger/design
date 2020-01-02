package design

import (
	"fmt"
	"testing"
)

func Test_27_1(t *testing.T) {
	eleA := new(ElementA)
	eleA.Name = "文件A"
	eleA.LookPerms = []string{"read"}
	eleB := new(ElementB)
	eleB.Name = "文件B"

	root := new(VisitorRoot)
	one := new(Visitor)

	data := new(Data)
	data.Add(eleA)
	data.Add(eleB)

	data.Show(root)
	data.Show(one)
	fmt.Println("----------")
	one.Perm = "read"
	data.Show(one)
}
