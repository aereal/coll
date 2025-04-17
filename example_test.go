package coll_test

import (
	"fmt"

	"github.com/aereal/coll"
)

func ExampleOrderedSet() {
	nums := coll.NewOrderedSet(3, 1, 2)
	fmt.Printf("nums contains 1?: %v\n", nums.Contains(1))
	fmt.Printf("nums contains 42?: %v\n", nums.Contains(42))
	nums.Append(42)
	fmt.Printf("appended nums contains 42?: %v\n", nums.Contains(42))
	fmt.Print("values:")
	for n := range nums.Values() {
		fmt.Printf(" %d", n)
	}
	fmt.Println()
	// Output:
	// nums contains 1?: true
	// nums contains 42?: false
	// appended nums contains 42?: true
	// values: 3 1 2 42
}
