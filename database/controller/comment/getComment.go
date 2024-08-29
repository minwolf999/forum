package comment

import (
	"forum/database/controller/likesComment"
	"forum/database/controller/users"
	"forum/database/initialisation"
	"forum/structure"
)

// Retrieve all comments on a Post by its identifier
func GetCommentsByPostId(id int) error {
	db, err := initialisation.OpenBDD()
	if err != nil {
		return err
	}
	defer db.Close()

	datas, err := db.Query("SELECT * FROM `Comments` WHERE `IdPost`=?", id)
	if err != nil {
		return err
	}
	defer datas.Close()

	structure.Html.Comments = []structure.Comment{}
	for datas.Next() {
		var tmp structure.Comment
		var idCreator int

		err = datas.Scan(&tmp.Id, &tmp.IdPost, &idCreator, &tmp.Text, &tmp.CreationDate, &tmp.ImagePath)
		if err != nil {
			return err
		}

		name, creatorImage, _ := users.GetUserById(idCreator)
		tmp.NameCreator = name
		tmp.CreatorImagePath = creatorImage

		tmp.Likes.UserId, _ = likesComment.GetLikesComment(tmp.Id)
		tmp.Likes.Quantity = len(tmp.Likes.UserId)

		tmp.Dislikes.UserId, _ = likesComment.GetDislikesComment(tmp.Id)
		tmp.Dislikes.Quantity = len(tmp.Dislikes.UserId)

		structure.Html.Comments = append(structure.Html.Comments, tmp)
	}

	return nil
}
