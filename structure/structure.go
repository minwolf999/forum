package structure

import (
	"html/template"
)

var (
	Html Htmls
	Tpl  *template.Template
)

type Htmls struct {
	User        Users
	Categories  []Categorie
	Posts       []Post
	Comments    []Comment
	AllReqMod   []Users
	AllMod      []Users
	ReportPosts []Post

	Error MyError
}

type Users struct {
	Id               int
	Username         string
	Email            string
	Password         string
	RegistrationDate string
	Role             int
	Connected        int
	PromoteRequest   int
	Uuid			 string

	ImagePath string
}

type Categorie struct {
	Id   int
	Name string
}

type Post struct {
	Id                  int
	NameCreator         string
	CreatorImageProfile string
	NameCategories      []string
	Name                string
	Description         string
	CreationDate        string
	ReplyQuantity       int
	Likes               Like
	Dislikes            Like
	ImagePath           any
}

type Comment struct {
	Id           int
	IdPost       int
	Text         string
	CreationDate string
	ImagePath    any
	Likes        Like
	Dislikes     Like

	NameCreator      string
	CreatorImagePath any
}

type Like struct {
	Quantity int
	UserId   []int
}

type MyError struct {
	LoginError            string
	RegisterError         string
	CreatPostError        string
	FilterError           string
	ModificationUserError string
	CreatCommentError     string
	LikeError             string
}
