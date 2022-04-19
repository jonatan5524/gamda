package gamda

import "fmt"

func ExampleAdjust() {
	arr := []int{1, 2, 3, 4}

	increaseElem := func(elem int) int {
		return elem + 1
	}

	newArr := Adjust(1, increaseElem, arr)

	fmt.Println(newArr)
	// Output: [1 3 3 4]
}

func ExampleFilter() {
	arr := []int{1, 2, 3, 4, 5, 6}

	filterByEven := func(elem int) bool {
		return elem%2 == 0
	}

	filteredArr := Filter(filterByEven, arr)

	fmt.Println(filteredArr)
	// Output: [2 4 6]
}

func ExampleFind() {
	type twoNums struct {
		numA int
		numB int
	}

	arr := []twoNums{
		{
			numA: 1,
			numB: 2,
		},
		{
			numA: 3,
			numB: 4,
		},
		{
			numA: 5,
			numB: 6,
		},
	}

	numBFour := func(elem twoNums) bool {
		return elem.numB == 4
	}

	objFind := Find(numBFour, arr)

	fmt.Println(objFind)
	// Output: {3 4}
}

func ExampleMap() {
	arr := []int{1, 2, 3, 4}

	doubleElems := func(elem int) int {
		return elem * 2
	}

	mappedArr := Map(doubleElems, arr)

	fmt.Println(mappedArr)
	// Output: [2 4 6 8]
}

func ExampleAny() {
	arr := []int{1, 2, 3, 5}

	evenElem := func(elem int) bool {
		return elem%2 == 0
	}

	mappedArr := Any(evenElem, arr)

	fmt.Println(mappedArr)
	// Output: true
}

func ExampleAll() {
	arr := []int{2, 4, 6, 8}

	evenElem := func(elem int) bool {
		return elem%2 == 0
	}

	mappedArr := All(evenElem, arr)

	fmt.Println(mappedArr)
	// Output: true
}

func ExampleForeach() {
	arr := []int{1, 2, 3, 4}

	newArr := []int{}
	doubleSlice := func(elem int) {
		newArr = append(newArr, elem*2)
	}

	Foreach(doubleSlice, arr)

	fmt.Println(newArr)
	// Output: [2 4 6 8]
}

func ExampleConcat() {
	arr := []int{1, 2, 3}
	secondArr := []int{3, 5, 6}

	combinedArr := Concat(arr, secondArr)

	fmt.Println(combinedArr)
	// Output: [1 2 3 3 5 6]
}
