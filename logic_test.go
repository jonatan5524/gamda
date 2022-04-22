package gamda

import (
	"fmt"

	"github.com/jonatan5524/gamda/internal/logic"
)

func ExampleAllPass() {
	checkOdd := func(x int) bool { return x%2 != 0 }
	checkGreaterThanFive := func(x int) bool { return x > 5 }
	checkLessThanTwenty := func(x int) bool { return x < 20 }

	checkAllPass := AllPass([]logic.CheckPass[int]{checkOdd, checkGreaterThanFive, checkLessThanTwenty})

	fmt.Println(checkAllPass(7, 9, 11))
	// Output: true
}

func ExampleAnd() {
	ret := And(true, false)
	ret2 := And(true, true)

	fmt.Println(ret, ret2)
	// Output: false true
}
