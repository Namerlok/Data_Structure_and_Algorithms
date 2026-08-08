func checkOnesSegment(s string) bool {
	for it := 1; it < len(s); it++ {
		if s[it] == '1' && s[it] != s[it-1] {
			return false
		}
	}
	return true
}