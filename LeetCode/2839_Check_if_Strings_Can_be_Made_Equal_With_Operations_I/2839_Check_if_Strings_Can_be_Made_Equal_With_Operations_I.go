func canBeEqual(s1 string, s2 string) bool {
	setF := make(map[byte]int, 2)
	setF[s1[0]]++
	setF[s1[2]]++
	if el, ok := setF[s2[0]]; !ok || el == 0 {
		return false
	} else {
		setF[s2[0]] = el - 1
	}
	if el, ok := setF[s2[2]]; !ok || el == 0 {
		return false
	} else {
		setF[s2[2]] = el - 1
	}

	setS := make(map[byte]int, 2)
	setS[s1[1]]++
	setS[s1[3]]++
	if el, ok := setS[s2[1]]; !ok || el == 0 {
		return false
	} else {
		setS[s2[1]] = el - 1
	}
	if el, ok := setS[s2[3]]; !ok || el == 0 {
		return false
	} else {
		setS[s2[3]] = el - 1
	}
	return true
}