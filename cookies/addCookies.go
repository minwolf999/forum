package cookies

import (
	"forum/database/controller/users"
	"forum/structure"
	"net/http"
	"time"
)

// Create a cookie for the user
func AddCookies(w http.ResponseWriter) {
	uuid := newUUID()
	cookieEmail := http.Cookie{
		Name:    "UUID",
		Value:   uuid.String(),
		Expires: time.Now().Add(48 * time.Hour),
		SameSite: http.SameSiteLaxMode,
		Path: "/",
	}
	users.UpdateUserUUID(structure.Html.User.Id,uuid.String())
	
	http.SetCookie(w, &cookieEmail)
}
