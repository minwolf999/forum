package handlefunc

import (
	"forum/cookies"
	"forum/database/controller/comment"
	"forum/database/controller/likesComment"
	likespost "forum/database/controller/likesPost"
	"forum/database/controller/posts"
	"forum/structure"
	"forum/verificationFunction"
	"net/http"
	"strconv"
	"strings"
	"time"
)

// this function manages the Post information page>
func PostHandle(w http.ResponseWriter, r *http.Request) {
	structure.Html.Error.CreatCommentError = ""
	structure.Html.User = structure.Users{}

	DisconnectHandle(w, r)
	cookies.ConnectCookies(w, r)

	PostIdStr := strings.Split(r.URL.Path, "/")[len(strings.Split(r.URL.Path, "/"))-1]

	PostId, err := strconv.Atoi(PostIdStr)
	if err != nil {
		DisconnectHandle(w, r)
		http.Redirect(w, r, "/home", http.StatusSeeOther)
		return
	}

	posts.GetPostsById(PostId)
	comment.GetCommentsByPostId(PostId)

	if len(structure.Html.Posts) == 0 {
		DisconnectHandle(w, r)
		http.Redirect(w, r, "/home", http.StatusSeeOther)
		return
	}

	if r.Method != "POST" {
		structure.Tpl.ExecuteTemplate(w, "post.html", structure.Html)

		structure.Html.Error.LoginError = ""
		structure.Html.Error.RegisterError = ""
		return
	}

	if r.FormValue("login-email") != "" && r.FormValue("login-psw") != "" || r.FormValue("userContent") != "" {
		err := LogInHandle(w, r)

		if err != nil {
			time.Sleep(2 * time.Second)

			structure.Html.Error.LoginError = err.Error()

			DisconnectHandle(w, r)
			http.Redirect(w, r, r.URL.Path+"#login-form", http.StatusSeeOther)
			return
		} else {
			structure.Html.Error.LoginError = ""
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
		} else {
			structure.Html.Error.LoginError = ""
		}

		DisconnectHandle(w, r)
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	like_dislike := r.FormValue("like/dislike-post")
	if like_dislike == "like" && structure.Html.User.Connected == 1 {
		likespost.AddLikePost(PostId, structure.Html.User.Id)
	} else if like_dislike == "dislike" && structure.Html.User.Connected == 1 {
		likespost.AddDislikePost(PostId, structure.Html.User.Id)
	}

	if r.FormValue("like-comment") != "" && structure.Html.User.Connected == 1 {
		like_dislike_comment := r.FormValue("like-comment")
		like_dislike_comment_id, _ := strconv.Atoi(like_dislike_comment)
		likesComment.AddLikeComment(like_dislike_comment_id, structure.Html.User.Id)
	} else if r.FormValue("dislike-comment") != "" && structure.Html.User.Connected == 1 {
		like_dislike_comment := r.FormValue("dislike-comment")
		like_dislike_comment_id, _ := strconv.Atoi(like_dislike_comment)
		likesComment.AddDislikeComment(like_dislike_comment_id, structure.Html.User.Id)
	}

	//Deleting a Post
	delete_post := r.FormValue("delete-post")
	if delete_post == "delete-post" && structure.Html.User.Role == 1 || delete_post == "delete-post" && structure.Html.User.Role == 2 {
		posts.DeletePost(PostId)
		//Deleting Comments link to the Post
		comment.DeleteCommentsOfPost(PostId)
		//Deleting Like/Dislike link to the Post
		likespost.DeleteLikePost(PostId)
		likespost.DeleteDislikePost(PostId)
		http.Redirect(w, r, "/home", http.StatusSeeOther)
		return
	}

	//Deleting a Comment
	delete_comment := r.FormValue("delete-comment")
	delete_comment_id, err := strconv.Atoi(delete_comment)
	if err == nil {
		if structure.Html.User.Role == 1 || structure.Html.User.Role == 2 {
			comment.DeleteComment(delete_comment_id)
			//Deleting Like/Dislike link to the comment
			likesComment.DeleteLikeComment(delete_comment_id)
			likesComment.DeleteDislikeComment(delete_comment_id)
			http.Redirect(w, r, r.URL.Path, http.StatusSeeOther)
			return
		}
	}

	//Reporting a Post
	report_post := r.FormValue("report-post")
	report_post_id, err := strconv.Atoi(report_post)
	if err == nil {
		if structure.Html.User.Role == 1 {
			posts.AddToReportPosts(report_post_id)
		}
	}

	if r.FormValue("text") != "" && structure.Html.User.Connected == 1 && verificationFunction.IsValidMessage(r.FormValue("text")) {
		comment.CreatNewComment(PostId, structure.Html.User.Id, r.FormValue("text"), "TEXT")
	} else if r.FormValue("text") != "" && structure.Html.User.Connected == 1 && !verificationFunction.IsValidMessage(r.FormValue("text")) {
		structure.Html.Error.CreatCommentError = "there is an empty field in your comment"
	}

	posts.GetPostsById(PostId)
	comment.GetCommentsByPostId(PostId)

	structure.Tpl.ExecuteTemplate(w, "post.html", structure.Html)
	structure.Html.Error.LoginError = ""
	structure.Html.Error.RegisterError = ""
}
