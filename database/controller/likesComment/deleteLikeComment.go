package likesComment

import (
	"forum/database/initialisation"
)

// Delete like of comment in database
func DeleteLikeComment(idComment int) error {
	db, err := initialisation.OpenBDD()
	if err != nil {
		return err
	}
	defer db.Close()

	_, err = db.Exec("DELETE FROM `LikesComments` WHERE `IdComment` =?", idComment)
	return err
}

// Delete dislike of comment in database
func DeleteDislikeComment(idComment int) error {
	db, err := initialisation.OpenBDD()
	if err != nil {
		return err
	}
	defer db.Close()

	_, err = db.Exec("DELETE FROM `DislikesComments` WHERE `IdComment` =?", idComment)
	return err
}
