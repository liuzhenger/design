package design

import (
	"fmt"
	"testing"
)

func Test_17_1(t *testing.T) {
	aggregate := ConcreteAggregate{}
	aggregate.Add("a")
	aggregate.Add("b")
	aggregate.Add("c")
	iterator := aggregate.CreateIterator()
	for iterator.HasNext() {
		fmt.Println(iterator.Next())
	}
}
