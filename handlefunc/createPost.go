package handlefunc

import (
	"forum/cookies"
	"forum/database/controller/categories"
	"forum/database/controller/posts"
	"forum/structure"
	"forum/verificationFunction"
	"net/http"
	"strconv"
	"strings"
)

// this function manages the new post creation page
func CreatePostHandle(w http.ResponseWriter, r *http.Request) {
	structure.Html.Error.CreatPostError = ""

	if r.URL.Path != "/createPost" {
		http.Redirect(w, r, "/createPost", http.StatusSeeOther)
		return
	}

	cookies.ConnectCookies(w, r)
	if structure.Html.User.Email == "" {
		http.Redirect(w, r, "/home", http.StatusSeeOther)
		return
	}

	if r.Method != "POST" {
		structure.Tpl.ExecuteTemplate(w, "createPost.html", structure.Html)
		return
	}
	/*
		err := r.ParseMultipartForm(29 << 20)
		if err != nil {
			fmt.Println("Error ParseMultipartForm: ", err)
			http.Redirect(w, r, "/creat-post/", http.StatusSeeOther)
		} */
	// Get the image from the page
	ImagePath, handler, _ := r.FormFile("Img")

	if handler.Size >= 3e+6 {
		structure.Html.Error.CreatPostError = "This size of image is above 30mb"

		structure.Tpl.ExecuteTemplate(w, "createPost.html", structure.Html)
		structure.Html.Error.CreatPostError = ""
		return
	}
	//call the function for storing the image’
	Path := posts.FormatImg(ImagePath, handler)
	postTitle := r.FormValue("post-title")
	postDescription := r.FormValue("description")
	postCategories := strings.Join(r.Form["categories"], "|")

	for _, y := range r.Form["categories"] {
		idCat, err := strconv.Atoi(y)
		if err != nil {
			structure.Html.Error.CreatPostError = "one of the categories does not exist"

			structure.Tpl.ExecuteTemplate(w, "createPost.html", structure.Html)
			structure.Html.Error.CreatPostError = ""
			return
		}

		str, err := categories.GetCategoriesById(idCat)
		if err != nil || str == "" {
			structure.Html.Error.CreatPostError = "one of the categories does not exist"

			structure.Tpl.ExecuteTemplate(w, "createPost.html", structure.Html)
			structure.Html.Error.CreatPostError = ""
			return
		}
	}

	if postCategories != "" && verificationFunction.IsValidMessage(postTitle) && verificationFunction.IsValidMessage(postDescription) {
		err := posts.CreateNewPost(postCategories, structure.Html.User.Id, postTitle, postDescription, Path)
		if err != nil {
			structure.Html.Error.CreatPostError = err.Error()

			structure.Tpl.ExecuteTemplate(w, "createPost.html", structure.Html)
			structure.Html.Error.CreatPostError = ""
			return
		}
	} else {
		structure.Html.Error.CreatPostError = "there is an empty field in your post"

		structure.Tpl.ExecuteTemplate(w, "createPost.html", structure.Html)
		structure.Html.Error.CreatPostError = ""
		return
	}

	posts.GetPostsByName(postTitle)

	DisconnectHandle(w, r)
	http.Redirect(w, r, "/post/"+strconv.Itoa(structure.Html.Posts[0].Id), http.StatusSeeOther)
}
