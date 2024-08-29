package verificationFunction

import (
	"strings"
)

// this function look if the string given is empty
func IsValidMessage(s string) bool {
	tmp := s

	tmp = strings.ReplaceAll(tmp, " ", "")
	tmp = strings.ReplaceAll(tmp, string(rune(10)), "")
	tmp = strings.ReplaceAll(tmp, string(rune(13)), "")

	return tmp != ""
}
