package elemententfernen

// Liefert eine Liste, die alle Elemente aus list enthält,
// außer dem an Stelle pos.
func RemoveElement(list []int, pos int) []int {
	if pos < 0 || len(list) == 0 {
		return list
	}
	if pos == 0 {
		return list[1:]
	}
	return append(list[0:1], RemoveElement(list[1:], pos-1)...)
}
