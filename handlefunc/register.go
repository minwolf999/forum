package handlefunc

import (
	"errors"
	"forum/cookies"
	"forum/database/controller/users"
	"forum/structure"
	"forum/verificationFunction"
	"net/http"
	"strings"

	"golang.org/x/crypto/bcrypt"
)

// this function manages the registration
func RegisterHandle(w http.ResponseWriter, r *http.Request) error {
	var user structure.Users
	code := strings.Split(r.URL.String(), "?")[len(strings.Split(r.URL.String(), "?"))-1]

	if r.FormValue("userContent") == "" && r.FormValue("FacebookUserContent") == "" && !strings.HasPrefix(code, "code=") {
		if r.FormValue("psw") != r.FormValue("confirm-psw") {
			return errors.New("password and password confirmation don't match")
		}

		username := r.FormValue("username")
		email := r.FormValue("email")
		password := r.FormValue("psw")

		if !verificationFunction.PasswordVerif(password) {
			return errors.New("incorrect password ! the password must contain 8 characters, 1 uppercase letter, 1 special character, 1 number")
		}

		if !verificationFunction.EmailVerif(email) {
			return errors.New("incorrect email format")
		}

		if username == "" || email == "" || password == "" {
			return errors.New("there is an empty field")
		}

		cryptPassword, _ := bcrypt.GenerateFromPassword([]byte(password), 12)
		err := users.AddUser(username, email, string(cryptPassword))
		if err != nil {
			return errors.New("username or email is already used by someone")
		}

		user, err = users.GetUser(email, string(cryptPassword))
		if err != nil {
			return err
		}

	} else if r.FormValue("userContent") != "" {
		email, username, _, err := GoogleLogDecoder(r.FormValue("userContent"))
		if err != nil {
			return err
		}

		err = users.AddUser(username, email, "")
		if err != nil {
			return errors.New("username or email is already used by someone")
		}

		structure.Html.User = structure.Users{}
		user, err = users.GetUserByEmail(email)
		if err != nil {
			return err
		}
	}

	users.ConnectUser(structure.Html.User.Id)
	user.Connected = 1
	structure.Html.User = user

	cookies.AddCookies(w)
	return nil
}
