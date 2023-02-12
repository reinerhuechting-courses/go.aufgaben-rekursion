package strings

// ContainsChain liefert true, falls s eine Kette von count aufeinander folgenden
// Vorkommen von symbol enthält.
func ContainsChain(s, symbol string, count int) bool {
	if count == 0 {
		return true
	}
	if s == "" {
		return symbol == ""
	}
	return (s[:1] == symbol && StartsWith(s[:1], Chain(symbol, count-1))) || ContainsChain(s[1:], symbol, count)
}
