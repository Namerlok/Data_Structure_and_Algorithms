package main

func MaxInt(left, right int) int {
	if left < right {
		return right
	} else {
		return left
	}
}

func addBinary(a string, b string) string {
	result := make([]byte, MaxInt(len(a), len(b))+1)
	posA, posB, posRes := len(a)-1, len(b)-1, len(result)-1
	one := 0
	for posRes >= 0 {
		numA := 0
		if posA >= 0 {
			numA = int(a[posA] - '0')
			posA--
		}
		numB := 0
		if posB >= 0 {
			numB = int(b[posB] - '0')
			posB--
		}
		result[posRes] = byte((numA+numB+one)%2) + '0'
		one = (numA + numB + one) / 2
		posRes--
	}
	if result[0] == '0' {
		return string(result[1:])
	} else {
		return string(result)
	}
}
