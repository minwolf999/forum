package handlefunc

import (
	"forum/cookies"
	"forum/database/controller/users"
	"forum/structure"
	"net/http"
	"time"
)

// this function manages deconnection in the bdd
func DisconnectHandle(w http.ResponseWriter, r *http.Request) {
	cookies.GetCookies(w, r)
	users.DisconnectUser(structure.Html.User.Id)
}

// this function manages deconnection in the bdd and delete the cookie in the browser of the user
func CompleteDisconnectHandle(w http.ResponseWriter, r *http.Request) {
	DisconnectHandle(w, r)

	cookieEmail := http.Cookie{
		Name:    "UUID",
		Value:   "",
		Expires: time.Now().Add(-time.Hour),
	}
	http.SetCookie(w, &cookieEmail)

	http.Redirect(w, r, "/home", http.StatusSeeOther)
}
