package listenvergleich

// Liefert true, falls die beiden Listen gleich sind.
func ListsEqual(list1, list2 []int) bool {
	if len(list1) != len(list2) {
		return false
	}
	if len(list1) == 0 {
		return true
	}
	return list1[0] == list2[0] && ListsEqual(list1[1:], list2[1:])
}
