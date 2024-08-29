package verificationFunction

import (
	"net/mail"
)

// this function look if the email format is good
func EmailVerif(email string) bool {
	_, err := mail.ParseAddress(email)
	return err == nil
}
