package strings

// Contains prüft, ob der String s die Sequenz seq enthält.
func Contains(s, seq string) bool {
	if seq == "" {
		return true
	}
	if s == "" {
		return false
	}
	return StartsWith(s, seq) || Contains(s[1:], seq)
}
