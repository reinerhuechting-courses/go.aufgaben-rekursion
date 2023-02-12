package suche

// BinSearch geht davon aus, dass die Liste aufsteigend sortiert ist und führt
// eine sog. "binäre Suche" durch.
// D.h. die Funktion vergleicht zuerst das mittlere Element mit key.
// Falls key dort gefunden wurde, wird die Position zurückgeliefert.
// Ansonsten wird nur im linken oder rechten Teil weitergesucht.
// Ist das Element nicht enthalten, wird die Länge der Liste geliefert.
func BinSearch(list []int, key int) int {
	if len(list) == 0 {
		return 0
	}

	m := len(list) / 2
	if list[m] == key {
		return m
	}

	if key < list[m] {
		pos := Search(list[:m], key)
		if pos == len(list[:m]) {
			return len(list)
		}
		return pos
	}
	return m + Search(list[m+1:], key) + 1
}
