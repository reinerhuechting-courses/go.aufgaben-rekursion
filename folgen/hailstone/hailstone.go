package hailstone

// Hailstone erwartet eine Zahl n und liefert eine Liste mit den
// Elementen der Hailstone-Folge ab n.
func Hailstone(n int) []int {
	result := []int{n}
	if n == 1 {
		return result
	}
	if n%2 == 0 {
		result = append(result, Hailstone(n/2)...)
	} else {
		result = append(result, Hailstone(3*n+1)...)
	}
	return result
}
