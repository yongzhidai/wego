package lists

import (
	"fmt"
	"testing"
)

type Entity struct {
	name    string
	age     int
	address string
}

func TestLists(t *testing.T) {
	list := []Entity{
		{name: "123", age: 10, address: "山东"},
		{name: "123", age: 20, address: "山东"},
		{name: "456", age: 10, address: "北京"},
	}

	fmt.Println(Arrays(list).Filter(func(i Entity) bool {
		return i.age < 20
	}))

	fmt.Println(Arrays(list).Map(func(i Entity) MapKeyType {
		return i.name
	}))

	fmt.Println(Arrays(list).GroupBy(func(i Entity) MapKeyType {
		return i.address
	}))
}
