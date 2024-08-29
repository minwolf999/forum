package likespost

import (
	"forum/database/initialisation"
	"forum/structure"
)

// Add a like to a post in the database or delete it if there is already one. Delete the dislike if there is one too
func AddLikePost(idPost, idUser int) error {
	db, err := initialisation.OpenBDD()
	if err != nil {
		return err
	}
	defer db.Close()

	for _, v := range structure.Html.Posts {
		for _, i := range v.Likes.UserId {
			if i == idUser && v.Id == idPost {
				return RemoveLikePost(idPost, idUser)
			}
		}

		for _, i := range v.Dislikes.UserId {
			if i == idUser && v.Id == idPost {
				RemoveDislikePost(idPost, idUser)
			}
		}
	}

	_, err = db.Exec("INSERT INTO `LikesPosts`(`IdPost`, `IdUser`) VALUES(?,?)", idPost, idUser)
	return err
}

// Add a dislike to a post in the database or delete it if there is already one. Delete the like if there is one too
func AddDislikePost(idPost, idUser int) error {
	db, err := initialisation.OpenBDD()
	if err != nil {
		return err
	}
	defer db.Close()

	for _, v := range structure.Html.Posts {
		for _, i := range v.Dislikes.UserId {
			if i == idUser && v.Id == idPost {
				return RemoveDislikePost(idPost, idUser)
			}
		}

		for _, i := range v.Likes.UserId {
			if i == idUser && v.Id == idPost {
				RemoveLikePost(idPost, idUser)
			}
		}
	}

	_, err = db.Exec("INSERT INTO `DislikesPosts`(`IdPost`, `IdUser`) VALUES(?,?)", idPost, idUser)
	return err
}