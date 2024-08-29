package verificationFunction

// this fonction look if a string is in a slice of string
func TabNotContain(tab []string, s string) bool {
	for _, v := range tab {
		if s == v {
			return false
		}
	}

	return true
}

// this fonction look if an int is in a slice of int
func TabNotContainInt(tab []int, s int) bool {
	for _, v := range tab {
		if s == v {
			return false
		}
	}

	return true
}