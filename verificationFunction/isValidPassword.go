package verificationFunction

// this function look if the string given is with a good format
// 1 UpperCase
// 1 Number
// 1 Special Character
// minimum size of 8 characters
func PasswordVerif(password string) bool {
	var isLongEnought bool = false
	var containUpper bool = false
	var containSpeChar bool = false
	var containNumber bool = false
	if len(password) >= 8 {
		isLongEnought = true
	}
	for _, r := range password {
		if r >= 'A' && r <= 'Z' {
			containUpper = true
		} else if r >= '0' && r <= '9' {
			containNumber = true
		} else if r < 'a' || r > 'z' {
			containSpeChar = true
		}
	}
	if isLongEnought && containNumber && containSpeChar && containUpper {
		return true
	}
	return false
}
