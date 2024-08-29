package posts

import (
	"fmt"
	"forum/database/controller/categories"
	"forum/database/controller/comment"
	likespost "forum/database/controller/likesPost"
	"forum/database/controller/users"
	"forum/database/initialisation"
	"forum/structure"
	verificationfunction "forum/verificationFunction"
	"strconv"
	"strings"
)

// Retrieve all Post in the database
func GetPosts() error {
	db, err := initialisation.OpenBDD()
	if err != nil {
		return err
	}
	defer db.Close()

	datas, err := db.Query("SELECT * FROM `Posts`")
	if err != nil {
		return err
	}
	defer datas.Close()

	structure.Html.Posts = []structure.Post{}
	for datas.Next() {
		var tmp structure.Post
		var idCategoriesString string
		var idCreator int

		err = datas.Scan(&tmp.Id, &idCategoriesString, &idCreator, &tmp.Name, &tmp.Description, &tmp.CreationDate, &tmp.ImagePath)
		if err != nil {
			return err
		}

		name, imagePath, _ := users.GetUserById(idCreator)
		tmp.NameCreator = name
		tmp.CreatorImageProfile = imagePath

		tmp.Likes.UserId, _ = likespost.GetLikesPost(tmp.Id)
		tmp.Likes.Quantity = len(tmp.Likes.UserId)

		tmp.Dislikes.UserId, _ = likespost.GetDislikesPost(tmp.Id)
		tmp.Dislikes.Quantity = len(tmp.Dislikes.UserId)

		idCategoriesTabString := strings.Split(idCategoriesString, "|")
		for _, v := range idCategoriesTabString {
			i, _ := strconv.Atoi(v)
			temp, err := categories.GetCategoriesById(i)
			if err != nil {
				fmt.Println(err)
				return err
			}

			tmp.NameCategories = append(tmp.NameCategories, temp)
		}

		comment.GetCommentsByPostId(tmp.Id)
		tmp.ReplyQuantity = len(structure.Html.Comments)
		structure.Html.Comments = []structure.Comment{}

		structure.Html.Posts = append(structure.Html.Posts, tmp)
	}

	return nil
}

// Retrieve all ReportPost in the database
func GetReportPosts() ([]int, error) {
	db, err := initialisation.OpenBDD()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	datas, err := db.Query("SELECT `IdPost` FROM `ReportPosts`")
	if err != nil {
		return nil, err
	}
	defer datas.Close()

	structure.Html.ReportPosts = []structure.Post{}
	var idPostTable []int
	for datas.Next() {
		var idPost int

		err = datas.Scan(&idPost)
		if err != nil {
			return nil, err
		}
		idPostTable = append(idPostTable, idPost)
	}

	return idPostTable, err
}

// Retrieve all Post with a given category in the database
func GetPostsByCategorieId(idCategorie int) error {
	db, err := initialisation.OpenBDD()
	if err != nil {
		return err
	}
	defer db.Close()

	datas, err := db.Query("SELECT * FROM `Posts`", idCategorie)
	if err != nil {
		return err
	}
	defer datas.Close()

	structure.Html.Posts = []structure.Post{}
	for datas.Next() {
		var tmp structure.Post
		var idCategoriesString string
		var idCreator int

		err = datas.Scan(&tmp.Id, &idCategoriesString, &idCreator, &tmp.Name, &tmp.Description, &tmp.CreationDate, &tmp.ImagePath)
		if err != nil {
			return err
		}

		name, imagePath, _ := users.GetUserById(idCreator)
		tmp.NameCreator = name
		tmp.CreatorImageProfile = imagePath

		idCategoriesTabString := strings.Split(idCategoriesString, "|")
		if verificationfunction.TabNotContain(idCategoriesTabString, strconv.Itoa(idCategorie)) {
			continue
		}

		tmp.Likes.UserId, _ = likespost.GetLikesPost(tmp.Id)
		tmp.Likes.Quantity = len(tmp.Likes.UserId)

		tmp.Dislikes.UserId, _ = likespost.GetDislikesPost(tmp.Id)
		tmp.Dislikes.Quantity = len(tmp.Dislikes.UserId)

		for _, v := range idCategoriesTabString {
			i, _ := strconv.Atoi(v)
			temp, err := categories.GetCategoriesById(i)
			if err != nil {
				fmt.Println(err)
				return err
			}

			tmp.NameCategories = append(tmp.NameCategories, temp)
		}

		comment.GetCommentsByPostId(tmp.Id)
		tmp.ReplyQuantity = len(structure.Html.Comments)
		structure.Html.Comments = []structure.Comment{}

		structure.Html.Posts = append(structure.Html.Posts, tmp)
	}

	return nil
}

