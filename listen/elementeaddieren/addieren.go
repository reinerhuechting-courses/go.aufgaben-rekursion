package elementeaddieren

// Liefert die Summe aller Elemente in list.
func AddElements(list []int) int {
	if len(list) == 0 {
		return 0
	}
	return list[0] + AddElements(list[1:])
}
