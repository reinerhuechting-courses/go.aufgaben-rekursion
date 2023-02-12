package hailstone

import "fmt"

func ExampleHailstone() {
	for n := 1; n <= 10; n++ {
		fmt.Println(Hailstone(n))
	}

	// Output:
	// [1]
	// [2 1]
	// [3 10 5 16 8 4 2 1]
	// [4 2 1]
	// [5 16 8 4 2 1]
	// [6 3 10 5 16 8 4 2 1]
	// [7 22 11 34 17 52 26 13 40 20 10 5 16 8 4 2 1]
	// [8 4 2 1]
	// [9 28 14 7 22 11 34 17 52 26 13 40 20 10 5 16 8 4 2 1]
	// [10 5 16 8 4 2 1]
}
