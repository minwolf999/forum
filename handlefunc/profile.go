package handlefunc

import (
	"forum/cookies"
	"forum/database/controller/posts"
	"forum/database/controller/users"
	"forum/structure"
	"forum/verificationFunction"
	"net/http"
	"strconv"
)

// this function manages the User information page
func ProfileHandle(w http.ResponseWriter, r *http.Request) {
	structure.Html.User = structure.Users{}
	structure.Html.Posts = []structure.Post{}
	DisconnectHandle(w, r)
	cookies.ConnectCookies(w, r)

	if structure.Html.User.Username == "Visitor" || structure.Html.User.Username == "" {
		DisconnectHandle(w, r)
		http.Redirect(w, r, "/home", http.StatusSeeOther)
		return
	}

	like := r.FormValue("filtre")
	if like == "liked" {
		posts.GetPosts()

	retry:
		for i := range structure.Html.Posts {
			if verificationFunction.TabNotContainInt(structure.Html.Posts[i].Likes.UserId, structure.Html.User.Id) {
				if i != len(structure.Html.Posts)-1 {
					structure.Html.Posts = append(structure.Html.Posts[:i], structure.Html.Posts[i+1:]...)
					goto retry
				} else {
					structure.Html.Posts = structure.Html.Posts[:i]
					goto retry
				}
			}
		}

		structure.Tpl.ExecuteTemplate(w, "profil.html", structure.Html)
		return
	}

	posts.GetPostsByUserId(structure.Html.User.Id)

	if r.FormValue("login-email") != "" && r.FormValue("login-psw") != "" {
		err := LogInHandle(w, r)

		if err != nil {
			structure.Html.Error.LoginError = err.Error()
		}

		DisconnectHandle(w, r)
		http.Redirect(w, r, r.URL.Path, http.StatusSeeOther)
		return
	}

	if r.FormValue("email") != "" && r.FormValue("psw") != "" && r.FormValue("username") != "" && r.FormValue("confirm-psw") != "" {
		err := RegisterHandle(w, r)

		if err != nil {
			structure.Html.Error.RegisterError = err.Error()
		}

		DisconnectHandle(w, r)
		http.Redirect(w, r, r.URL.Path, http.StatusSeeOther)
		return
	}

	//request to be moderator
	request_mod := r.FormValue("mod-request")
	request_mod_id, err := strconv.Atoi(request_mod)
	if err == nil {
		users.UpdatePromoteRequest(request_mod_id, 1)
	}

	structure.Tpl.ExecuteTemplate(w, "profil.html", structure.Html)
}