// Retrieve a post in the database with his name
func GetPostsByName(s string) error {
	db, err := initialisation.OpenBDD()
	if err != nil {
		return err
	}
	defer db.Close()

	s += "%"
	datas, err := db.Query("SELECT * FROM `Posts` WHERE `Name` LIKE ?", s)
	if err != nil {
		return err
	}
	defer datas.Close()

	structure.Html.Posts = []structure.Post{}
	for datas.Next() {
		var tmp structure.Post
		var idCategoriesString string
		var idCreator int

		err = datas.Scan(&tmp.Id, &idCategoriesString, &idCreator, &tmp.Name, &tmp.Description, &tmp.CreationDate, &tmp.ImagePath)
		if err != nil {
			return err
		}

		name, imagePath, _ := users.GetUserById(idCreator)
		tmp.NameCreator = name
		tmp.CreatorImageProfile = imagePath

		tmp.Likes.UserId, _ = likespost.GetLikesPost(tmp.Id)
		tmp.Likes.Quantity = len(tmp.Likes.UserId)

		tmp.Dislikes.UserId, _ = likespost.GetDislikesPost(tmp.Id)
		tmp.Dislikes.Quantity = len(tmp.Dislikes.UserId)

		idCategoriesTabString := strings.Split(idCategoriesString, "|")
		for _, v := range idCategoriesTabString {
			i, _ := strconv.Atoi(v)
			temp, err := categories.GetCategoriesById(i)
			if err != nil {
				fmt.Println(err)
				return err
			}

			tmp.NameCategories = append(tmp.NameCategories, temp)
		}

		comment.GetCommentsByPostId(tmp.Id)
		tmp.ReplyQuantity = len(structure.Html.Comments)
		structure.Html.Comments = []structure.Comment{}

		structure.Html.Posts = append(structure.Html.Posts, tmp)
	}

	return nil
}

// Retrieve a post in the database with the ID of the creator
func GetPostsByUserId(id int) error {
	db, err := initialisation.OpenBDD()
	if err != nil {
		return err
	}
	defer db.Close()

	datas, err := db.Query("SELECT * FROM `Posts` WHERE `IdCreator`=?", id)
	if err != nil {
		return err
	}
	defer datas.Close()

	structure.Html.Posts = []structure.Post{}
	for datas.Next() {
		var tmp structure.Post
		var idCategoriesString string
		var idCreator int

		err = datas.Scan(&tmp.Id, &idCategoriesString, &idCreator, &tmp.Name, &tmp.Description, &tmp.CreationDate, &tmp.ImagePath)
		if err != nil {
			return err
		}

		name, imagePath, _ := users.GetUserById(idCreator)
		tmp.NameCreator = name
		tmp.CreatorImageProfile = imagePath

		tmp.Likes.UserId, _ = likespost.GetLikesPost(tmp.Id)
		tmp.Likes.Quantity = len(tmp.Likes.UserId)

		tmp.Dislikes.UserId, _ = likespost.GetDislikesPost(tmp.Id)
		tmp.Dislikes.Quantity = len(tmp.Dislikes.UserId)

		idCategoriesTabString := strings.Split(idCategoriesString, "|")
		for _, v := range idCategoriesTabString {
			i, _ := strconv.Atoi(v)
			temp, err := categories.GetCategoriesById(i)
			if err != nil {
				fmt.Println(err)
				return err
			}

			tmp.NameCategories = append(tmp.NameCategories, temp)
		}

		comment.GetCommentsByPostId(tmp.Id)
		tmp.ReplyQuantity = len(structure.Html.Comments)
		structure.Html.Comments = []structure.Comment{}

		structure.Html.Posts = append(structure.Html.Posts, tmp)
	}

	return nil
}

