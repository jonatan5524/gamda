package gamda

import (
	"fmt"
)

func ExampleAlways() {
	ten := Always(10)

	fmt.Println(ten())
	// Output: 10
}

func ExampleAp() {
	arr := []int{1, 2, 3}
	multiply := func(x int) int { return x * 2 }
	plus := func(x int) int { return x + 3 }

	fmt.Println(Ap([]func(int) int{multiply, plus}, arr))

	// Output: [2 4 6 4 5 6]
}
