package likespost

import "forum/database/initialisation"

// Delete a likes of a post using its ID and the ID of the User
func RemoveLikePost(idPost, idUser int) error {
	db, err := initialisation.OpenBDD()
	if err != nil {
		return err
	}
	defer db.Close()

	_, err = db.Exec("DELETE FROM `LikesPosts` WHERE `IdPost`=? AND `IdUser`=?", idPost, idUser)
	return err
}

// Delete a dislikes of a post using its ID and the ID of the User
func RemoveDislikePost(idPost, idUser int) error {
	db, err := initialisation.OpenBDD()
	if err != nil {
		return err
	}
	defer db.Close()

	_, err = db.Exec("DELETE FROM `DislikesPosts` WHERE `IdPost`=? AND `IdUser`=?", idPost, idUser)
	return err
}
