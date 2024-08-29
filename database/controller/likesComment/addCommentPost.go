package likesComment

import (
	"forum/database/initialisation"
	"forum/structure"
)

// Add a like to a comment in the database or delete it if there is already one. Delete the dislike if there is one too
func AddLikeComment(idComment, idUser int) error {
	db, err := initialisation.OpenBDD()
	if err != nil {
		return err
	}
	defer db.Close()

	for _, v := range structure.Html.Comments {
		for _, i := range v.Likes.UserId {
			if i == idUser && v.Id == idComment {
				return RemoveLikeComment(idComment, idUser)
			}
		}

		for _, i := range v.Dislikes.UserId {
			if i == idUser && v.Id == idComment {
				RemoveDislikeComment(idComment, idUser)
			}
		}
	}

	_, err = db.Exec("INSERT INTO `LikesComments`(`IdComment`, `IdUser`) VALUES(?,?)", idComment, idUser)
	return err
}

// Add a dislike to a comment in the database or delete it if there is already one. Delete the like if there is one too
func AddDislikeComment(idComment, idUser int) error {
	db, err := initialisation.OpenBDD()
	if err != nil {
		return err
	}
	defer db.Close()

	for _, v := range structure.Html.Comments {
		for _, i := range v.Dislikes.UserId {
			if i == idUser && v.Id == idComment {
				return RemoveDislikeComment(idComment, idUser)
			}
		}

		for _, i := range v.Likes.UserId {
			if i == idUser && v.Id == idComment {
				RemoveLikeComment(idComment, idUser)
			}
		}
	}

	_, err = db.Exec("INSERT INTO `DislikesComments`(`IdComment`, `IdUser`) VALUES(?,?)", idComment, idUser)
	return err
}
