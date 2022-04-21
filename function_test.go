package gamda

import (
	"fmt"
)

func ExampleAlways() {
	ten := Always(10)

	fmt.Println(ten())
	// Output: 10
}
