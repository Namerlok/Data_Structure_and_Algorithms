func checkStrings(s1 string, s2 string) bool {
	var counts [52]int

	for i := 0; i < len(s1); i++ {
		offset := (i & 1) * 26
		counts[offset+int(s1[i])-int('a')]++
		counts[offset+int(s2[i])-int('a')]--
	}

	for _, count := range counts {
		if count != 0 {
			return false
		}
	}

	return true
}