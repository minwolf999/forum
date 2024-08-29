package handlefunc

import (
	"forum/cookies"
	"forum/database/controller/posts"
	"forum/database/controller/users"
	"forum/structure"
	"net/http"
	"strconv"
)

func AdminPanelHandle(w http.ResponseWriter, r *http.Request) {
	DisconnectHandle(w, r)
	cookies.ConnectCookies(w, r)

	//If not an Admin
	if structure.Html.User.Role != 2 {
		DisconnectHandle(w, r)
		http.Redirect(w, r, "/home", http.StatusSeeOther)
		return
	}

	//For Promote/Demote
	selected_id := r.FormValue("all-users")
	selected_id_int, err := strconv.Atoi(selected_id)
	if err == nil {
		if r.FormValue("promote") == "promote" {
			users.UpdateUserRole(selected_id_int, 1)
			users.UpdatePromoteRequest(selected_id_int, 0)
		} else if r.FormValue("deny") == "deny" {
			users.UpdatePromoteRequest(selected_id_int, 0)
		}
	}

	selected_mod_id := r.FormValue("all-mod")
	selected_mod_id_int, err := strconv.Atoi(selected_mod_id)
	if err == nil {
		if r.FormValue("demote") == "demote" {
			users.UpdateUserRole(selected_mod_id_int, 0)
		}
	}

	//Deleting a ReportPost from list
	delete_report_post := r.FormValue("delete-report-post")
	delete_report_post_id, err := strconv.Atoi(delete_report_post)
	if err == nil {
		posts.DeleteReportPost(delete_report_post_id)
	}

	allReportPostID, err := posts.GetReportPosts()
	if err == nil {
		structure.Html.ReportPosts = []structure.Post{}
		for _, id := range allReportPostID {
			posts.GetPostIntoReportPost(id)
		}
	}

	users.GetAllReqMod()
	users.GetAllMod()

	structure.Tpl.ExecuteTemplate(w, "adminpanel.html", structure.Html)
}
