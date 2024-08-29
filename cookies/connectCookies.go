package cookies

import (
	"errors"
	"forum/database/controller/users"
	"forum/structure"
	"net/http"
)

// Connect a User with his Cookie
func ConnectCookies(w http.ResponseWriter, r *http.Request) error {
	user := structure.Users{}

	UUID, err := r.Cookie("UUID")
	if err != nil {
		return err
	}

	user, err = users.GetUserByUUID(UUID.Value)
	if err != nil || user.Email == "" {
		return err
	}

	if user.Connected == 1 {
		return errors.New("already connected on another device")
	}

	err = users.ConnectUser(user.Id)

	user.Connected = 1
	structure.Html.User = user
	AddCookies(w)

	return err
}

// Get the information of an User with a Cookie
func GetCookies(w http.ResponseWriter, r *http.Request) error {
	var user structure.Users

	UUID, err := r.Cookie("UUID")
	if err != nil {
		return err
	}

	user, err = users.GetUserByUUID(UUID.Value)
	if err != nil {
		return err
	} else if user.Email == "" {
		return errors.New("nobody found")
	}

	structure.Html.User = user
	return nil
}