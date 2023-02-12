package filter

// Liefert eine Liste mit allen Elementen aus list, die kleiner oder gleich key sind.
func FilterLess(list []int, key int) []int {
	if len(list) == 0 {
		return []int{}
	}
	head, tail := list[0], FilterLess(list[1:], key)
	if head <= key {
		return append([]int{head}, tail...)
	}
	return tail
}

// Liefert eine Liste mit allen Elementen aus list, die größer als key sind.
func FilterGreater(list []int, key int) []int {
	if len(list) == 0 {
		return []int{}
	}
	head, tail := list[0], FilterGreater(list[1:], key)
	if head > key {
		return append([]int{head}, tail...)
	}
	return tail
}
