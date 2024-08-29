package likesComment

import "forum/database/initialisation"

// Delete a likes of a comment using its ID and the ID of the User
func RemoveLikeComment(idComment, idUser int) error {
	db, err := initialisation.OpenBDD()
	if err != nil {
		return err
	}
	defer db.Close()

	_, err = db.Exec("DELETE FROM `LikesComments` WHERE `IdComment`=? AND `IdUser`=?", idComment, idUser)
	return err
}

// Delete a dislikes of a comment using its ID and the ID of the User
func RemoveDislikeComment(idComment, idUser int) error {
	db, err := initialisation.OpenBDD()
	if err != nil {
		return err
	}
	defer db.Close()

	_, err = db.Exec("DELETE FROM `DislikesComments` WHERE `IdComment`=? AND `IdUser`=?", idComment, idUser)
	return err
}