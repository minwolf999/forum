package likespost

import (
	"forum/database/initialisation"
)

// Delete likes of post in database
func DeleteLikePost(idPost int) error {
	db, err := initialisation.OpenBDD()
	if err != nil {
		return err
	}
	defer db.Close()

	_, err = db.Exec("DELETE FROM `LikesPosts` WHERE `IdPost` =?", idPost)
	return err
}

// Delete dislikes of post in database
func DeleteDislikePost(idPost int) error {
	db, err := initialisation.OpenBDD()
	if err != nil {
		return err
	}
	defer db.Close()

	_, err = db.Exec("DELETE FROM `DislikesPosts` WHERE `IdPost` =?", idPost)
	return err
}
