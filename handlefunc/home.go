package handlefunc

import (
	"forum/cookies"
	"forum/database/controller/posts"
	"forum/structure"
	"net/http"
	"strconv"
	"strings"
	"time"
)

// this function manages the home page
func HomeHandle(w http.ResponseWriter, r *http.Request) {
	structure.Html.User = structure.Users{}
	cookies.ConnectCookies(w, r)

	if r.URL.Path != "/home" {
		DisconnectHandle(w, r)
		http.Redirect(w, r, "/home", http.StatusSeeOther)
		return
	}

	code := strings.Split(r.URL.String(), "?")[len(strings.Split(r.URL.String(), "?"))-1]

	if r.FormValue("login-email") != "" && r.FormValue("login-psw") != "" || r.FormValue("userContent") != "" || r.FormValue("FacebookUserContent") != "" || strings.HasPrefix(code, "code=") {
		err := LogInHandle(w, r)

		if err != nil {
			time.Sleep(2 * time.Second)

			structure.Html.Error.LoginError = err.Error()

			DisconnectHandle(w, r)
			http.Redirect(w, r, r.URL.Path+"#login-form", http.StatusSeeOther)
			return
		}

		DisconnectHandle(w, r)
		http.Redirect(w, r, "/home", http.StatusSeeOther)
		return
	}

	if r.FormValue("email") != "" && r.FormValue("psw") != "" && r.FormValue("username") != "" && r.FormValue("confirm-psw") != "" {
		err := RegisterHandle(w, r)

		if err != nil {
			structure.Html.Error.RegisterError = err.Error()

			DisconnectHandle(w, r)
			http.Redirect(w, r, r.URL.Path+"#register-form", http.StatusSeeOther)
			return
		}

		DisconnectHandle(w, r)
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	if r.Method != "POST" {
		posts.GetPosts()
		structure.Tpl.ExecuteTemplate(w, "user.html", structure.Html)

		structure.Html.Error.LoginError = ""
		structure.Html.Error.RegisterError = ""
		return
	}

	filterCategories := r.FormValue("categories")
	if filterCategories != "" && filterCategories != "all" {
		i, _ := strconv.Atoi(filterCategories)
		posts.GetPostsByCategorieId(i)

		structure.Tpl.ExecuteTemplate(w, "user.html", structure.Html)
		return
	}

	posts.GetPosts()
	structure.Tpl.ExecuteTemplate(w, "user.html", structure.Html)
}