// Retrieve a post in the database with his ID
func GetPostsById(id int) error {
	db, err := initialisation.OpenBDD()
	if err != nil {
		return err
	}
	defer db.Close()

	datas, err := db.Query("SELECT * FROM `Posts` WHERE `Id`=?", id)
	if err != nil {
		return err
	}
	defer datas.Close()

	structure.Html.Posts = []structure.Post{}
	for datas.Next() {
		var tmp structure.Post
		var idCategoriesString string
		var idCreator int

		err = datas.Scan(&tmp.Id, &idCategoriesString, &idCreator, &tmp.Name, &tmp.Description, &tmp.CreationDate, &tmp.ImagePath)
		if err != nil {
			return err
		}

		name, imagePath, _ := users.GetUserById(idCreator)
		tmp.NameCreator = name
		tmp.CreatorImageProfile = imagePath

		tmp.Likes.UserId, _ = likespost.GetLikesPost(tmp.Id)
		tmp.Likes.Quantity = len(tmp.Likes.UserId)

		tmp.Dislikes.UserId, _ = likespost.GetDislikesPost(tmp.Id)
		tmp.Dislikes.Quantity = len(tmp.Dislikes.UserId)

		idCategoriesTabString := strings.Split(idCategoriesString, "|")
		for _, v := range idCategoriesTabString {
			i, _ := strconv.Atoi(v)
			temp, err := categories.GetCategoriesById(i)
			if err != nil {
				fmt.Println(err)
				return err
			}

			tmp.NameCategories = append(tmp.NameCategories, temp)
		}

		comment.GetCommentsByPostId(tmp.Id)
		tmp.ReplyQuantity = len(structure.Html.Comments)
		structure.Html.Comments = []structure.Comment{}

		structure.Html.Posts = append(structure.Html.Posts, tmp)
	}

	return nil
}

// Put post that get request to ReportPost
func GetPostIntoReportPost(id int) error {
	db, err := initialisation.OpenBDD()
	if err != nil {
		return err
	}
	defer db.Close()

	datas, err := db.Query("SELECT * FROM `Posts` WHERE `Id`=?", id)
	if err != nil {
		return err
	}
	defer datas.Close()

	for datas.Next() {
		var tmp structure.Post
		var idCategoriesString string
		var idCreator int

		err = datas.Scan(&tmp.Id, &idCategoriesString, &idCreator, &tmp.Name, &tmp.Description, &tmp.CreationDate, &tmp.ImagePath)
		if err != nil {
			return err
		}

		name, imagePath, _ := users.GetUserById(idCreator)
		tmp.NameCreator = name
		tmp.CreatorImageProfile = imagePath

		tmp.Likes.UserId, _ = likespost.GetLikesPost(tmp.Id)
		tmp.Likes.Quantity = len(tmp.Likes.UserId)

		tmp.Dislikes.UserId, _ = likespost.GetDislikesPost(tmp.Id)
		tmp.Dislikes.Quantity = len(tmp.Dislikes.UserId)

		idCategoriesTabString := strings.Split(idCategoriesString, "|")
		for _, v := range idCategoriesTabString {
			i, _ := strconv.Atoi(v)
			temp, err := categories.GetCategoriesById(i)
			if err != nil {
				fmt.Println(err)
				return err
			}

			tmp.NameCategories = append(tmp.NameCategories, temp)
		}

		comment.GetCommentsByPostId(tmp.Id)
		tmp.ReplyQuantity = len(structure.Html.Comments)
		structure.Html.Comments = []structure.Comment{}

		structure.Html.ReportPosts = append(structure.Html.ReportPosts, tmp)
	}

	return nil
}
