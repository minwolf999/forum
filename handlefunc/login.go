package handlefunc

import (
	"encoding/json"
	"errors"
	"forum/cookies"
	"forum/database/controller/users"
	"forum/structure"
	"net/http"
	"strings"

	"golang.org/x/crypto/bcrypt"
)

// this function manages the login
func LogInHandle(w http.ResponseWriter, r *http.Request) error {
	structure.Html.User = structure.Users{}

	code := strings.Split(r.URL.String(), "?")[len(strings.Split(r.URL.String(), "?"))-1]

	var (
		email    string
		username string
		err      error
	)

	if r.FormValue("login-email") != "" && r.FormValue("login-psw") != "" {
		email = r.FormValue("login-email")
	} else if r.FormValue("userContent") != "" {
		email, _, _, err = GoogleLogDecoder(r.FormValue("userContent"))
		if err != nil {
			return err
		}
	} else if strings.HasPrefix(code, "code=") {
		email, username, err = GetEmailGithub(w, r)
		if err != nil {
			return err
		}
	}

	password := r.FormValue("login-psw")

	structure.Html.User = structure.Users{}
	user, err := users.GetUserByEmail(email)
	if err != nil {
		return err
	}

	if user.Username == "" && r.FormValue("userContent") != "" {
		return RegisterHandle(w, r)
	}

	if user.Username == "" && strings.HasPrefix(code, "code=") {
		err = users.AddUser(username, email, "")
		if err != nil {
			return errors.New("username or email is already used by someone")
		}

		structure.Html.User = structure.Users{}
		user, err = users.GetUserByEmail(email)
		if err != nil {
			return err
		}

		users.ConnectUser(structure.Html.User.Id)
		user.Connected = 1
		structure.Html.User = user

		cookies.AddCookies(w)

		return nil
	}

	if user.Username == "" && r.FormValue("userContent") != "" {
		return RegisterHandle(w, r)
	}

	if user.Username == "" || (bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)) != nil && r.FormValue("userContent") == "" && !strings.HasPrefix(code, "code=")) {
		return errors.New("email or password invalid")
	} else if user.Connected == 1 {
		return errors.New("already connected on another device")
	}

	users.ConnectUser(user.Id)
	user.Connected = 1
	structure.Html.User = user

	cookies.AddCookies(w)

	return nil
}

func FaceBookGetData(s string) (string, string, error) {
	datas := make(map[string]string)

	err := json.Unmarshal([]byte(s), &datas)
	if datas["email"] == "" {
		err = errors.New("your facebook email address isn't in public")
	}
	return datas["email"], datas["name"], err
}
