package main

func numSteps(s string) int {
	var count int
	var increes int
	for i := len(s) - 1; i > 0; i-- {
		if (int(s[i]-'0')+increes)%2 == 1 {
			count++
		}
		increes = (int(s[i]-'0') + increes + 1) / 2
		count++
	}
	if increes != 0 {
		count++
	}
	return count
}
