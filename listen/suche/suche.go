package suche

// Search führt eine lineare Suche durch.
// D.h. es durchsucht die Liste von Anfang bis ende nach key und liefert die Position,
// an der Key steht. Liefert die Länge der Liste, falls key nicht enthalten ist.
func Search(list []int, key int) int {
	if len(list) == 0 || list[0] == key {
		return 0
	}
	return 1 + Search(list[1:], key)
}
