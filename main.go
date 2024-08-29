package main

import (
	"fmt"
	"forum/database/controller/categories"
	"forum/database/controller/posts"
	"forum/handlefunc"
	ratelimit "forum/rate-limit"
	"forum/structure"
	"html/template"
	"log"
	"net/http"

	_ "github.com/mattn/go-sqlite3"
)

//const  LIMIT_TIME = time.Second / 10

func init() {
	// os.Remove("structure/database.db")
	// initialisation.CreateBDD()

	categories.GetCategories()
	posts.GetPosts()

	// Parsing all html files
	structure.Tpl = template.Must(template.New("").ParseGlob("front/**/*.html"))

	// Listen statics files
	fs := http.FileServer(http.Dir("front"))
	http.Handle("/front/", http.StripPrefix("/front", fs))
}

func main() {
	tb := ratelimit.NewTokenBucket(10, 1)
	fmt.Print("\033[96m")
	fmt.Println("\033[96mServer started at: https://localhost:8080\033[0m")
	fmt.Print("\033[0m")
	// Create the routes for the differents page
	http.HandleFunc("/", ratelimit.RateLimitedHandler(tb, handlefunc.HomeHandle))
	http.HandleFunc("/home", ratelimit.RateLimitedHandler(tb, handlefunc.HomeHandle))
	http.HandleFunc("/createPost", ratelimit.RateLimitedHandler(tb, handlefunc.CreatePostHandle))
	http.HandleFunc("/profile", ratelimit.RateLimitedHandler(tb, handlefunc.ProfileHandle))
	http.HandleFunc("/post/", ratelimit.RateLimitedHandler(tb, handlefunc.PostHandle))
	http.HandleFunc("/adminPanel", ratelimit.RateLimitedHandler(tb, handlefunc.AdminPanelHandle))
	http.HandleFunc("/disconnect", ratelimit.RateLimitedHandler(tb, handlefunc.CompleteDisconnectHandle))
	http.HandleFunc("/bye", ratelimit.RateLimitedHandler(tb, handlefunc.DisconnectHandle))
	err := http.ListenAndServeTLS(":8080", "certificates/full-cert.crt", "certificates/private-key.key", nil)
	log.Fatal(err)

	//http.ListenAndServe(":8080", nil)
}
