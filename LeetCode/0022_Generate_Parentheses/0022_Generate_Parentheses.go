func ConstructionBrackets(result *[]string, pref string, open int, close int, n int) {
	if open == n && close == n {
		*result = append(*result, pref)
	}

	if open < n {
		ConstructionBrackets(result, pref+"(", open+1, close, n)
	}
	if open > close {
		ConstructionBrackets(result, pref+")", open, close+1, n)
	}
}

func generateParenthesis(n int) []string {
	if n == 0 {
		return []string{""}
	}
	var result []string
	ConstructionBrackets(&result, "", 0, 0, n)
	return result
}
