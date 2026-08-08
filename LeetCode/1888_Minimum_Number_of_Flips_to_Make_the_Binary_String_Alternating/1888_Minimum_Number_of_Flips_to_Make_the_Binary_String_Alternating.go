func minInt(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

func minFlips(s string) int {
	var minFlip, curFlipOne, curFlipZero int
	for pos := 0; pos < len(s); pos++ {
		if int(s[pos]-'0') == pos%2 {
			curFlipOne++
		} else {
			curFlipZero++
		}
	}
	minFlip = minInt(curFlipZero, curFlipOne)
	for div := 0; div < len(s); div++ {
		if int(s[div]-'0') == div%2 {
			curFlipOne--
		} else {
			curFlipZero--
		}
		if int(s[div]-'0') == (len(s)+div)%2 {
			curFlipOne++
		} else {
			curFlipZero++
		}
		minFlip = minInt(minFlip, minInt(curFlipZero, curFlipOne))
	}
	return minFlip
}